package main

import (
	"sync"

	"github.com/teleclimber/DropServer/cmd/ds-host/domain"
)

// AppVersionEvents notifies that a change was made to an app version
type AppVersionEvents struct {
	subLock     sync.Mutex
	subscribers []chan<- domain.AppID
}

// Send sends an app is changed
func (e *AppVersionEvents) Send(appID domain.AppID) {
	e.subLock.Lock()
	defer e.subLock.Unlock()
	for _, ch := range e.subscribers {
		ch <- appID
	}
}

// Subscribe to an event for when an app is changed
func (e *AppVersionEvents) Subscribe(ch chan<- domain.AppID) {
	e.subLock.Lock()
	defer e.subLock.Unlock()
	e.removeSubscriber(ch)
	e.subscribers = append(e.subscribers, ch)
}

// Unsubscribe to an event for when an app is changed
func (e *AppVersionEvents) Unsubscribe(ch chan<- domain.AppID) {
	e.subLock.Lock()
	defer e.subLock.Unlock()
	e.removeSubscriber(ch)
}

func (e *AppVersionEvents) removeSubscriber(ch chan<- domain.AppID) {
	for i, c := range e.subscribers {
		if c == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
			close(ch)
		}
	}
}

// InspectSandboxEvents notifes of changes to the sandbox inspect state
type InspectSandboxEvents struct {
	subLock     sync.Mutex
	subscribers []chan<- bool
}

// Send sends an app is changed
func (e *InspectSandboxEvents) Send(inspect bool) {
	e.subLock.Lock()
	defer e.subLock.Unlock()
	for _, ch := range e.subscribers {
		ch <- inspect
	}
}

// Subscribe to an event for when an app is changed
func (e *InspectSandboxEvents) Subscribe(ch chan<- bool) {
	e.subLock.Lock()
	defer e.subLock.Unlock()
	e.removeSubscriber(ch)
	e.subscribers = append(e.subscribers, ch)
}

// Unsubscribe to an event for when an app is changed
func (e *InspectSandboxEvents) Unsubscribe(ch chan<- bool) {
	e.subLock.Lock()
	defer e.subLock.Unlock()
	e.removeSubscriber(ch)
}

func (e *InspectSandboxEvents) removeSubscriber(ch chan<- bool) {
	for i, c := range e.subscribers {
		if c == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
			close(ch)
		}
	}
}
