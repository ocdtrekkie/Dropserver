package events

import (
	"time"

	"github.com/teleclimber/DropServer/cmd/ds-host/domain"
)

// AppspacePausedEvents handles appspace pause and unpause events
type AppspacePausedEvents struct {
	subscribers []chan<- domain.AppspacePausedEvent
}

// Send sends an appspace paused or unpaused event
func (e *AppspacePausedEvents) Send(appspaceID domain.AppspaceID, paused bool) {
	p := domain.AppspacePausedEvent{AppspaceID: appspaceID, Paused: paused}
	for _, ch := range e.subscribers {
		ch <- p
	}
}

// Subscribe to an event for when an appspace is paused or unpaused
func (e *AppspacePausedEvents) Subscribe(ch chan<- domain.AppspacePausedEvent) {
	e.removeSubscriber(ch)
	e.subscribers = append(e.subscribers, ch)
}

// Unsubscribe to an event for when an appspace is paused or unpaused
func (e *AppspacePausedEvents) Unsubscribe(ch chan<- domain.AppspacePausedEvent) {
	e.removeSubscriber(ch)
}

func (e *AppspacePausedEvents) removeSubscriber(ch chan<- domain.AppspacePausedEvent) {
	// get a feeling you'll need a mutex to cover subscribers?
	for i, c := range e.subscribers {
		if c == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
		}
	}
}

/////////////////////////////////////////
// migration job events

//MigrationJobStatusEvents forwards events related to the progress of migration jobs
type MigrationJobStatusEvents struct {
	subscribers []chan<- domain.MigrationStatusData
}

// Send sends an appspace status event
func (e *MigrationJobStatusEvents) Send(event domain.MigrationStatusData) {
	for _, ch := range e.subscribers {
		ch <- event
	}
}

// Subscribe to an event to know when the status of an appspace has changed
func (e *MigrationJobStatusEvents) Subscribe(ch chan<- domain.MigrationStatusData) {
	e.removeSubscriber(ch)
	e.subscribers = append(e.subscribers, ch)
}

// Unsubscribe to the event
func (e *MigrationJobStatusEvents) Unsubscribe(appspaceID domain.AppspaceID, ch chan<- domain.MigrationStatusData) {
	e.removeSubscriber(ch)
}

func (e *MigrationJobStatusEvents) removeSubscriber(ch chan<- domain.MigrationStatusData) {
	for i, c := range e.subscribers {
		if c == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
		}
	}
}

////// Apppsace Files Event

// AppspaceFilesEvents notify subscribers that appsapce files
// have been written to outside of normal appspace use.
// Usually this means they were imported, or a backup restored
type AppspaceFilesEvents struct {
	subscribers []chan<- domain.AppspaceID
}

// Send sends an appspace paused or unpaused event
func (e *AppspaceFilesEvents) Send(appspaceID domain.AppspaceID) {
	for _, ch := range e.subscribers {
		ch <- appspaceID
	}
}

// Subscribe to an event for when an appspace is paused or unpaused
func (e *AppspaceFilesEvents) Subscribe(ch chan<- domain.AppspaceID) {
	e.removeSubscriber(ch)
	e.subscribers = append(e.subscribers, ch)
}

// Unsubscribe to an event for when an appspace is paused or unpaused
func (e *AppspaceFilesEvents) Unsubscribe(ch chan<- domain.AppspaceID) {
	e.removeSubscriber(ch)
}

func (e *AppspaceFilesEvents) removeSubscriber(ch chan<- domain.AppspaceID) {
	// get a feeling you'll need a mutex to cover subscribers?
	for i, c := range e.subscribers {
		if c == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
		}
	}
}

////////////////////////////////////////
// Appspace Status events
type appspaceStatusSubscriber struct {
	appspaceID domain.AppspaceID
	ch         chan<- domain.AppspaceStatusEvent
}

// AppspaceStatusEvents handles appspace pause and unpause events
type AppspaceStatusEvents struct {
	subscribers []appspaceStatusSubscriber
}

// Send sends an appspace status event
func (e *AppspaceStatusEvents) Send(appspaceID domain.AppspaceID, event domain.AppspaceStatusEvent) {
	for _, sub := range e.subscribers {
		if sub.appspaceID == appspaceID {
			sub.ch <- event
		}
	}
}

// Subscribe to an event to know when the status of an appspace has changed
func (e *AppspaceStatusEvents) Subscribe(appspaceID domain.AppspaceID, ch chan<- domain.AppspaceStatusEvent) {
	e.removeSubscriber(appspaceID, ch)
	e.subscribers = append(e.subscribers, appspaceStatusSubscriber{appspaceID, ch})
}

// Unsubscribe to the event
func (e *AppspaceStatusEvents) Unsubscribe(appspaceID domain.AppspaceID, ch chan<- domain.AppspaceStatusEvent) {
	e.removeSubscriber(appspaceID, ch)
}

// UnsubscribeChannel removes the channel from all subscriptions
func (e *AppspaceStatusEvents) UnsubscribeChannel(ch chan<- domain.AppspaceStatusEvent) {
	for i := len(e.subscribers) - 1; i >= 0; i-- {
		if e.subscribers[i].ch == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
		}
	}
}

