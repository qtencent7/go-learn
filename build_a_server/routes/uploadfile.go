package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "只允许post请求", http.StatusBadRequest)
		return
	}
	err := r.ParseMultipartForm(32 << 20)

	if err != nil {
		http.Error(w, "解析form错误", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "提取文件错误", http.StatusBadRequest)
		return
	}

	defer file.Close()

	newFile, err := os.Create(filepath.Join("files", header.Filename))

	if err != nil {
		http.Error(w, "文件创建错误", http.StatusInternalServerError)
		return
	}

	defer newFile.Close()

	_, err = io.Copy(newFile, file)

	if err != nil {
		http.Error(w, "文件复制错误", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", header.Filename)
}
