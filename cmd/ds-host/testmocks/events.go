package testmocks

import (
	"github.com/teleclimber/DropServer/cmd/ds-host/domain"
)

//go:generate mockgen -destination=events_mocks.go -package=testmocks github.com/teleclimber/DropServer/cmd/ds-host/testmocks AppspacePausedEvents,AppspaceStatusEvents

// xxx go:generate mockgen -source=$GOFILE
// ^^ the above fails with an internal error: nil Pkg imports which I have no idea how to fix.

// AppspacePausedEvents interface for mocking
type AppspacePausedEvents interface {
	Subscribe(ch chan<- domain.AppspacePausedEvent)
	Unsubscribe(ch chan<- domain.AppspacePausedEvent)
	Send(domain.AppspaceID, bool)
}

// AppspaceStatusEvents interface for mocking
type AppspaceStatusEvents interface {
	Subscribe(appspaceID domain.AppspaceID, ch chan<- domain.AppspaceStatusEvent)
	Unsubscribe(appspaceID domain.AppspaceID, ch chan<- domain.AppspaceStatusEvent)
	UnsubscribeChannel(chan<- domain.AppspaceStatusEvent)
	Send(domain.AppspaceID, domain.AppspaceStatusEvent)
}
