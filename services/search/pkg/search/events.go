package search

import (
	"time"

	user "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	"github.com/cs3org/reva/v2/pkg/events"
	"github.com/cs3org/reva/v2/pkg/storagespace"
	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/services/search/pkg/config"
)

// HandleEvents listens to the needed events,
// it handles the whole resource indexing livecycle.
func HandleEvents(s Searcher, bus events.Consumer, logger log.Logger, cfg *config.Config) error {
	evts := []events.Unmarshaller{
		events.ItemTrashed{},
		events.ItemRestored{},
		events.ItemMoved{},
		events.ContainerCreated{},
		events.FileTouched{},
		events.FileVersionRestored{},
		events.TagsAdded{},
		events.TagsRemoved{},
		events.SpaceRenamed{},
	}

	if cfg.Events.AsyncUploads {
		evts = append(evts, events.UploadReady{})
	} else {
		evts = append(evts, events.FileUploaded{})
	}

	ch, err := events.Consume(bus, "search", evts...)
	if err != nil {
		return err
	}

	if cfg.Events.NumConsumers == 0 {
		cfg.Events.NumConsumers = 1
	}

	getSpaceID := func(ref *provider.Reference) *provider.StorageSpaceId {
		return &provider.StorageSpaceId{
			OpaqueId: storagespace.FormatResourceID(
				provider.ResourceId{
					StorageId: ref.GetResourceId().GetStorageId(),
					SpaceId:   ref.GetResourceId().GetSpaceId(),
				},
			),
		}
	}

	getUser := func(users ...*user.UserId) *user.UserId {
		for _, u := range users {
			if u == nil {
				continue
			}

			return u
		}

		return nil
	}

	indexSpaceDebouncer := NewSpaceDebouncer(time.Duration(cfg.Events.DebounceDuration)*time.Millisecond, func(id *provider.StorageSpaceId, userID *user.UserId) {
		if err := s.IndexSpace(id, userID); err != nil {
			logger.Error().Err(err).Interface("spaceID", id).Interface("userID", userID).Msg("error while indexing a space")
		}
	})

	for i := 0; i < cfg.Events.NumConsumers; i++ {
		go func(s Searcher, ch <-chan events.Event) {
			for event := range ch {
				e := event
				go func() {
					logger.Debug().Interface("event", e).Msg("updating index")

					var err error

					switch ev := e.Event.(type) {
					case events.ItemTrashed:
						u := getUser(ev.SpaceOwner, ev.Executant)
						s.TrashItem(ev.ID)
						indexSpaceDebouncer.Debounce(getSpaceID(ev.Ref), u)
					case events.ItemMoved:
						u := getUser(ev.SpaceOwner, ev.Executant)
						s.MoveItem(ev.Ref, u)
						indexSpaceDebouncer.Debounce(getSpaceID(ev.Ref), getUser(ev.SpaceOwner, ev.Executant))
					case events.ItemRestored:
						u := getUser(ev.SpaceOwner, ev.Executant)
						s.RestoreItem(ev.Ref, u)
						indexSpaceDebouncer.Debounce(getSpaceID(ev.Ref), u)
					case events.ContainerCreated:
						indexSpaceDebouncer.Debounce(getSpaceID(ev.Ref), getUser(ev.SpaceOwner, ev.Executant))
					case events.FileTouched:
						indexSpaceDebouncer.Debounce(getSpaceID(ev.Ref), getUser(ev.SpaceOwner, ev.Executant))
					case events.FileVersionRestored:
						indexSpaceDebouncer.Debounce(getSpaceID(ev.Ref), getUser(ev.SpaceOwner, ev.Executant))
					case events.TagsAdded:
						s.UpsertItem(ev.Ref, ev.Executant)
					case events.TagsRemoved:
						s.UpsertItem(ev.Ref, ev.Executant)
					case events.FileUploaded:
						indexSpaceDebouncer.Debounce(getSpaceID(ev.Ref), getUser(ev.SpaceOwner, ev.Executant))
					case events.UploadReady:
						indexSpaceDebouncer.Debounce(getSpaceID(ev.FileRef), getUser(ev.SpaceOwner, ev.ExecutingUser.Id))
					case events.SpaceRenamed:
						indexSpaceDebouncer.Debounce(ev.ID, getUser(ev.Executant))
					}

					if err != nil {
						logger.Error().Err(err).Interface("event", e)
					}
				}()
			}
		}(
			s,
			ch,
		)
	}

	return nil
}
