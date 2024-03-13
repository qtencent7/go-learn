package build_a_server

import "net/http"

// corsMiddleware 是一个中间件函数，用于设置CORS头部
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置Access-Control-Allow-Origin头部，允许所有域名跨域
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// 如果是预检请求（OPTIONS方法），则设置其他CORS相关的头部
		if r.Method == "OPTIONS" {
			// 设置Access-Control-Allow-Methods头部，允许的HTTP方法
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			// 设置Access-Control-Allow-Headers头部，允许的HTTP请求头
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// 设置Access-Control-Max-Age头部，预检请求的结果可以缓存的时间（秒）
			w.Header().Set("Access-Control-Max-Age", "3600")
			// 返回204 No Content响应
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}
