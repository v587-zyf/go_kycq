package httpserver

import (
	rnd "crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	"cqserver/golibs/nw/httpserver/seqqueue"
)

var ErrMissed = errors.New("session missed")

var sessionsMu sync.RWMutex
var sessions = make(map[string]*Session)

type Session struct {
	SessionId   string
	SessionData interface{}
	UserData    interface{}
	kickStatus  int32
	SeqQueue    *seqqueue.SeqQueue
}

// 代表一个被kick的session
type sessionElement struct {
	next       *sessionElement
	session    *Session
	removeTime int64
}

// 管理被kick的session的删除，被kick的session需要延迟一段时间删除，以便提醒客户端该session被kick
// 被延迟删除的session构成一条根据时间排好序的链表
type sessionDeleteManager struct {
	first *sessionElement
	last  *sessionElement
	mu    sync.Mutex
}

var deleteManager = &sessionDeleteManager{}

func init() {
	go deleteManager.doRemove()
}

func (this *sessionDeleteManager) doRemove() {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			this.mu.Lock()
			first, last := this.first, this.last
			this.mu.Unlock()
			elem := first
			now := time.Now().Unix()
			removed := false
			for elem != nil {
				if elem.removeTime > now {
					break
				}
				removed = true
				sessionsMu.Lock()
				delete(sessions, elem.session.SessionId)
				sessionsMu.Unlock()
				if elem == last {
					break
				}
				elem = elem.next
			}
			if removed {
				this.mu.Lock()
				if elem != this.first {
					this.first = elem
				} else {
					if this.last == this.first {
						this.first = nil
						this.last = nil
					} else {
						this.first = this.first.next
					}
				}
				this.mu.Unlock()
			}
		}
	}
	ticker.Stop()
}

func (this *sessionDeleteManager) add(session *Session) {
	removeTime := time.Now().Add(time.Minute).Unix()
	elem := &sessionElement{
		session:    session,
		removeTime: removeTime,
	}
	this.mu.Lock()
	defer this.mu.Unlock()
	if this.first == nil {
		this.first = elem
	}
	if this.last != nil {
		this.last.next = elem
	}
	this.last = elem
}

func GenerateSessionId() (string, error) {
	k := make([]byte, 16)
	if _, err := io.ReadFull(rnd.Reader, k); err != nil {
		return "", nil
	}
	return hex.EncodeToString(k), nil
}

func New(sessionData interface{}, userData interface{}) (*Session, error) {
	sessionId, err := GenerateSessionId()
	if err != nil {
		return nil, err
	}
	queue := seqqueue.New(userData, sessionData.(int))
	session := &Session{
		SessionId:   sessionId,
		SessionData: sessionData,
		UserData:    userData,
		SeqQueue:    queue,
	}
	return session, nil
}

func Get(sessionId string) (*Session, error) {
	sessionsMu.RLock()
	session, ok := sessions[sessionId]
	sessionsMu.RUnlock()
	if !ok {
		return nil, ErrMissed
	}
	return session, nil
}

func Set(sessionId string, session *Session) {
	sessionsMu.Lock()
	sessions[sessionId] = session
	sessionsMu.Unlock()
}

func Remove(sessionId string) {
	session, err := Get(sessionId)
	if err == ErrMissed {
		return
	}
	RemoveSession(session)
}

func RemoveSession(session *Session) {
	if !session.IsKicked() {
		session.SeqQueue.Stop(false)
		sessionsMu.Lock()
		delete(sessions, session.SessionId)
		sessionsMu.Unlock()
	}
}

func (this *Session) Kick(needSync bool, kickType int) {
	this.kickStatus = int32(kickType)
	this.SeqQueue.Stop(needSync)
	deleteManager.add(this)
}

func (this *Session) IsKicked() bool {
	return this.kickStatus > 0
}

func GetWithCookieName(r *http.Request, name string) (*Session, error) {
	sessionId := r.FormValue(name)
	if len(sessionId) == 0 {
		cookie, err := r.Cookie(name)
		if err == nil && cookie != nil {
			sessionId = cookie.Value
		}
	}
	return Get(sessionId)
}

func SetWithCookieName(w http.ResponseWriter, name string, sessionId string, session *Session) {
	cookie := &http.Cookie{
		Name:  name,
		Value: sessionId,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	Set(sessionId, session)
}

func (this *Session) GetKickStatus() int {
	return int(this.kickStatus)
}