func (e *AppspaceStatusEvents) removeSubscriber(appspaceID domain.AppspaceID, ch chan<- domain.AppspaceStatusEvent) {
	// get a feeling you'll need a mutex to cover subscribers?
	for i, sub := range e.subscribers {
		if sub.appspaceID == appspaceID && sub.ch == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
		}
	}
}

////////////////////////////////////////
// Appspace Log events
type appspaceLogSubscriber struct {
	appspaceID domain.AppspaceID
	ch         chan<- domain.AppspaceLogEvent
}

// AppspaceLogEvents handles appspace pause and unpause events
type AppspaceLogEvents struct {
	subscribers []appspaceLogSubscriber
}

// Send sends an appspace status event
// Should this send arrays of events (for buffeing?)
// .. or should the channel do the buffering?
func (e *AppspaceLogEvents) Send(event domain.AppspaceLogEvent) {
	for _, sub := range e.subscribers {
		if sub.appspaceID == event.AppspaceID {
			sub.ch <- event
		}
	}
}

// Subscribe to an event to know when the status of an appspace has changed
func (e *AppspaceLogEvents) Subscribe(appspaceID domain.AppspaceID, ch chan<- domain.AppspaceLogEvent) {
	e.removeSubscriber(appspaceID, ch)
	e.subscribers = append(e.subscribers, appspaceLogSubscriber{appspaceID, ch})
}

// Unsubscribe to the event
func (e *AppspaceLogEvents) Unsubscribe(appspaceID domain.AppspaceID, ch chan<- domain.AppspaceLogEvent) {
	e.removeSubscriber(appspaceID, ch)
}

// UnsubscribeChannel removes the channel from all subscriptions
func (e *AppspaceLogEvents) UnsubscribeChannel(ch chan<- domain.AppspaceLogEvent) {
	for i := len(e.subscribers) - 1; i >= 0; i-- {
		if e.subscribers[i].ch == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
		}
	}
}

func (e *AppspaceLogEvents) removeSubscriber(appspaceID domain.AppspaceID, ch chan<- domain.AppspaceLogEvent) {
	// get a feeling you'll need a mutex to cover subscribers?
	for i, sub := range e.subscribers {
		if sub.appspaceID == appspaceID && sub.ch == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
		}
	}
}

//////////////////////////////////////////
// Appspace Route Event
// TODO: Shouldn't subscribers be for specific appspaces?

// AppspaceRouteHitEvents handles appspace pause and unpause events
type AppspaceRouteHitEvents struct {
	subscribers []chan<- *domain.AppspaceRouteHitEvent
}

// Send sends an appspace paused or unpaused event
// Event's timestamp is set if needed
func (e *AppspaceRouteHitEvents) Send(routeEvent *domain.AppspaceRouteHitEvent) {
	if routeEvent.Timestamp.IsZero() {
		routeEvent.Timestamp = time.Now()
	}
	for _, ch := range e.subscribers {
		ch <- routeEvent
	}
}

// Subscribe to an event for when an appspace is paused or unpaused
func (e *AppspaceRouteHitEvents) Subscribe(ch chan<- *domain.AppspaceRouteHitEvent) {
	e.removeSubscriber(ch)
	e.subscribers = append(e.subscribers, ch)
}

// Unsubscribe to an event for when an appspace is paused or unpaused
func (e *AppspaceRouteHitEvents) Unsubscribe(ch chan<- *domain.AppspaceRouteHitEvent) {
	e.removeSubscriber(ch)
}

func (e *AppspaceRouteHitEvents) removeSubscriber(ch chan<- *domain.AppspaceRouteHitEvent) {
	// get a feeling you'll need a mutex to cover subscribers?
	for i, c := range e.subscribers {
		if c == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
		}
	}
}

///////////////////////////////
// App Version data change event

// AppVersionEvents forwards events about changes to an app version's metadata
// This is only useful in ds-dev.
type AppVersionEvents struct {
	subscribers []chan<- domain.AppID
}

// Send sends an appspace paused or unpaused event
func (e *AppVersionEvents) Send(appID domain.AppID) {
	for _, ch := range e.subscribers {
		ch <- appID
	}
}

// Subscribe to an event for when an appspace is paused or unpaused
func (e *AppVersionEvents) Subscribe(ch chan<- domain.AppID) {
	e.removeSubscriber(ch)
	e.subscribers = append(e.subscribers, ch)
}

// Unsubscribe to an event for when an appspace is paused or unpaused
func (e *AppVersionEvents) Unsubscribe(ch chan<- domain.AppID) {
	e.removeSubscriber(ch)
}

func (e *AppVersionEvents) removeSubscriber(ch chan<- domain.AppID) {
	// get a feeling you'll need a mutex to cover subscribers?
	for i, c := range e.subscribers {
		if c == ch {
			e.subscribers[i] = e.subscribers[len(e.subscribers)-1]
			e.subscribers = e.subscribers[:len(e.subscribers)-1]
		}
	}
}
