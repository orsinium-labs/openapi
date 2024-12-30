package openapi

type OpenAPI struct {
	Version           string              `json:"openapi,omitempty"`
	Info              Info                `json:"info,omitempty"`
	JSONSchemaDialect string              `json:"jsonSchemaDialect,omitempty"`
	Servers           []Server            `json:"servers,omitempty"`
	Paths             Paths               `json:"paths,omitempty"`
	Webhooks          map[string]PathItem `json:"webhooks,omitempty"`
	Components        Components          `json:"components,omitempty"`
	Security          SecurityRequirement `json:"security,omitempty"`
	Tags              []Tag               `json:"tags,omitempty"`
	ExternalDocs      []ExternalDoc       `json:"externalDocs,omitempty"`
}

type Info struct {
	Title          string  `json:"title"`
	Summary        string  `json:"summary,omitempty"`
	Description    string  `json:"description,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
	Version        string  `json:"version,omitempty"`
}

type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type License struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier,omitempty"`
	URL        string `json:"url,omitempty"`
}

type Server struct {
	URL         string                    `json:"url,omitempty"`
	Description string                    `json:"description,omitempty"`
	Variables   map[string]ServerVariable `json:"variables,omitempty"`
}

type ServerVariable struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
}

type Components struct {
	Schemas         map[string]Schema         `json:"schemas,omitempty"`
	Responses       map[string]Response       `json:"responses,omitempty"`
	Parameters      map[string]Parameter      `json:"parameters,omitempty"`
	Examples        map[string]Example        `json:"examples,omitempty"`
	RequestBodies   map[string]RequestBody    `json:"requestBodies,omitempty"`
	Headers         map[string]Header         `json:"headers,omitempty"`
	SecuritySchemes map[string]SecurityScheme `json:"security_schemes,omitempty"`
	Links           map[string]Link           `json:"links,omitempty"`
	Callbacks       map[string]Callback       `json:"callbacks,omitempty"`
	PathItems       map[string]PathItem       `json:"pathItems,omitempty"`
}

type Paths map[string]PathItem

type PathItem struct {
	Ref         string      `json:"$ref,omitempty"`
	Summary     string      `json:"summary,omitempty"`
	Description string      `json:"description,omitempty"`
	Get         Operation   `json:"get,omitempty"`
	Put         Operation   `json:"put,omitempty"`
	Post        Operation   `json:"post,omitempty"`
	Delete      Operation   `json:"delete,omitempty"`
	Options     Operation   `json:"options,omitempty"`
	Head        Operation   `json:"head,omitempty"`
	Patch       Operation   `json:"patch,omitempty"`
	Trace       Operation   `json:"trace,omitempty"`
	Servers     []Server    `json:"servers,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty"`
}

type Operation struct {
	Tags         []string              `json:"tags,omitempty"`
	Summary      string                `json:"summary,omitempty"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs ExternalDoc           `json:"externalDocs,omitempty"`
	OperationID  string                `json:"operationId,omitempty"`
	Parameters   Parameter             `json:"parameters,omitempty"`
	RequestBody  RequestBody           `json:"requestBody,omitempty"`
	Responses    Responses             `json:"responses,omitempty"`
	Callbacks    map[string]Callback   `json:"callbacks,omitempty"`
	Deprecated   bool                  `json:"deprecated,omitempty"`
	Security     []SecurityRequirement `json:"security,omitempty"`
	Servers      []Server              `json:"servers,omitempty"`
}

type ExternalDoc struct {
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

type Parameter struct {
	Name            string `json:"name,omitempty"`
	In              string `json:"in,omitempty"`
	Description     string `json:"description,omitempty"`
	Required        bool   `json:"required,omitempty"`
	Deprecated      bool   `json:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty"`

	Style         string             `json:"style,omitempty"`
	Explode       bool               `json:"explode,omitempty"`
	AllowReserved bool               `json:"allow_reserved,omitempty"`
	Schema        Schema             `json:"schema,omitempty"`
	Example       any                `json:"example,omitempty"`
	Examples      map[string]Example `json:"examples,omitempty"`

	Content map[string]MediaType `json:"content,omitempty"`
}

