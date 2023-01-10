package identity

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-ldap/ldap/v3"
	"github.com/gofrs/uuid"
	libregraph "github.com/owncloud/libre-graph-api-go"
	oldap "github.com/owncloud/ocis/v2/ocis-pkg/ldap"
	"github.com/owncloud/ocis/v2/services/graph/pkg/config"
	"github.com/owncloud/ocis/v2/services/graph/pkg/service/v0/errorcode"
)

type educationConfig struct {
	schoolBaseDN      string
	schoolFilter      string
	schoolObjectClass string
	schoolScope       int
	// memberOfSchoolAttribute defines the AttributeType on the user/group objects
	// which contains the school Id to which the user/group is assigned
	memberOfSchoolAttribute string
	schoolAttributeMap      schoolAttributeMap

	userObjectClass  string
	userAttributeMap educationUserAttributeMap
}

type schoolAttributeMap struct {
	displayName  string
	schoolNumber string
	id           string
}

type schoolUpdateOperation uint8

const (
	tooManyValues schoolUpdateOperation = iota
	schoolUnchanged
	displayNameUpdated
	schoolNumberUpdated
)

func defaultEducationConfig() educationConfig {
	return educationConfig{
		schoolObjectClass:       "ocEducationSchool",
		schoolScope:             ldap.ScopeWholeSubtree,
		memberOfSchoolAttribute: "ocMemberOfSchool",
		schoolAttributeMap:      newSchoolAttributeMap(),

		userObjectClass:  "ocEducationUser",
		userAttributeMap: newEducationUserAttributeMap(),
	}
}

func newEducationConfig(config config.LDAP) (educationConfig, error) {
	if config.EducationResourcesEnabled {
		var err error
		eduCfg := defaultEducationConfig()
		eduCfg.schoolBaseDN = config.EducationConfig.SchoolBaseDN
		if config.EducationConfig.SchoolSearchScope != "" {
			if eduCfg.schoolScope, err = stringToScope(config.EducationConfig.SchoolSearchScope); err != nil {
				return educationConfig{}, fmt.Errorf("error configuring school search scope: %w", err)
			}
		}
		if config.EducationConfig.SchoolFilter != "" {
			eduCfg.schoolFilter = config.EducationConfig.SchoolFilter
		}
		if config.EducationConfig.SchoolObjectClass != "" {
			eduCfg.schoolObjectClass = config.EducationConfig.SchoolObjectClass
		}

		// Attribute mapping config
		if config.EducationConfig.SchoolNameAttribute != "" {
			eduCfg.schoolAttributeMap.displayName = config.EducationConfig.SchoolNameAttribute
		}
		if config.EducationConfig.SchoolNumberAttribute != "" {
			eduCfg.schoolAttributeMap.schoolNumber = config.EducationConfig.SchoolNumberAttribute
		}
		if config.EducationConfig.SchoolIDAttribute != "" {
			eduCfg.schoolAttributeMap.id = config.EducationConfig.SchoolIDAttribute
		}

		return eduCfg, nil
	}
	return educationConfig{}, nil
}

func newSchoolAttributeMap() schoolAttributeMap {
	return schoolAttributeMap{
		displayName:  "ou",
		schoolNumber: "ocEducationSchoolNumber",
		id:           "owncloudUUID",
	}
}

// CreateEducationSchool creates the supplied school in the identity backend.
func (i *LDAP) CreateEducationSchool(ctx context.Context, school libregraph.EducationSchool) (*libregraph.EducationSchool, error) {
	logger := i.logger.SubloggerWithRequestID(ctx)
	logger.Debug().Str("backend", "ldap").Msg("CreateEducationSchool")
	if !i.writeEnabled {
		return nil, ErrReadOnly
	}

	// Here we should verify that the school number is not already used

	dn := fmt.Sprintf("%s=%s,%s",
		i.educationConfig.schoolAttributeMap.displayName,
		oldap.EscapeDNAttributeValue(school.GetDisplayName()),
		i.educationConfig.schoolBaseDN,
	)
	ar := ldap.NewAddRequest(dn, nil)
	ar.Attribute(i.educationConfig.schoolAttributeMap.displayName, []string{school.GetDisplayName()})
	ar.Attribute(i.educationConfig.schoolAttributeMap.schoolNumber, []string{school.GetSchoolNumber()})
	if !i.useServerUUID {
		ar.Attribute(i.educationConfig.schoolAttributeMap.id, []string{uuid.Must(uuid.NewV4()).String()})
	}
	objectClasses := []string{"organizationalUnit", i.educationConfig.schoolObjectClass, "top"}
	ar.Attribute("objectClass", objectClasses)

	if err := i.conn.Add(ar); err != nil {
		var lerr *ldap.Error
		logger.Debug().Err(err).Msg("error adding school")
		if errors.As(err, &lerr) {
			if lerr.ResultCode == ldap.LDAPResultEntryAlreadyExists {
				err = errorcode.New(errorcode.NameAlreadyExists, lerr.Error())
			}
		}
		return nil, err
	}

	// Read	back school from LDAP to get the generated UUID
	e, err := i.getSchoolByDN(ar.DN)
	if err != nil {
		return nil, err
	}
	return i.createSchoolModelFromLDAP(e), nil
}

