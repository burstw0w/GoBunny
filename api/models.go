package api

type PatternTrigger struct {
	Type                int      `json:"Type"`
	PatternMatches      []string `json:"PatternMatches"`
	PatternMatchingType int      `json:"PatternMatchingType"`
	Parameter1          string   `json:"Parameter1"`
}

type EdgeRule struct {
	Guid                string           `json:"Guid"`
	ActionType          int              `json:"ActionType"`
	ActionParameter1    string           `json:"ActionParameter1"`
	Description         string           `json:"Description"`
	Enabled             bool             `json:"Enabled"`
	OrderIndex          int              `json:"OrderIndex"`
	Triggers            []PatternTrigger `json:"Triggers"`
}

type HostnameBasic struct {
	Id    int    `json:"Id"`
	Value string `json:"Value"`
}

type PullZoneBasic struct {
	Id        int             `json:"Id"`
	Name      string          `json:"Name"`
	OriginUrl string          `json:"OriginUrl"`
	Enabled   bool            `json:"Enabled"`
	Suspended bool            `json:"Suspended"`
	Hostnames []HostnameBasic `json:"Hostnames"`
	EdgeRules []EdgeRule      `json:"EdgeRules"`
}

type PatternTriggerFull struct {
	Type                int      `json:"Type"`
	PatternMatches      []string `json:"PatternMatches"`
	PatternMatchingType int      `json:"PatternMatchingType"`
	Parameter1          *string  `json:"Parameter1"` // Pointer handles null values
}

type EdgeRuleFull struct {
	Guid                string               `json:"Guid"`
	ActionType          int                  `json:"ActionType"`
	ActionParameter1    string               `json:"ActionParameter1"`
	ActionParameter2    string               `json:"ActionParameter2"`
	ActionParameter3    string               `json:"ActionParameter3"`
	Triggers            []PatternTriggerFull `json:"Triggers"`
	ExtraActions        []interface{}        `json:"ExtraActions"` // Interfaces used as these arrays are empty in the JSON
	TriggerMatchingType int                  `json:"TriggerMatchingType"`
	Description         string               `json:"Description"`
	Enabled             bool                 `json:"Enabled"`
	OrderIndex          int                  `json:"OrderIndex"`
	ReadOnly            bool                 `json:"ReadOnly"`
}

type HostnameFull struct {
	Id                       int     `json:"Id"`
	Value                    string  `json:"Value"`
	ForceSSL                 bool    `json:"ForceSSL"`
	IsSystemHostname         bool    `json:"IsSystemHostname"`
	IsManagedHostname        bool    `json:"IsManagedHostname"`
	HasCertificate           bool    `json:"HasCertificate"`
	Certificate              *string `json:"Certificate"`
	CertificateKey           *string `json:"CertificateKey"`
	CertificateKeyType       int     `json:"CertificateKeyType"`
	CertificateProvisionType int     `json:"CertificateProvisionType"`
}

type BunnyAiImageBlueprintProperties struct {
	PrePrompt  string `json:"PrePrompt"`
	PostPrompt string `json:"PostPrompt"`
}

type BunnyAiImageBlueprint struct {
	Name       string                          `json:"Name"`
	Properties BunnyAiImageBlueprintProperties `json:"Properties"`
}

