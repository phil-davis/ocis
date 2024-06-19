package service

import (
	"embed"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	revactx "github.com/cs3org/reva/v2/pkg/ctx"
	"github.com/cs3org/reva/v2/pkg/events"
	"github.com/cs3org/reva/v2/pkg/storagespace"
	"github.com/cs3org/reva/v2/pkg/utils"
	libregraph "github.com/owncloud/libre-graph-api-go"

	"github.com/owncloud/ocis/v2/ocis-pkg/ast"
	"github.com/owncloud/ocis/v2/ocis-pkg/kql"
	"github.com/owncloud/ocis/v2/ocis-pkg/l10n"
	ehmsg "github.com/owncloud/ocis/v2/protogen/gen/ocis/messages/eventhistory/v0"
	ehsvc "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/eventhistory/v0"
)

var (
	//go:embed l10n/locale
	_localeFS embed.FS

	// subfolder where the translation files are stored
	_localeSubPath = "l10n/locale"

	// domain of the activitylog service (transifex)
	_domain = "activitylog"
)

// ServeHTTP implements the http.Handler interface.
func (s *ActivitylogService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// HandleGetItemActivities handles the request to get the activities of an item.
func (s *ActivitylogService) HandleGetItemActivities(w http.ResponseWriter, r *http.Request) {
	activeUser, ok := revactx.ContextGetUser(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rid, limit, rawActivityAccepted, activityAccepted, err := s.getFilters(r.URL.Query().Get("kql"))
	if err != nil {
		s.log.Info().Str("query", r.URL.Query().Get("kql")).Err(err).Msg("error getting filters")
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	raw, err := s.Activities(rid)
	if err != nil {
		s.log.Error().Err(err).Msg("error getting activities")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ids := make([]string, 0, len(raw))
	toDelete := make(map[string]struct{}, len(raw))
	for _, a := range raw {
		if !rawActivityAccepted(a) {
			continue
		}
		ids = append(ids, a.EventID)
		toDelete[a.EventID] = struct{}{}
	}

	evRes, err := s.evHistory.GetEvents(r.Context(), &ehsvc.GetEventsRequest{Ids: ids})
	if err != nil {
		s.log.Error().Err(err).Msg("error getting events")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var resp GetActivitiesResponse
	for _, e := range evRes.GetEvents() {
		delete(toDelete, e.GetId())

		if limit != 0 && len(resp.Activities) >= limit {
			continue
		}

		if !activityAccepted(e) {
			continue
		}

		var (
			message string
			res     Resource
			act     Actor
			ts      libregraph.ActivityTimes
		)

		switch ev := s.unwrapEvent(e).(type) {
		case nil:
			// error already logged in unwrapEvent
			continue
		case events.UploadReady:
			message = MessageResourceCreated
			res, act, ts, err = s.ResponseData(ev.FileRef, ev.ExecutingUser.GetId(), ev.ExecutingUser.GetDisplayName(), utils.TSToTime(ev.Timestamp))
		case events.FileTouched:
			message = MessageResourceCreated
			res, act, ts, err = s.ResponseData(ev.Ref, ev.Executant, "", utils.TSToTime(ev.Timestamp))
		case events.ContainerCreated:
			message = MessageResourceCreated
			res, act, ts, err = s.ResponseData(ev.Ref, ev.Executant, "", utils.TSToTime(ev.Timestamp))
		case events.ItemTrashed:
			message = MessageResourceTrashed
			res, act, ts, err = s.ResponseData(ev.Ref, ev.Executant, "", utils.TSToTime(ev.Timestamp))
		case events.ItemPurged:
			message = MessageResourcePurged
			res, act, ts, err = s.ResponseData(ev.Ref, ev.Executant, "", utils.TSToTime(ev.Timestamp))
		case events.ItemMoved:
			message = MessageResourceMoved
			res, act, ts, err = s.ResponseData(ev.Ref, ev.Executant, "", utils.TSToTime(ev.Timestamp))
		case events.ShareCreated:
			message = MessageShareCreated
			res, act, ts, err = s.ResponseData(toRef(ev.ItemID), ev.Executant, "", utils.TSToTime(ev.CTime))
		case events.ShareUpdated:
			message = MessageShareUpdated
			res, act, ts, err = s.ResponseData(toRef(ev.ItemID), ev.Executant, "", utils.TSToTime(ev.MTime))
		case events.ShareRemoved:
			message = MessageShareDeleted
			res, act, ts, err = s.ResponseData(toRef(ev.ItemID), ev.Executant, "", ev.Timestamp)
		case events.LinkCreated:
			message = MessageLinkCreated
			res, act, ts, err = s.ResponseData(toRef(ev.ItemID), ev.Executant, "", utils.TSToTime(ev.CTime))
		case events.LinkUpdated:
			message = MessageLinkUpdated
			res, act, ts, err = s.ResponseData(toRef(ev.ItemID), ev.Executant, "", utils.TSToTime(ev.CTime))
		case events.LinkRemoved:
			message = MessageLinkDeleted
			res, act, ts, err = s.ResponseData(toRef(ev.ItemID), ev.Executant, "", utils.TSToTime(ev.Timestamp))
		case events.SpaceShared:
			message = MessageSpaceShared
			res, act, ts, err = s.ResponseData(sToRef(ev.ID), ev.Executant, "", ev.Timestamp)
		case events.SpaceShareUpdated:
			message = MessageSpaceShareUpdated
			res, act, ts, err = s.ResponseData(sToRef(ev.ID), ev.Executant, "", ev.Timestamp)
		case events.SpaceUnshared:
			message = MessageSpaceUnshared
			res, act, ts, err = s.ResponseData(sToRef(ev.ID), ev.Executant, "", ev.Timestamp)
		}

		if err != nil {
			s.log.Error().Err(err).Msg("error getting response data")
			continue
		}

		// FIXME: configurable default locale?
		loc := l10n.MustGetUserLocale(r.Context(), activeUser.GetId().GetOpaqueId(), r.Header.Get(l10n.HeaderAcceptLanguage), s.valService)
		t := l10n.NewTranslatorFromCommonConfig("en", _domain, "", _localeFS, _localeSubPath)

		resp.Activities = append(resp.Activities, NewActivity(t.Translate(message, loc), res, act, ts, e.GetId()))
	}

	// delete activities in separate go routine
	if len(toDelete) > 0 {
		go func() {
			err := s.RemoveActivities(rid, toDelete)
			if err != nil {
				s.log.Error().Err(err).Msg("error removing activities")
			}
		}()
	}

	b, err := json.Marshal(resp)
	if err != nil {
		s.log.Error().Err(err).Msg("error marshalling activities")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		s.log.Error().Err(err).Msg("error writing response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *ActivitylogService) unwrapEvent(e *ehmsg.Event) interface{} {
	etype, ok := s.registeredEvents[e.GetType()]
	if !ok {
		s.log.Error().Str("eventid", e.GetId()).Str("eventtype", e.GetType()).Msg("event not registered")
		return nil
	}

	einterface, err := etype.Unmarshal(e.GetEvent())
	if err != nil {
		s.log.Error().Str("eventid", e.GetId()).Str("eventtype", e.GetType()).Msg("failed to umarshal event")
		return nil
	}

	return einterface
}

func (s *ActivitylogService) getFilters(query string) (*provider.ResourceId, int, func(RawActivity) bool, func(*ehmsg.Event) bool, error) {
	qast, err := kql.Builder{}.Build(query)
	if err != nil {
		return nil, 0, nil, nil, err
	}

	prefilters := make([]func(RawActivity) bool, 0)
	postfilters := make([]func(*ehmsg.Event) bool, 0)

	var (
		itemID string
		limit  int
	)

	for _, n := range qast.Nodes {
		switch v := n.(type) {
		case *ast.StringNode:
			switch strings.ToLower(v.Key) {
			case "itemid":
				itemID = v.Value
			case "depth":
				depth, err := strconv.Atoi(v.Value)
				if err != nil {
					return nil, limit, nil, nil, err
				}

				prefilters = append(prefilters, func(a RawActivity) bool {
					return a.Depth <= depth
				})
			case "limit":
				l, err := strconv.Atoi(v.Value)
				if err != nil {
					return nil, limit, nil, nil, err
				}

				limit = l
			}
		case *ast.DateTimeNode:
			switch v.Operator.Value {
			case "<", "<=":
				prefilters = append(prefilters, func(a RawActivity) bool {
					return a.Timestamp.Before(v.Value)
				})
			case ">", ">=":
				prefilters = append(prefilters, func(a RawActivity) bool {
					return a.Timestamp.After(v.Value)
				})
			}
		case *ast.OperatorNode:
			if v.Value != "AND" {
				return nil, limit, nil, nil, errors.New("only AND operator is supported")
			}
		}
	}

	rid, err := storagespace.ParseID(itemID)
	if err != nil {
		return nil, limit, nil, nil, err
	}
	pref := func(a RawActivity) bool {
		for _, f := range prefilters {
			if !f(a) {
				return false
			}
		}
		return true
	}
	postf := func(e *ehmsg.Event) bool {
		for _, f := range postfilters {
			if !f(e) {
				return false
			}
		}
		return true
	}
	return &rid, limit, pref, postf, nil
}