// UpdateEducationSchoolOperation contains the logic for which update operation to apply to a school
func (i *LDAP) UpdateEducationSchoolOperation(
	schoolUpdate libregraph.EducationSchool,
	currentSchool libregraph.EducationSchool,
) schoolUpdateOperation {
	providedDisplayName := schoolUpdate.GetDisplayName()
	schoolNumber := schoolUpdate.GetSchoolNumber()

	if providedDisplayName != "" && schoolNumber != "" {
		return tooManyValues
	}

	if providedDisplayName != "" && providedDisplayName != currentSchool.GetDisplayName() {
		return displayNameUpdated
	}

	if schoolNumber != "" && schoolNumber != currentSchool.GetSchoolNumber() {
		return schoolNumberUpdated
	}

	return schoolUnchanged
}

// updateDisplayName updates the school OU in the identity backend
func (i *LDAP) updateDisplayName(ctx context.Context, dn string, providedDisplayName string) error {
	logger := i.logger.SubloggerWithRequestID(ctx)
	newDisplayName := fmt.Sprintf(
		"%s=%s",
		i.educationConfig.schoolAttributeMap.displayName,
		providedDisplayName,
	)

	mrdn := ldap.NewModifyDNRequest(dn, newDisplayName, true, "")
	i.logger.Debug().Str("backend", "ldap").
		Str("dn", mrdn.DN).
		Str("newrdn", mrdn.NewRDN).
		Msg("updateDisplayName")

	if err := i.conn.ModifyDN(mrdn); err != nil {
		var lerr *ldap.Error
		logger.Debug().Err(err).Msg("error updating school name")
		if errors.As(err, &lerr) {
			if lerr.ResultCode == ldap.LDAPResultEntryAlreadyExists {
				err = errorcode.New(errorcode.NameAlreadyExists, lerr.Error())
			}
		}
		return err
	}

	return nil
}

// updateSchoolNumber checks if a school number is already taken, and if not updates the school number
func (i *LDAP) updateSchoolNumber(ctx context.Context, dn string, schoolNumber string) error {
	logger := i.logger.SubloggerWithRequestID(ctx)
	_, err := i.getSchoolByNumberOrID(schoolNumber)
	if err == nil {
		errmsg := fmt.Sprintf("school number '%s' already exists", schoolNumber)
		err = fmt.Errorf(errmsg)
		return err
	}

	mr := ldap.NewModifyRequest(dn, nil)
	mr.Replace(i.educationConfig.schoolAttributeMap.schoolNumber, []string{schoolNumber})

	if err := i.conn.Modify(mr); err != nil {
		var lerr *ldap.Error
		logger.Debug().Err(err).Msg("error updating school number")
		if errors.As(err, &lerr) {
			if lerr.ResultCode == ldap.LDAPResultEntryAlreadyExists {
				err = errorcode.New(errorcode.NameAlreadyExists, lerr.Error())
			}
		}
		return err
	}

	return nil
}

