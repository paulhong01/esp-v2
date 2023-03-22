// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"time"
)

const (
	// DefaultRootCAPaths is the default certs path.
	DefaultRootCAPaths = "/etc/ssl/certs/ca-certificates.crt"

	// ESPv2 custom http filters.

	// JwtPayloadMetadataName is the field name passed into metadata
	JwtPayloadMetadataName = "jwt_payloads"

	// Supported Http Methods.

	GET     = "GET"
	PUT     = "PUT"
	POST    = "POST"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	OPTIONS = "OPTIONS"
	CUSTOM  = "CUSTOM"

	// Rollout strategy

	FixedRolloutStrategy   = "fixed"
	ManagedRolloutStrategy = "managed"

	// Metadata suffix

	ConfigIDPath          = "/computeMetadata/v1/instance/attributes/endpoints-service-version"
	GAEServerSoftwarePath = "/computeMetadata/v1/instance/attributes/gae_server_software"
	KubeEnvPath           = "/computeMetadata/v1/instance/attributes/kube-env"
	RolloutStrategyPath   = "/computeMetadata/v1/instance/attributes/endpoints-rollout-strategy"
	ServiceNamePath       = "/computeMetadata/v1/instance/attributes/endpoints-service-name"

	AccessTokenPath   = "/computeMetadata/v1/instance/service-accounts/default/token"
	IdentityTokenPath = "/computeMetadata/v1/instance/service-accounts/default/identity"
	ProjectIDPath     = "/computeMetadata/v1/project/project-id"

	// Cloud Run platform is regional, use the region path.
	RegionPath = "/computeMetadata/v1/instance/region"

	// GKE/GCE platforms are zonal. Regional path does not exist in IMDS.
	ZonePath = "/computeMetadata/v1/instance/zone"

	// The path of getting access token from token agent server
	TokenAgentAccessTokenPath = "/local/access_token"

	// b/147591854: This string must NOT have a trailing slash
	OpenIDDiscoveryCfgURLSuffix = "/.well-known/openid-configuration"

	// Platforms
	GAEFlex = "GAE_FLEX(ESPv2)"
	GKE     = "GKE(ESPv2)"
	GCE     = "GCE(ESPv2)"

	// System Parameter Name
	ApiKeyParameterName = "api_key"

	// retriable-status-codes retryOn policy
	RetryOnRetriableStatusCodes = "retriable-status-codes"
	// Default response deadline used if user does not specify one in the BackendRule.
	DefaultResponseDeadline = 15 * time.Second

	// Default idle timeout applied globally if not specified via flag.
	DefaultIdleTimeout = 5 * time.Minute

	// A limit configured to restrict resource usage in Envoy's SafeRegex GoogleRE2 matcher.
	// It will be validated on configmanager side though it may use different GoogleRE2 library.
	// b/148606900: It is safe to set this to a fairly high value.
	// This won't impact resource usage for customers who have short UriTemplates.
	GoogleRE2MaxProgramSize = 1000

	// Default jwt locations
	DefaultJwtHeaderNameAuthorization          = "Authorization"
	DefaultJwtHeaderValuePrefixBearer          = "Bearer "
	DefaultJwtHeaderNameXGoogleIapJwtAssertion = "X-Goog-Iap-Jwt-Assertion"
	DefaultJwtQueryParamAccessToken            = "access_token"

	// The suffix of jwtAuthn filter header to forward payload
	JwtAuthnForwardPayloadHeaderSuffix = "API-UserInfo"

	// Default api key locations
	DefaultApiKeyQueryParamKey    = "key"
	DefaultApiKeyQueryParamApiKey = "api_key"

	// Strict Transport Security header key and value
	HSTSHeaderKey   = "Strict-Transport-Security"
	HSTSHeaderValue = "max-age=31536000; includeSubdomains"

	// Standard type url prefix.
	TypeUrlPrefix = "type.googleapis.com/"

	// Loopback Address
	LoopbackIPv4Addr = "127.0.0.1"

	// All operations auto-generated by ESPv2 be in the format:
	// `{prefix}_{component}`, with an optional `_{formatted_path}` suffix.
	AutogeneratedOperationPrefix = "ESPv2_Autogenerated"

	// For operations not tied to a specific API.
	EspOperation = "espv2_deployment"

	// All traces created by ESPv2 should have this prefix.
	SpanNamePrefix = "ingress"

	// The maximum byte number of a span name. This restriction is from StackDriver.
	SpanNameMaxByteNum = 128

	// The stat prefix.
	StatPrefix = "ingress_http"

	// The suffix that forms the operation name header.
	OperationHeaderSuffix = "Api-Operation-Name"

	// The serverless platform for the flag --compute_platform_override
	// It is copied from SERVERLESS_PLATFORM at "docker/start_proxy.py"
	ServerlessPlatform = "Cloud Run(ESPv2)"

	// HTTPBackendProtocolKey is the HTTP backend rule key defined in backend rules.
	HTTPBackendProtocolKey = "http"
)

type BackendProtocol int32

type GetAccessTokenFunc func() (string, time.Duration, error)
type GetNewRolloutIdFunc func() (string, error)

// Backend protocol.
const (
	UNKNOWN BackendProtocol = iota
	HTTP1
	HTTP2
	GRPC
)

func MaybeTruncateSpanName(spanName string) string {
	if len(spanName) <= SpanNameMaxByteNum {
		return spanName
	}
	newSpanName := spanName[:SpanNameMaxByteNum-3] + "..."
	return newSpanName
}

// HardCodedSkipServiceControlMethods is a list of methods that should skip
// service control by default.
var HardCodedSkipServiceControlMethods = []string{
	"grpc.health.v1.Health.Check",
	"grpc.health.v1.Health.Watch",
}
