package config

import (
	"context"

	"github.com/owncloud/ocis/ocis-pkg/shared"
)

// Config combines all available configuration parts.
type Config struct {
	*shared.Commons `yaml:"-"`

	Service Service `yaml:"-"`

	Tracing *Tracing `yaml:"tracing,omitempty"`
	Log     *Log     `yaml:"log,omitempty"`
	Debug   Debug    `yaml:"debug,omitempty"`

	HTTP HTTP `yaml:"http,omitempty"`

	Reva         *Reva         `yaml:"reva,omitempty"`
	TokenManager *TokenManager `yaml:"token_manager,omitempty"`

	Spaces   Spaces   `yaml:"spaces,omitempty"`
	Identity Identity `yaml:"identity,omitempty"`
	Events   Events   `yaml:"events,omitempty"`

	Context context.Context `yaml:"-"`
}

type Spaces struct {
	WebDavBase                      string `yaml:"webdav_base,omitempty" env:"OCIS_URL;GRAPH_SPACES_WEBDAV_BASE"`
	WebDavPath                      string `yaml:"webdav_path,omitempty" env:"GRAPH_SPACES_WEBDAV_PATH"`
	DefaultQuota                    string `yaml:"default_quota,omitempty" env:"GRAPH_SPACES_DEFAULT_QUOTA"`
	Insecure                        bool   `yaml:"insecure,omitempty" env:"OCIS_INSECURE;GRAPH_SPACES_INSECURE"`
	ExtendedSpacePropertiesCacheTTL int    `yaml:"extended_space_properties_cache_ttl,omitempty" env:"GRAPH_SPACES_EXTENDED_SPACE_PROPERTIES_CACHE_TTL"`
}

type LDAP struct {
	URI           string `yaml:"uri,omitempty" env:"LDAP_URI;GRAPH_LDAP_URI"`
	Insecure      bool   `yaml:"insecure,omitempty" env:"OCIS_INSECURE;GRAPH_LDAP_INSECURE"`
	BindDN        string `yaml:"bind_dn,omitempty" env:"LDAP_BIND_DN;GRAPH_LDAP_BIND_DN"`
	BindPassword  string `yaml:"bind_password,omitempty" env:"LDAP_BIND_PASSWORD;GRAPH_LDAP_BIND_PASSWORD"`
	UseServerUUID bool   `yaml:"use_server_uuid,omitempty" env:"GRAPH_LDAP_SERVER_UUID"`
	WriteEnabled  bool   `yaml:"write_enabled,omitempty" env:"GRAPH_LDAP_SERVER_WRITE_ENABLED"`

	UserBaseDN               string `yaml:"user_base_dn,omitempty" env:"LDAP_USER_BASE_DN;GRAPH_LDAP_USER_BASE_DN"`
	UserSearchScope          string `yaml:"user_search_scope,omitempty" env:"LDAP_USER_SCOPE;GRAPH_LDAP_USER_SCOPE"`
	UserFilter               string `yaml:"user_filter,omitempty" env:"LDAP_USER_FILTER;GRAPH_LDAP_USER_FILTER"`
	UserObjectClass          string `yaml:"user_objectclass,omitempty" env:"LDAP_USER_OBJECTCLASS;GRAPH_LDAP_USER_OBJECTCLASS"`
	UserEmailAttribute       string `yaml:"user_mail_attribute,omitempty" env:"LDAP_USER_SCHEMA_MAIL;GRAPH_LDAP_USER_EMAIL_ATTRIBUTE"`
	UserDisplayNameAttribute string `yaml:"user_displayname_attribute,omitempty" env:"LDAP_USER_SCHEMA_DISPLAY_NAME;GRAPH_LDAP_USER_DISPLAYNAME_ATTRIBUTE"`
	UserNameAttribute        string `yaml:"user_name_attribute,omitempty" env:"LDAP_USER_SCHEMA_USERNAME;GRAPH_LDAP_USER_NAME_ATTRIBUTE"`
	UserIDAttribute          string `yaml:"user_id_attribute,omitempty" env:"LDAP_USER_SCHEMA_ID;GRAPH_LDAP_USER_UID_ATTRIBUTE"`

	GroupBaseDN        string `yaml:"group_base_dn,omitempty" env:"LDAP_GROUP_BASE_DN;GRAPH_LDAP_GROUP_BASE_DN"`
	GroupSearchScope   string `yaml:"group_search_scope,omitempty" env:"LDAP_GROUP_SCOPE;GRAPH_LDAP_GROUP_SEARCH_SCOPE"`
	GroupFilter        string `yaml:"group_filter,omitempty" env:"LDAP_GROUP_FILTER;GRAPH_LDAP_GROUP_FILTER"`
	GroupObjectClass   string `yaml:"group_objectclass,omitempty" env:"LDAP_GROUP_OBJECTCLASS;GRAPH_LDAP_GROUP_OBJECTCLASS"`
	GroupNameAttribute string `yaml:"group_name_attribute,omitempty" env:"LDAP_GROUP_SCHEMA_GROUPNAME;GRAPH_LDAP_GROUP_NAME_ATTRIBUTE"`
	GroupIDAttribute   string `yaml:"group_id_attribute,omitempty" env:"LDAP_GROUP_SCHEMA_ID;GRAPH_LDAP_GROUP_ID_ATTRIBUTE"`
}

type Identity struct {
	Backend string `yaml:"backend,omitempty" env:"GRAPH_IDENTITY_BACKEND"`
	LDAP    LDAP   `yaml:"ldap,omitempty"`
}

// Events combines the configuration options for the event bus.
type Events struct {
	Endpoint string `yaml:"events_endpoint,omitempty" env:"GRAPH_EVENTS_ENDPOINT" desc:"the address of the streaming service"`
	Cluster  string `yaml:"events_cluster,omitempty" env:"GRAPH_EVENTS_CLUSTER" desc:"the clusterID of the streaming service. Mandatory when using nats"`
}