// UpdateEducationSchool updates the supplied school in the identity backend
func (i *LDAP) UpdateEducationSchool(ctx context.Context, numberOrID string, school libregraph.EducationSchool) (*libregraph.EducationSchool, error) {
	logger := i.logger.SubloggerWithRequestID(ctx)
	logger.Debug().Str("backend", "ldap").Msg("UpdateEducationSchool")
	if !i.writeEnabled {
		return nil, ErrReadOnly
	}

	providedDisplayName := school.GetDisplayName()
	schoolNumber := school.GetSchoolNumber()

	if providedDisplayName != "" && schoolNumber != "" {
		return nil, fmt.Errorf("school name and school number cannot be updated in the same request")
	}

	e, err := i.getSchoolByNumberOrID(numberOrID)
	if err != nil {
		return nil, err
	}

	currentSchool := i.createSchoolModelFromLDAP(e)
	switch i.UpdateEducationSchoolOperation(school, *currentSchool) {
	case tooManyValues:
		return nil, fmt.Errorf("school name and school number cannot be updated in the same request")
	case schoolUnchanged:
		logger.Debug().Str("backend", "ldap").Msg("UpdateEducationSchool: Nothing changed")
		return i.createSchoolModelFromLDAP(e), nil
	case displayNameUpdated:
		if err := i.updateDisplayName(ctx, e.DN, providedDisplayName); err != nil {
			return nil, err
		}
	case schoolNumberUpdated:
		if err := i.updateSchoolNumber(ctx, e.DN, schoolNumber); err != nil {
			return nil, err
		}
	}

	// Read	back school from LDAP
	e, err = i.getSchoolByNumberOrID(i.getID(e))
	if err != nil {
		return nil, err
	}
	return i.createSchoolModelFromLDAP(e), nil
}

// DeleteEducationSchool deletes a given school, identified by id
func (i *LDAP) DeleteEducationSchool(ctx context.Context, id string) error {
	logger := i.logger.SubloggerWithRequestID(ctx)
	logger.Debug().Str("backend", "ldap").Msg("DeleteEducationSchool")
	if !i.writeEnabled {
		return ErrReadOnly
	}
	e, err := i.getSchoolByNumberOrID(id)
	if err != nil {
		return err
	}

	dr := ldap.DelRequest{DN: e.DN}
	if err = i.conn.Del(&dr); err != nil {
		return err
	}

	// TODO update any users that are member of this school
	return nil
}

// GetEducationSchool implements the EducationBackend interface for the LDAP backend.
func (i *LDAP) GetEducationSchool(ctx context.Context, numberOrID string, queryParam url.Values) (*libregraph.EducationSchool, error) {
	logger := i.logger.SubloggerWithRequestID(ctx)
	logger.Debug().Str("backend", "ldap").Msg("GetEducationSchool")
	e, err := i.getSchoolByNumberOrID(numberOrID)
	if err != nil {
		return nil, err
	}

	return i.createSchoolModelFromLDAP(e), nil
}

// GetEducationSchools implements the EducationBackend interface for the LDAP backend.
func (i *LDAP) GetEducationSchools(ctx context.Context, queryParam url.Values) ([]*libregraph.EducationSchool, error) {
	var filter string
	filter = fmt.Sprintf("(objectClass=%s)", i.educationConfig.schoolObjectClass)

	if i.educationConfig.schoolFilter != "" {
		filter = fmt.Sprintf("(&%s%s)", i.educationConfig.schoolFilter, filter)
	}

	searchRequest := ldap.NewSearchRequest(
		i.educationConfig.schoolBaseDN,
		i.educationConfig.schoolScope,
		ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{
			i.educationConfig.schoolAttributeMap.displayName,
			i.educationConfig.schoolAttributeMap.id,
			i.educationConfig.schoolAttributeMap.schoolNumber,
		},
		nil,
	)
	i.logger.Debug().Str("backend", "ldap").
		Str("base", searchRequest.BaseDN).
		Str("filter", searchRequest.Filter).
		Int("scope", searchRequest.Scope).
		Int("sizelimit", searchRequest.SizeLimit).
		Interface("attributes", searchRequest.Attributes).
		Msg("GetEducationSchools")
	res, err := i.conn.Search(searchRequest)
	if err != nil {
		return nil, errorcode.New(errorcode.ItemNotFound, err.Error())
	}

	schools := make([]*libregraph.EducationSchool, 0, len(res.Entries))
	for _, e := range res.Entries {
		school := i.createSchoolModelFromLDAP(e)
		// Skip invalid LDAP entries
		if school == nil {
			continue
		}
		schools = append(schools, school)
	}
	return schools, nil
}

