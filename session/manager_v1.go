//package session
//
//import (
//	"errors"
//	"sync"
//	"time"
//)
//
//// Session 结构体存储会话信息
//type Session struct {
//	ID     string
//	Values map[string]interface{}
//	Expiry time.Time
//}
//
//// SessionManager 管理会话
//type SessionManager struct {
//	sessions map[string]Session
//	mutex    sync.RWMutex
//}
//
//// NewSessionManager 创建一个新的会话管理器
//func NewSessionManager() *SessionManager {
//	return &SessionManager{
//		sessions: make(map[string]Session),
//	}
//}
//
//// GenerateID 生成一个新的会话ID
//func (sm *SessionManager) GenerateID() string {
//	return "session_" + time.Now().Format("2006-01-02T15:04:05")
//}
//
//// NewSession 创建一个新的会话
//func (sm *SessionManager) NewSession() (string, error) {
//	sm.mutex.Lock()
//	defer sm.mutex.Unlock()
//
//	id := sm.GenerateID()
//	session := Session{
//		ID:     id,
//		Values: make(map[string]interface{}),
//		Expiry: time.Now().Add(30 * time.Minute), // 设置会话过期时间为30分钟
//	}
//	sm.sessions[id] = session
//	return id, nil
//}
//
//// GetSession 获取会话
//func (sm *SessionManager) GetSession(id string) (Session, error) {
//	sm.mutex.RLock()
//	defer sm.mutex.RUnlock()
//
//	session, exists := sm.sessions[id]
//	if !exists || time.Now().After(session.Expiry) {
//		return Session{}, ErrSessionNotFound
//	}
//	return session, nil
//}
//
//// UpdateSession 更新会话
//func (sm *SessionManager) UpdateSession(id string, values map[string]interface{}) error {
//	sm.mutex.Lock()
//	defer sm.mutex.Unlock()
//
//	session, exists := sm.sessions[id]
//	if !exists {
//		return ErrSessionNotFound
//	}
//
//	for key, value := range values {
//		session.Values[key] = value
//	}
//	session.Expiry = time.Now().Add(30 * time.Minute) // 更新过期时间
//	sm.sessions[id] = session
//	return nil
//}
//
//// DeleteSession 删除会话
//func (sm *SessionManager) DeleteSession(id string) error {
//	sm.mutex.Lock()
//	defer sm.mutex.Unlock()
//
//	delete(sm.sessions, id)
//	return nil
//}
//
//// Cleanup 清理过期会话
//func (sm *SessionManager) Cleanup() {
//	sm.mutex.Lock()
//	defer sm.mutex.Unlock()
//
//	for id, session := range sm.sessions {
//		if time.Now().After(session.Expiry) {
//			delete(sm.sessions, id)
//		}
//	}
//}
//
//// ErrSessionNotFound 会话未找到错误
//var ErrSessionNotFound = errors.New("session not found")
//
//func main() {
//	// 示例：创建会话管理器并创建一个新会话
//	manager := NewSessionManager()
//	sessionID, err := manager.NewSession()
//	if err != nil {
//		panic(err)
//	}
//
//	// 示例：获取并更新会话
//	session, err := manager.GetSession(sessionID)
//	if err != nil {
//		panic(err)
//	}
//	session.Values["user"] = "John Doe"
//	err = manager.UpdateSession(sessionID, session.Values)
//	if err != nil {
//		panic(err)
//	}
//
//	// 示例：清理过期会话
//	manager.Cleanup()
//}