type RequestBody struct {
	Description string               `json:"description,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
	Required    bool                 `json:"required,omitempty"`
}

type MediaType struct {
	Schema   Schema              `json:"schema,omitempty"`
	Example  any                 `json:"example,omitempty"`
	Examples map[string]Example  `json:"examples,omitempty"`
	Encoding map[string]Encoding `json:"encoding,omitempty"`
}

type Encoding struct {
	ContentType   string            `json:"contentType,omitempty"`
	Headers       map[string]Header `json:"headers,omitempty"`
	Style         string            `json:"style,omitempty"`
	Explode       bool              `json:"explode,omitempty"`
	AllowReserved bool              `json:"allow_reserved,omitempty"`
}

type Responses struct {
	Default                       Response
	Continue                      Response `json:"100,omitempty"`
	SwitchingProtocols            Response `json:"101,omitempty"`
	Processing                    Response `json:"102,omitempty"`
	EarlyHints                    Response `json:"103,omitempty"`
	OK                            Response `json:"200,omitempty"`
	Created                       Response `json:"201,omitempty"`
	Accepted                      Response `json:"202,omitempty"`
	NonAuthoritativeInfo          Response `json:"203,omitempty"`
	NoContent                     Response `json:"204,omitempty"`
	ResetContent                  Response `json:"205,omitempty"`
	PartialContent                Response `json:"206,omitempty"`
	MultiStatus                   Response `json:"207,omitempty"`
	AlreadyReported               Response `json:"208,omitempty"`
	IMUsed                        Response `json:"226,omitempty"`
	MultipleChoices               Response `json:"300,omitempty"`
	MovedPermanently              Response `json:"301,omitempty"`
	Found                         Response `json:"302,omitempty"`
	SeeOther                      Response `json:"303,omitempty"`
	NotModified                   Response `json:"304,omitempty"`
	UseProxy                      Response `json:"305,omitempty"`
	TemporaryRedirect             Response `json:"307,omitempty"`
	PermanentRedirect             Response `json:"308,omitempty"`
	BadRequest                    Response `json:"400,omitempty"`
	Unauthorized                  Response `json:"401,omitempty"`
	PaymentRequired               Response `json:"402,omitempty"`
	Forbidden                     Response `json:"403,omitempty"`
	NotFound                      Response `json:"404,omitempty"`
	MethodNotAllowed              Response `json:"405,omitempty"`
	NotAcceptable                 Response `json:"406,omitempty"`
	ProxyAuthRequired             Response `json:"407,omitempty"`
	RequestTimeout                Response `json:"408,omitempty"`
	Conflict                      Response `json:"409,omitempty"`
	Gone                          Response `json:"410,omitempty"`
	LengthRequired                Response `json:"411,omitempty"`
	PreconditionFailed            Response `json:"412,omitempty"`
	RequestEntityTooLarge         Response `json:"413,omitempty"`
	RequestURITooLong             Response `json:"414,omitempty"`
	UnsupportedMediaType          Response `json:"415,omitempty"`
	RequestedRangeNotSatisfiable  Response `json:"416,omitempty"`
	ExpectationFailed             Response `json:"417,omitempty"`
	Teapot                        Response `json:"418,omitempty"`
	MisdirectedRequest            Response `json:"421,omitempty"`
	UnprocessableEntity           Response `json:"422,omitempty"`
	Locked                        Response `json:"423,omitempty"`
	FailedDependency              Response `json:"424,omitempty"`
	TooEarly                      Response `json:"425,omitempty"`
	UpgradeRequired               Response `json:"426,omitempty"`
	PreconditionRequired          Response `json:"428,omitempty"`
	TooManyRequests               Response `json:"429,omitempty"`
	RequestHeaderFieldsTooLarge   Response `json:"431,omitempty"`
	UnavailableForLegalReasons    Response `json:"451,omitempty"`
	InternalServerError           Response `json:"500,omitempty"`
	NotImplemented                Response `json:"501,omitempty"`
	BadGateway                    Response `json:"502,omitempty"`
	ServiceUnavailable            Response `json:"503,omitempty"`
	GatewayTimeout                Response `json:"504,omitempty"`
	HTTPVersionNotSupported       Response `json:"505,omitempty"`
	VariantAlsoNegotiates         Response `json:"506,omitempty"`
	InsufficientStorage           Response `json:"507,omitempty"`
	LoopDetected                  Response `json:"508,omitempty"`
	NotExtended                   Response `json:"510,omitempty"`
	NetworkAuthenticationRequired Response `json:"511,omitempty"`
}

type Response struct {
	Description string               `json:"description,omitempty"`
	Headers     map[string]Header    `json:"headers,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
	Links       map[string]Link      `json:"links,omitempty"`
}

type Callback map[string]PathItem

type Example struct {
	Summary       string `json:"summary,omitempty"`
	Description   string `json:"description,omitempty"`
	Value         any    `json:"value,omitempty"`
	ExternalValue string `json:"externalValue,omitempty"`
}

type Link struct {
	OperationRef string         `json:"operationRef,omitempty"`
	OperationID  string         `json:"operationId,omitempty"`
	Parameters   map[string]any `json:"parameters,omitempty"`
	RequestBody  any            `json:"requestBody,omitempty"`
	Description  string         `json:"description,omitempty"`
	Server       Server         `json:"server,omitempty"`
}

type Header struct {
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Deprecated  bool   `json:"deprecated,omitempty"`

	Style         string             `json:"style,omitempty"`
	Explode       bool               `json:"explode,omitempty"`
	AllowReserved bool               `json:"allow_reserved,omitempty"`
	Schema        Schema             `json:"schema,omitempty"`
	Example       any                `json:"example,omitempty"`
	Examples      map[string]Example `json:"examples,omitempty"`

	Content map[string]MediaType `json:"content,omitempty"`
}

type Tag struct {
	Name         string      `json:"name,omitempty"`
	Description  string      `json:"description,omitempty"`
	ExternalDocs ExternalDoc `json:"externalDocs,omitempty"`
}

type Reference struct {
	Ref         string `json:"$ref,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}

type Schema = any

type SecurityScheme struct {
	Type             string     `json:"type,omitempty"`
	Description      string     `json:"description,omitempty"`
	Name             string     `json:"name,omitempty"`
	In               string     `json:"in,omitempty"`
	Scheme           string     `json:"scheme,omitempty"`
	BearerFormat     string     `json:"bearerFormat,omitempty"`
	Flows            OAuthFlows `json:"flows,omitempty"`
	OpenIDConnectURL string     `json:"openIdConnectUrl,omitempty"`
}

type OAuthFlows struct {
	Implicit          OAuthFlow `json:"implicit,omitempty"`
	Password          OAuthFlow `json:"password,omitempty"`
	ClientCredentials OAuthFlow `json:"clientCredentials,omitempty"`
	AuthorizationCode OAuthFlow `json:"authorizationCode,omitempty"`
}

type OAuthFlow struct {
	AuthorizationUrl string            `json:"authorizationUrl,omitempty"`
	TokenUrl         string            `json:"tokenUrl,omitempty"`
	RefreshUrl       string            `json:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty"`
}

type SecurityRequirement map[string][]string
