package oidc

import (
	"net/http"

	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/services/proxy/pkg/config"

	gOidc "github.com/coreos/go-oidc/v3/oidc"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	// HTTPClient to use for requests
	HTTPClient *http.Client
	// Logger to use for logging, must be set
	Logger log.Logger
	// The OpenID Connect Issuer URL
	OidcIssuer string
	// JWKSOptions to use when retrieving keys
	JWKSOptions config.JWKS
	// KeySet to use when verifiing signatures
	KeySet KeySet
	// AccessTokenVerifyMethod to use when verifying access tokens
	// TODO pass a function or interface to verify? an AccessTokenVerifier?
	AccessTokenVerifyMethod string
	// ClientID the client id to expect in tokens. If not set SkipClientIDCheck must be true
	// TODO also check in access token
	ClientID string
	// SkipClientIDCheck must be true if ClientID is empty
	SkipClientIDCheck bool
	// Config to use
	Config *gOidc.Config

	// ProviderMetadata to use
	ProviderMetadata *ProviderMetadata
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// WithLogger provides a function to set the openid connect issuer option.
func WithOidcIssuer(val string) Option {
	return func(o *Options) {
		o.OidcIssuer = val
	}
}

// WithLogger provides a function to set the logger option.
func WithLogger(val log.Logger) Option {
	return func(o *Options) {
		o.Logger = val
	}
}

// WithAccessTokenVerifyMethod provides a function to set the accessTokenVerifyMethod option.
func WithAccessTokenVerifyMethod(val string) Option {
	return func(o *Options) {
		o.AccessTokenVerifyMethod = val
	}
}

// WithHTTPClient provides a function to set the httpClient option.
func WithHTTPClient(val *http.Client) Option {
	return func(o *Options) {
		o.HTTPClient = val
	}
}

// WithJWKSOptions provides a function to set the jwksOptions option.
func WithJWKSOptions(val config.JWKS) Option {
	return func(o *Options) {
		o.JWKSOptions = val
	}
}

// WithKeySet provides a function to set the KeySet option.
func WithKeySet(val KeySet) Option {
	return func(o *Options) {
		o.KeySet = val
	}
}

// WithClientID provides a function to set the clientID option.
func WithClientID(val string) Option {
	return func(o *Options) {
		o.ClientID = val
	}
}

// WithSkipClientIDCheck provides a function to set the skipClientIDCheck option.
func WithSkipClientIDCheck(val bool) Option {
	return func(o *Options) {
		o.SkipClientIDCheck = val
	}
}

// WithSkipClientIDCheck provides a function to set the skipClientIDCheck option.
func WithConfig(val *gOidc.Config) Option {
	return func(o *Options) {
		o.Config = val
	}
}

// WithProviderMetadata provides a function to set the provider option.
func WithProviderMetadata(val *ProviderMetadata) Option {
	return func(o *Options) {
		o.ProviderMetadata = val
	}
}
