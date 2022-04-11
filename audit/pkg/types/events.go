package types

import (
	"github.com/cs3org/reva/v2/pkg/events"
)

// RegisteredEvents returns the events the service is registered for
func RegisteredEvents() []events.Unmarshaller {
	return []events.Unmarshaller{
		events.ShareCreated{},
		events.ShareUpdated{},
		events.LinkCreated{},
		events.LinkUpdated{},
		events.ShareRemoved{},
		events.LinkRemoved{},
		events.ReceivedShareUpdated{},
		events.LinkAccessed{},
		events.LinkAccessFailed{},
		events.FileUploaded{},
		events.FileDownloaded{},
		events.ItemTrashed{},
		events.ItemMoved{},
		events.ItemPurged{},
		events.ItemRestored{},
		events.FileVersionRestored{},
		events.SpaceCreated{},
		events.SpaceRenamed{},
		events.SpaceEnabled{},
		events.SpaceDisabled{},
		events.SpaceDeleted{},
		events.UserCreated{},
		events.UserDeleted{},
		events.UserFeatureChanged{},
		events.GroupCreated{},
		events.GroupDeleted{},
		events.GroupMemberAdded{},
		events.GroupMemberRemoved{},
	}
}
