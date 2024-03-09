package session

import (
	"fmt"
	"sync"
	"time"
)

type Session struct {
	ID     string
	Values map[string]interface{}
	Expire time.Time
}

type SessionManager struct {
	sessions map[string]Session
	rwmutex  sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]Session),
	}
}
func (sm *SessionManager) GenerateID() string {
	return "session_" + time.Now().Format("")
}

func (sm *SessionManager) NewSession() string {
	sm.rwmutex.Lock()
	defer sm.rwmutex.Unlock()

	id := sm.GenerateID()
	session := Session{
		ID:     id,
		Values: make(map[string]interface{}),
		Expire: time.Now().Add(30 * time.Minute),
	}
	sm.sessions[id] = session
	return id
}

func (sm *SessionManager) GetSession(id string) Session {
	sm.rwmutex.RLock()
	defer sm.rwmutex.RUnlock()

	session := sm.sessions[id]
	return session
}

func main() {
	manager := NewSessionManager()
	sessionID := manager.NewSession()
	session := manager.GetSession(sessionID)
	fmt.Println(session)
}
