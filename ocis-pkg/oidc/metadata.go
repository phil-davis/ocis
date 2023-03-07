package oidc

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/owncloud/ocis/v2/ocis-pkg/log"
)

const wellknownPath = "/.well-known/openid-configuration"

// The ProviderMetadata describes an idp.
// see https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata
type ProviderMetadata struct {
	AuthorizationEndpoint string `json:"authorization_endpoint,omitempty"`
	//claims_parameter_supported
	ClaimsSupported []string `json:"claims_supported,omitempty"`
	//grant_types_supported
	IDTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported,omitempty"`
	Issuer                           string   `json:"issuer,omitempty"`
	JwksURI                          string   `json:"jwks_uri,omitempty"`
	//registration_endpoint
	//request_object_signing_alg_values_supported
	//request_parameter_supported
	//request_uri_parameter_supported
	//require_request_uri_registration
	//response_modes_supported
	ResponseTypesSupported []string `json:"response_types_supported,omitempty"`
	ScopesSupported        []string `json:"scopes_supported,omitempty"`
	SubjectTypesSupported  []string `json:"subject_types_supported,omitempty"`
	TokenEndpoint          string   `json:"token_endpoint,omitempty"`
	//token_endpoint_auth_methods_supported
	//token_endpoint_auth_signing_alg_values_supported
	UserinfoEndpoint string `json:"userinfo_endpoint,omitempty"`
	//userinfo_signing_alg_values_supported
	//code_challenge_methods_supported
	IntrospectionEndpoint string `json:"introspection_endpoint,omitempty"`
	//introspection_endpoint_auth_methods_supported
	//introspection_endpoint_auth_signing_alg_values_supported
	RevocationEndpoint string `json:"revocation_endpoint,omitempty"`
	//revocation_endpoint_auth_methods_supported
	//revocation_endpoint_auth_signing_alg_values_supported
	//id_token_encryption_alg_values_supported
	//id_token_encryption_enc_values_supported
	//userinfo_encryption_alg_values_supported
	//userinfo_encryption_enc_values_supported
	//request_object_encryption_alg_values_supported
	//request_object_encryption_enc_values_supported
	CheckSessionIframe string `json:"check_session_iframe,omitempty"`
	EndSessionEndpoint string `json:"end_session_endpoint,omitempty"`
	//claim_types_supported
}

func GetIDPMetadata(logger log.Logger, client *http.Client, idpURI string) (ProviderMetadata, error) {
	wellknownURI := strings.TrimSuffix(idpURI, "/") + wellknownPath

	resp, err := client.Get(wellknownURI)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to set request for .well-known/openid-configuration")
		return ProviderMetadata{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error().Err(err).Msg("unable to read discovery response body")
		return ProviderMetadata{}, err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Error().Str("status", resp.Status).Str("body", string(body)).Msg("error requesting openid-configuration")
		return ProviderMetadata{}, err
	}

	var oidcMetadata ProviderMetadata
	err = json.Unmarshal(body, &oidcMetadata)
	if err != nil {
		logger.Error().Err(err).Msg("failed to decode provider openid-configuration")
		return ProviderMetadata{}, err
	}
	return oidcMetadata, nil
}