type PullZoneFull struct {
	Id                                  int                     `json:"Id"`
	Name                                string                  `json:"Name"`
	OriginUrl                           string                  `json:"OriginUrl"`
	Enabled                             bool                    `json:"Enabled"`
	Suspended                           bool                    `json:"Suspended"`
	Hostnames                           []HostnameFull          `json:"Hostnames"`
	StorageZoneId                       int                     `json:"StorageZoneId"`
	EdgeScriptId                        int                     `json:"EdgeScriptId"`
	EdgeScriptExecutionPhase            int                     `json:"EdgeScriptExecutionPhase"`
	MiddlewareScriptId                  *string                 `json:"MiddlewareScriptId"`
	MagicContainersAppId                *string                 `json:"MagicContainersAppId"`
	MagicContainersEndpointId           *string                 `json:"MagicContainersEndpointId"`
	AllowedReferrers                    []string                `json:"AllowedReferrers"`
	BlockedReferrers                    []string                `json:"BlockedReferrers"`
	BlockedIps                          []string                `json:"BlockedIps"`
	EnableGeoZoneUS                     bool                    `json:"EnableGeoZoneUS"`
	EnableGeoZoneEU                     bool                    `json:"EnableGeoZoneEU"`
	EnableGeoZoneASIA                   bool                    `json:"EnableGeoZoneASIA"`
	EnableGeoZoneSA                     bool                    `json:"EnableGeoZoneSA"`
	EnableGeoZoneAF                     bool                    `json:"EnableGeoZoneAF"`
	ZoneSecurityEnabled                 bool                    `json:"ZoneSecurityEnabled"`
	ZoneSecurityKey                     string                  `json:"ZoneSecurityKey"`
	ZoneSecurityIncludeHashRemoteIP     bool                    `json:"ZoneSecurityIncludeHashRemoteIP"`
	IgnoreQueryStrings                  bool                    `json:"IgnoreQueryStrings"`
	MonthlyBandwidthLimit               int64                   `json:"MonthlyBandwidthLimit"`
	MonthlyBandwidthUsed                int64                   `json:"MonthlyBandwidthUsed"` // Int64 to prevent overflow
	MonthlyCharges                      float64                 `json:"MonthlyCharges"`
	AddHostHeader                       bool                    `json:"AddHostHeader"`
	OriginHostHeader                    string                  `json:"OriginHostHeader"`
	Type                                int                     `json:"Type"`
	AccessControlOriginHeaderExtensions []string                `json:"AccessControlOriginHeaderExtensions"`
	EnableAccessControlOriginHeader     bool                    `json:"EnableAccessControlOriginHeader"`
	DisableCookies                      bool                    `json:"DisableCookies"`
	BudgetRedirectedCountries           []string                `json:"BudgetRedirectedCountries"`
	BlockedCountries                    []string                `json:"BlockedCountries"`
	EnableOriginShield                  bool                    `json:"EnableOriginShield"`
	CacheControlMaxAgeOverride          int                     `json:"CacheControlMaxAgeOverride"`
	CacheControlPublicMaxAgeOverride    int                     `json:"CacheControlPublicMaxAgeOverride"`
	BurstSize                           int                     `json:"BurstSize"`
	RequestLimit                        int                     `json:"RequestLimit"`
	BlockRootPathAccess                 bool                    `json:"BlockRootPathAccess"`
	BlockPostRequests                   bool                    `json:"BlockPostRequests"`
	LimitRateAfter      				float64 				`json:"LimitRateAfter"`
	LimitRatePerSecond  				float64 				`json:"LimitRatePerSecond"`
	ConnectionLimitPerIPCount           int                     `json:"ConnectionLimitPerIPCount"`
	PriceOverride                       float64                 `json:"PriceOverride"`
	OptimizerPricing                    float64                 `json:"OptimizerPricing"`
	AddCanonicalHeader                  bool                    `json:"AddCanonicalHeader"`
	EnableLogging                       bool                    `json:"EnableLogging"`
	EnableCacheSlice                    bool                    `json:"EnableCacheSlice"`
	EnableSmartCache                    bool                    `json:"EnableSmartCache"`
	EdgeRules                           []EdgeRuleFull          `json:"EdgeRules"`
	EnableWebPVary                      bool                    `json:"EnableWebPVary"`
	EnableAvifVary                      bool                    `json:"EnableAvifVary"`
	EnableCountryCodeVary               bool                    `json:"EnableCountryCodeVary"`
	EnableCountryStateCodeVary          bool                    `json:"EnableCountryStateCodeVary"`
	EnableMobileVary                    bool                    `json:"EnableMobileVary"`
	EnableCookieVary                    bool                    `json:"EnableCookieVary"`
	CookieVaryParameters                []string                `json:"CookieVaryParameters"`
	EnableHostnameVary                  bool                    `json:"EnableHostnameVary"`
	CnameDomain                         string                  `json:"CnameDomain"`
	AWSSigningEnabled                   bool                    `json:"AWSSigningEnabled"`
	AWSSigningKey                       *string                 `json:"AWSSigningKey"`
	AWSSigningSecret                    *string                 `json:"AWSSigningSecret"`
	AWSSigningRegionName                *string                 `json:"AWSSigningRegionName"`
	LoggingIPAnonymizationEnabled       bool                    `json:"LoggingIPAnonymizationEnabled"`
	EnableTLS1                          bool                    `json:"EnableTLS1"`
	EnableTLS1_1                        bool                    `json:"EnableTLS1_1"`
	VerifyOriginSSL                     bool                    `json:"VerifyOriginSSL"`
	ErrorPageEnableCustomCode           bool                    `json:"ErrorPageEnableCustomCode"`
	ErrorPageCustomCode                 *string                 `json:"ErrorPageCustomCode"`
	ErrorPageEnableStatuspageWidget     bool                    `json:"ErrorPageEnableStatuspageWidget"`
	ErrorPageStatuspageCode             *string                 `json:"ErrorPageStatuspageCode"`
	ErrorPageWhitelabel                 bool                    `json:"ErrorPageWhitelabel"`
	OriginShieldZoneCode                string                  `json:"OriginShieldZoneCode"`
	LogForwardingEnabled                bool                    `json:"LogForwardingEnabled"`
	LogForwardingHostname               *string                 `json:"LogForwardingHostname"`
	LogForwardingPort                   int                     `json:"LogForwardingPort"`
	LogForwardingToken                  *string                 `json:"LogForwardingToken"`
	LogForwardingProtocol               int                     `json:"LogForwardingProtocol"`
	LoggingSaveToStorage                bool                    `json:"LoggingSaveToStorage"`
	LoggingStorageZoneId                int                     `json:"LoggingStorageZoneId"`
	FollowRedirects                     bool                    `json:"FollowRedirects"`
	VideoLibraryId                      int                     `json:"VideoLibraryId"`
	DnsRecordId                         int                     `json:"DnsRecordId"`
	DnsZoneId                           int                     `json:"DnsZoneId"`
	DnsRecordValue                      *string                 `json:"DnsRecordValue"`
	OptimizerEnabled                    bool                    `json:"OptimizerEnabled"`
	OptimizerTunnelEnabled              bool                    `json:"OptimizerTunnelEnabled"`
	OptimizerDesktopMaxWidth            int                     `json:"OptimizerDesktopMaxWidth"`
	OptimizerMobileMaxWidth             int                     `json:"OptimizerMobileMaxWidth"`
	OptimizerImageQuality               int                     `json:"OptimizerImageQuality"`
	OptimizerMobileImageQuality         int                     `json:"OptimizerMobileImageQuality"`
	OptimizerEnableWebP                 bool                    `json:"OptimizerEnableWebP"`
	OptimizerPrerenderHtml              bool                    `json:"OptimizerPrerenderHtml"`
	OptimizerEnableManipulationEngine   bool                    `json:"OptimizerEnableManipulationEngine"`
	OptimizerMinifyCSS                  bool                    `json:"OptimizerMinifyCSS"`
	OptimizerMinifyJavaScript           bool                    `json:"OptimizerMinifyJavaScript"`
	OptimizerWatermarkEnabled           bool                    `json:"OptimizerWatermarkEnabled"`
	OptimizerWatermarkUrl               string                  `json:"OptimizerWatermarkUrl"`
	OptimizerWatermarkPosition          int                     `json:"OptimizerWatermarkPosition"`
	OptimizerWatermarkOffset            float64                 `json:"OptimizerWatermarkOffset"`
	OptimizerWatermarkMinImageSize      int                     `json:"OptimizerWatermarkMinImageSize"`
	OptimizerAutomaticOptimizationEnabled bool                  `json:"OptimizerAutomaticOptimizationEnabled"`
	PermaCacheStorageZoneId             int                     `json:"PermaCacheStorageZoneId"`
	PermaCacheType                      int                     `json:"PermaCacheType"`
	OriginRetries                       int                     `json:"OriginRetries"`
	OriginConnectTimeout                int                     `json:"OriginConnectTimeout"`
	OriginResponseTimeout               int                     `json:"OriginResponseTimeout"`
	UseStaleWhileUpdating               bool                    `json:"UseStaleWhileUpdating"`
	UseStaleWhileOffline                bool                    `json:"UseStaleWhileOffline"`
	OriginRetry5XXResponses             bool                    `json:"OriginRetry5XXResponses"`
	OriginRetryConnectionTimeout        bool                    `json:"OriginRetryConnectionTimeout"`
	OriginRetryResponseTimeout          bool                    `json:"OriginRetryResponseTimeout"`
	OriginRetryDelay                    int                     `json:"OriginRetryDelay"`
	QueryStringVaryParameters           []string                `json:"QueryStringVaryParameters"`
	OriginShieldEnableConcurrencyLimit  bool                    `json:"OriginShieldEnableConcurrencyLimit"`
	OriginShieldMaxConcurrentRequests   int                     `json:"OriginShieldMaxConcurrentRequests"`
	EnableSafeHop                       bool                    `json:"EnableSafeHop"`
	CacheErrorResponses                 bool                    `json:"CacheErrorResponses"`
	OriginShieldQueueMaxWaitTime        int                     `json:"OriginShieldQueueMaxWaitTime"`
	OriginShieldMaxQueuedRequests       int                     `json:"OriginShieldMaxQueuedRequests"`
	OptimizerClasses                    []string                `json:"OptimizerClasses"`
	OptimizerForceClasses               bool                    `json:"OptimizerForceClasses"`
	OptimizerStaticHtmlEnabled          bool                    `json:"OptimizerStaticHtmlEnabled"`
	OptimizerStaticHtmlWordPressPath    *string                 `json:"OptimizerStaticHtmlWordPressPath"`
	OptimizerStaticHtmlWordPressBypassCookie *string            `json:"OptimizerStaticHtmlWordPressBypassCookie"`
	UseBackgroundUpdate                 bool                    `json:"UseBackgroundUpdate"`
	EnableAutoSSL                       bool                    `json:"EnableAutoSSL"`
	EnableQueryStringOrdering           bool                    `json:"EnableQueryStringOrdering"`
	LogAnonymizationType                int                     `json:"LogAnonymizationType"`
	LogFormat                           int                     `json:"LogFormat"`
	LogForwardingFormat                 int                     `json:"LogForwardingFormat"`
	ShieldDDosProtectionType            int                     `json:"ShieldDDosProtectionType"`
	ShieldDDosProtectionEnabled         bool                    `json:"ShieldDDosProtectionEnabled"`
	OriginType                          int                     `json:"OriginType"`
	EnableRequestCoalescing             bool                    `json:"EnableRequestCoalescing"`
	RequestCoalescingTimeout            int                     `json:"RequestCoalescingTimeout"`
	OriginLinkValue                     string                  `json:"OriginLinkValue"`
	DisableLetsEncrypt                  bool                    `json:"DisableLetsEncrypt"`
	EnableBunnyImageAi                  bool                    `json:"EnableBunnyImageAi"`
	BunnyAiImageBlueprints              []BunnyAiImageBlueprint `json:"BunnyAiImageBlueprints"`
	PreloadingScreenEnabled             bool                    `json:"PreloadingScreenEnabled"`
	PreloadingScreenShowOnFirstVisit    bool                    `json:"PreloadingScreenShowOnFirstVisit"`
	PreloadingScreenCode                string                  `json:"PreloadingScreenCode"`
	PreloadingScreenLogoUrl             *string                 `json:"PreloadingScreenLogoUrl"`
	PreloadingScreenCodeEnabled         bool                    `json:"PreloadingScreenCodeEnabled"`
	PreloadingScreenTheme               int                     `json:"PreloadingScreenTheme"`
	PreloadingScreenDelay               int                     `json:"PreloadingScreenDelay"`
	EUUSDiscount                        float64                 `json:"EUUSDiscount"`
	SouthAmericaDiscount                float64                 `json:"SouthAmericaDiscount"`
	AfricaDiscount                      float64                 `json:"AfricaDiscount"`
	AsiaOceaniaDiscount                 float64                 `json:"AsiaOceaniaDiscount"`
	RoutingFilters                      []string                `json:"RoutingFilters"`
	BlockNoneReferrer                   bool                    `json:"BlockNoneReferrer"`
	StickySessionType                   int                     `json:"StickySessionType"`
	StickySessionCookieName             *string                 `json:"StickySessionCookieName"`
	StickySessionClientHeaders          *string                 `json:"StickySessionClientHeaders"`
	UserId                              string                  `json:"UserId"`
	CacheVersion                        int                     `json:"CacheVersion"`
	OptimizerEnableUpscaling            bool                    `json:"OptimizerEnableUpscaling"`
	EnableWebSockets                    bool                    `json:"EnableWebSockets"`
	MaxWebSocketConnections             int                     `json:"MaxWebSocketConnections"`
	EnableExtendedLogging               bool                    `json:"EnableExtendedLogging"`
}