// GetEducationSchoolUsers implements the EducationBackend interface for the LDAP backend.
func (i *LDAP) GetEducationSchoolUsers(ctx context.Context, id string) ([]*libregraph.EducationUser, error) {
	logger := i.logger.SubloggerWithRequestID(ctx)
	logger.Debug().Str("backend", "ldap").Msg("GetEducationSchoolUsers")

	schoolEntry, err := i.getSchoolByNumberOrID(id)
	if err != nil {
		return nil, err
	}

	if schoolEntry == nil {
		return nil, ErrNotFound
	}
	id = ldap.EscapeFilter(id)
	idFilter := fmt.Sprintf("(%s=%s)", i.educationConfig.memberOfSchoolAttribute, id)
	userFilter := fmt.Sprintf("(&%s(objectClass=%s)%s)", i.userFilter, i.educationConfig.userObjectClass, idFilter)

	searchRequest := ldap.NewSearchRequest(
		i.userBaseDN,
		i.userScope,
		ldap.NeverDerefAliases, 0, 0, false,
		userFilter,
		i.getEducationUserAttrTypes(),
		nil,
	)
	logger.Debug().Str("backend", "ldap").
		Str("base", searchRequest.BaseDN).
		Str("filter", searchRequest.Filter).
		Int("scope", searchRequest.Scope).
		Int("sizelimit", searchRequest.SizeLimit).
		Interface("attributes", searchRequest.Attributes).
		Msg("GetEducationUsers")
	res, err := i.conn.Search(searchRequest)
	if err != nil {
		return nil, errorcode.New(errorcode.ItemNotFound, err.Error())
	}

	users := make([]*libregraph.EducationUser, 0, len(res.Entries))

	for _, e := range res.Entries {
		u := i.createEducationUserModelFromLDAP(e)
		// Skip invalid LDAP users
		if u == nil {
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

// AddUsersToEducationSchool adds new members (reference by a slice of IDs) to supplied school in the identity backend.
func (i *LDAP) AddUsersToEducationSchool(ctx context.Context, schoolID string, memberIDs []string) error {
	logger := i.logger.SubloggerWithRequestID(ctx)
	logger.Debug().Str("backend", "ldap").Msg("AddUsersToEducationSchool")

	schoolEntry, err := i.getSchoolByNumberOrID(schoolID)
	if err != nil {
		return err
	}

	if schoolEntry == nil {
		return ErrNotFound
	}

	userEntries := make([]*ldap.Entry, 0, len(memberIDs))
	for _, memberID := range memberIDs {
		user, err := i.getEducationUserByNameOrID(memberID)
		if err != nil {
			i.logger.Warn().Str("userid", memberID).Msg("User does not exist")
			return err
		}
		userEntries = append(userEntries, user)
	}

	for _, userEntry := range userEntries {
		currentSchools := userEntry.GetEqualFoldAttributeValues(i.educationConfig.memberOfSchoolAttribute)
		found := false
		for _, currentSchool := range currentSchools {
			if currentSchool == schoolID {
				found = true
				break
			}
		}
		if !found {
			mr := ldap.ModifyRequest{DN: userEntry.DN}
			mr.Add(i.educationConfig.memberOfSchoolAttribute, []string{schoolID})
			if err := i.conn.Modify(&mr); err != nil {
				return err
			}
		}
	}

	return nil
}

// RemoveUserFromEducationSchool removes a single member (by ID) from a school
func (i *LDAP) RemoveUserFromEducationSchool(ctx context.Context, schoolID string, memberID string) error {
	logger := i.logger.SubloggerWithRequestID(ctx)
	logger.Debug().Str("backend", "ldap").Msg("RemoveUserFromEducationSchool")

	schoolEntry, err := i.getSchoolByNumberOrID(schoolID)
	if err != nil {
		return err
	}

	if schoolEntry == nil {
		return ErrNotFound
	}
	user, err := i.getEducationUserByNameOrID(memberID)
	if err != nil {
		i.logger.Warn().Str("userid", memberID).Msg("User does not exist")
		return err
	}
	currentSchools := user.GetEqualFoldAttributeValues(i.educationConfig.memberOfSchoolAttribute)
	for _, currentSchool := range currentSchools {
		if currentSchool == schoolID {
			mr := ldap.ModifyRequest{DN: user.DN}
			mr.Delete(i.educationConfig.memberOfSchoolAttribute, []string{schoolID})
			if err := i.conn.Modify(&mr); err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func (i *LDAP) getSchoolByDN(dn string) (*ldap.Entry, error) {
	attrs := []string{
		i.educationConfig.schoolAttributeMap.displayName,
		i.educationConfig.schoolAttributeMap.id,
		i.educationConfig.schoolAttributeMap.schoolNumber,
	}
	filter := fmt.Sprintf("(objectClass=%s)", i.educationConfig.schoolObjectClass)

	if i.educationConfig.schoolFilter != "" {
		filter = fmt.Sprintf("(&%s(%s))", filter, i.educationConfig.schoolFilter)
	}
	return i.getEntryByDN(dn, attrs, filter)
}

func (i *LDAP) getSchoolByNumberOrID(numberOrID string) (*ldap.Entry, error) {
	numberOrID = ldap.EscapeFilter(numberOrID)
	filter := fmt.Sprintf(
		"(|(%s=%s)(%s=%s))",
		i.educationConfig.schoolAttributeMap.id,
		numberOrID,
		i.educationConfig.schoolAttributeMap.schoolNumber,
		numberOrID,
	)
	return i.getSchoolByFilter(filter)
}

func (i *LDAP) getSchoolByFilter(filter string) (*ldap.Entry, error) {
	filter = fmt.Sprintf("(&%s(objectClass=%s)%s)",
		i.educationConfig.schoolFilter,
		i.educationConfig.schoolObjectClass,
		filter,
	)
	searchRequest := ldap.NewSearchRequest(
		i.educationConfig.schoolBaseDN,
		i.educationConfig.schoolScope,
		ldap.NeverDerefAliases, 1, 0, false,
		filter,
		[]string{
			i.educationConfig.schoolAttributeMap.displayName,
			i.educationConfig.schoolAttributeMap.id,
			i.educationConfig.schoolAttributeMap.schoolNumber,
		},
		nil,
	)
	i.logger.Debug().Str("backend", "ldap").
		Str("base", searchRequest.BaseDN).
		Str("filter", searchRequest.Filter).
		Int("scope", searchRequest.Scope).
		Int("sizelimit", searchRequest.SizeLimit).
		Interface("attributes", searchRequest.Attributes).
		Msg("getSchoolByFilter")
	res, err := i.conn.Search(searchRequest)
	if err != nil {
		var errmsg string
		if lerr, ok := err.(*ldap.Error); ok {
			if lerr.ResultCode == ldap.LDAPResultSizeLimitExceeded {
				errmsg = fmt.Sprintf("too many results searching for school '%s'", filter)
				i.logger.Debug().Str("backend", "ldap").Err(lerr).
					Str("schoolfilter", filter).Msg("too many results searching for school")
			}
		}
		return nil, errorcode.New(errorcode.ItemNotFound, errmsg)
	}
	if len(res.Entries) == 0 {
		return nil, ErrNotFound
	}

	return res.Entries[0], nil
}

func (i *LDAP) createSchoolModelFromLDAP(e *ldap.Entry) *libregraph.EducationSchool {
	if e == nil {
		return nil
	}

	displayName := i.getDisplayName(e)
	id := i.getID(e)
	schoolNumber := i.getSchoolNumber(e)

	if id != "" && displayName != "" && schoolNumber != "" {
		school := libregraph.NewEducationSchool()
		school.SetDisplayName(displayName)
		school.SetSchoolNumber(schoolNumber)
		school.SetId(id)
		return school
	}
	i.logger.Warn().Str("dn", e.DN).Str("id", id).Str("displayName", displayName).Str("schoolNumber", schoolNumber).Msg("Invalid School. Missing required attribute")
	return nil
}

func (i *LDAP) getSchoolNumber(e *ldap.Entry) string {
	schoolNumber := e.GetEqualFoldAttributeValue(i.educationConfig.schoolAttributeMap.schoolNumber)
	return schoolNumber
}

func (i *LDAP) getID(e *ldap.Entry) string {
	id := e.GetEqualFoldAttributeValue(i.educationConfig.schoolAttributeMap.id)
	return id
}

func (i *LDAP) getDisplayName(e *ldap.Entry) string {
	displayName := e.GetEqualFoldAttributeValue(i.educationConfig.schoolAttributeMap.displayName)
	return displayName
}
