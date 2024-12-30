package openapi

// This is the root object of the OpenAPI Description.
type OpenAPI struct {
	// REQUIRED. This string MUST be the version number of the OpenAPI Specification that the OpenAPI Document uses. The openapi field SHOULD be used by tooling to interpret the OpenAPI Document. This is not related to the API info.version string.
	Version string `json:"openapi,omitempty"`
	//  	REQUIRED. Provides metadata about the API. The metadata MAY be used by tooling as required.
	Info Info `json:"info,omitempty"`
	// The default value for the $schema keyword within Schema Objects contained within this OAS document. This MUST be in the form of a URI.
	JSONSchemaDialect string `json:"jsonSchemaDialect,omitempty"`
	// An array of Server Objects, which provide connectivity information to a target server. If the servers field is not provided, or is an empty array, the default value would be a Server Object with a url value of /.
	Servers []Server `json:"servers,omitempty"`
	// The available paths and operations for the API.
	Paths Paths `json:"paths,omitempty"`
	// The incoming webhooks that MAY be received as part of this API and that the API consumer MAY choose to implement. Closely related to the callbacks feature, this section describes requests initiated other than by an API call, for example by an out of band registration. The key name is a unique string to refer to each webhook, while the (optionally referenced) Path Item Object describes a request that may be initiated by the API provider and the expected responses. An example is available.
	Webhooks map[string]PathItem `json:"webhooks,omitempty"`
	// An element to hold various Objects for the OpenAPI Description.
	Components Components `json:"components,omitempty"`
	// A declaration of which security mechanisms can be used across the API. The list of values includes alternative Security Requirement Objects that can be used. Only one of the Security Requirement Objects need to be satisfied to authorize a request. Individual operations can override this definition. The list can be incomplete, up to being empty or absent. To make security explicitly optional, an empty security requirement ({}) can be included in the array.
	Security SecurityRequirement `json:"security,omitempty"`
	// A list of tags used by the OpenAPI Description with additional metadata. The order of the tags can be used to reflect on their order by the parsing tools. Not all tags that are used by the Operation Object must be declared. The tags that are not declared MAY be organized randomly or based on the tools' logic. Each tag name in the list MUST be unique.
	Tags []Tag `json:"tags,omitempty"`
	// Additional external documentation.
	ExternalDocs []ExternalDoc `json:"externalDocs,omitempty"`
}

type Info struct {
	// REQUIRED. The title of the API.
	Title string `json:"title"`
	// A short summary of the API.
	Summary string `json:"summary,omitempty"`
	// A description of the API. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// A URI for the Terms of Service for the API. This MUST be in the form of a URI.
	TermsOfService string `json:"termsOfService,omitempty"`
	// The contact information for the exposed API.
	Contact Contact `json:"contact,omitempty"`
	// The license information for the exposed API.
	License License `json:"license,omitempty"`
	// REQUIRED. The version of the OpenAPI Document (which is distinct from the OpenAPI Specification version or the version of the API being described or the version of the OpenAPI Description).
	Version string `json:"version,omitempty"`
}

type Contact struct {
	// The identifying name of the contact person/organization.
	Name string `json:"name,omitempty"`
	// The URI for the contact information. This MUST be in the form of a URI.
	URL string `json:"url,omitempty"`
	// The email address of the contact person/organization. This MUST be in the form of an email address.
	Email string `json:"email,omitempty"`
}

type License struct {
	// REQUIRED. The license name used for the API.
	Name string `json:"name"`
	// An SPDX license expression for the API. The identifier field is mutually exclusive of the url field.
	Identifier string `json:"identifier,omitempty"`
	// A URI for the license used for the API. This MUST be in the form of a URI. The url field is mutually exclusive of the identifier field.
	URL string `json:"url,omitempty"`
}

// An object representing a Server.
type Server struct {
	// REQUIRED. A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the document containing the Server Object is being served. Variable substitutions will be made when a variable is named in {braces}.
	URL string `json:"url,omitempty"`
	// An optional string describing the host designated by the URL. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// A map between a variable name and its value. The value is used for substitution in the server's URL template.
	Variables map[string]ServerVariable `json:"variables,omitempty"`
}

// An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	// An enumeration of string values to be used if the substitution options are from a limited set. The array MUST NOT be empty.
	Enum []string `json:"enum,omitempty"`
	// REQUIRED. The default value to use for substitution, which SHALL be sent if an alternate value is not supplied. If the enum is defined, the value MUST exist in the enum's values. Note that this behavior is different from the Schema Object's default keyword, which documents the receiver's behavior rather than inserting the value into the data.
	Default string `json:"default,omitempty"`
	// An optional description for the server variable. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
}

// Holds a set of reusable objects for different aspects of the OAS. All objects defined within the Components Object will have no effect on the API unless they are explicitly referenced from outside the Components Object.
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

// Holds the relative paths to the individual endpoints and their operations. The path is appended to the URL from the Server Object in order to construct the full URL. The Paths Object MAY be empty, due to Access Control List (ACL) constraints.
type Paths map[string]PathItem

// Describes the operations available on a single path. A Path Item MAY be empty, due to ACL constraints. The path itself is still exposed to the documentation viewer but they will not know which operations and parameters are available.
type PathItem struct {
	// Allows for a referenced definition of this path item. The value MUST be in the form of a URI, and the referenced structure MUST be in the form of a Path Item Object. In case a Path Item Object field appears both in the defined object and the referenced object, the behavior is undefined. See the rules for resolving Relative References.
	Ref string `json:"$ref,omitempty"`
	// An optional string summary, intended to apply to all operations in this path.
	Summary string `json:"summary,omitempty"`
	// An optional string description, intended to apply to all operations in this path. CommonMark syntax MAY be used for rich text representation.
	Description string    `json:"description,omitempty"`
	Get         Operation `json:"get,omitempty"`
	Put         Operation `json:"put,omitempty"`
	Post        Operation `json:"post,omitempty"`
	Delete      Operation `json:"delete,omitempty"`
	Options     Operation `json:"options,omitempty"`
	Head        Operation `json:"head,omitempty"`
	Patch       Operation `json:"patch,omitempty"`
	Trace       Operation `json:"trace,omitempty"`
	// An alternative servers array to service all operations in this path. If a servers array is specified at the OpenAPI Object level, it will be overridden by this value.
	Servers []Server `json:"servers,omitempty"`
	// A list of parameters that are applicable for all the operations described under this path. These parameters can be overridden at the operation level, but cannot be removed there. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined in the OpenAPI Object's components.parameters.
	Parameters []Parameter `json:"parameters,omitempty"`
}

// Describes a single API operation on a path.
type Operation struct {
	// A list of tags for API documentation control. Tags can be used for logical grouping of operations by resources or any other qualifier.
	Tags []string `json:"tags,omitempty"`
	// A short summary of what the operation does.
	Summary string `json:"summary,omitempty"`
	// A verbose explanation of the operation behavior. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// Additional external documentation for this operation.
	ExternalDocs ExternalDoc `json:"externalDocs,omitempty"`
	// Unique string used to identify the operation. The id MUST be unique among all operations described in the API. The operationId value is case-sensitive. Tools and libraries MAY use the operationId to uniquely identify an operation, therefore, it is RECOMMENDED to follow common programming naming conventions.
	OperationID string `json:"operationId,omitempty"`
	// A list of parameters that are applicable for this operation. If a parameter is already defined at the Path Item, the new definition will override it but can never remove it. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined in the OpenAPI Object's components.parameters.
	Parameters Parameter `json:"parameters,omitempty"`
	// The request body applicable for this operation. The requestBody is fully supported in HTTP methods where the HTTP 1.1 specification RFC7231 has explicitly defined semantics for request bodies. In other cases where the HTTP spec is vague (such as GET, HEAD and DELETE), requestBody is permitted but does not have well-defined semantics and SHOULD be avoided if possible.
	RequestBody RequestBody `json:"requestBody,omitempty"`
	// The list of possible responses as they are returned from executing this operation.
	Responses Responses `json:"responses,omitempty"`
	// A map of possible out-of band callbacks related to the parent operation. The key is a unique identifier for the Callback Object. Each value in the map is a Callback Object that describes a request that may be initiated by the API provider and the expected responses.
	Callbacks map[string]Callback `json:"callbacks,omitempty"`
	// Declares this operation to be deprecated. Consumers SHOULD refrain from usage of the declared operation. Default value is false.
	Deprecated bool `json:"deprecated,omitempty"`
	// A declaration of which security mechanisms can be used for this operation. The list of values includes alternative Security Requirement Objects that can be used. Only one of the Security Requirement Objects need to be satisfied to authorize a request. To make security optional, an empty security requirement ({}) can be included in the array. This definition overrides any declared top-level security. To remove a top-level security declaration, an empty array can be used.
	Security []SecurityRequirement `json:"security,omitempty"`
	// An alternative servers array to service this operation. If a servers array is specified at the Path Item Object or OpenAPI Object level, it will be overridden by this value.
	Servers []Server `json:"servers,omitempty"`
}

// Allows referencing an external resource for extended documentation.
type ExternalDoc struct {
	// REQUIRED. The URI for the target documentation. This MUST be in the form of a URI.
	URL string `json:"url,omitempty"`
	// A description of the target documentation. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
}

// Describes a single operation parameter.
type Parameter struct {
	// REQUIRED. The name of the parameter. Parameter names are case sensitive.
	Name string `json:"name,omitempty"`
	// REQUIRED. The location of the parameter. Possible values are "query", "header", "path" or "cookie".
	In string `json:"in,omitempty"`
	// A brief description of the parameter. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// Determines whether this parameter is mandatory. If the parameter location is "path", this field is REQUIRED and its value MUST be true. Otherwise, the field MAY be included and its default value is false.
	Required bool `json:"required,omitempty"`
	// Specifies that a parameter is deprecated and SHOULD be transitioned out of usage. Default value is false.
	Deprecated bool `json:"deprecated,omitempty"`
	// If true, clients MAY pass a zero-length string value in place of parameters that would otherwise be omitted entirely, which the server SHOULD interpret as the parameter being unused. Default value is false. If style is used, and if behavior is n/a (cannot be serialized), the value of allowEmptyValue SHALL be ignored. Interactions between this field and the parameter's Schema Object are implementation-defined. This field is valid only for query parameters. Use of this field is NOT RECOMMENDED, and it is likely to be removed in a later revision.
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty"`

	// Describes how the parameter value will be serialized depending on the type of the parameter value. Default values (based on value of in): for "query" - "form"; for "path" - "simple"; for "header" - "simple"; for "cookie" - "form".
	Style string `json:"style,omitempty"`
	// When this is true, parameter values of type array or object generate separate parameters for each value of the array or key-value pair of the map. For other types of parameters this field has no effect. When style is "form", the default value is true. For all other styles, the default value is false. Note that despite false being the default for deepObject, the combination of false with deepObject is undefined.
	Explode bool `json:"explode,omitempty"`
	// When this is true, parameter values are serialized using reserved expansion, as defined by RFC6570, which allows RFC3986's reserved character set, as well as percent-encoded triples, to pass through unchanged, while still percent-encoding all other disallowed characters (including % outside of percent-encoded triples). Applications are still responsible for percent-encoding reserved characters that are not allowed in the query string ([, ], #), or have a special meaning in application/x-www-form-urlencoded (-, &, +); see Appendices C and E for details. This field only applies to parameters with an in value of query. The default value is false.
	AllowReserved bool `json:"allow_reserved,omitempty"`
	// The schema defining the type used for the parameter.
	Schema Schema `json:"schema,omitempty"`
	// Example of the parameter's potential value; see Working With Examples.
	Example any `json:"example,omitempty"`
	// Examples of the parameter's potential value; see Working With Examples.
	Examples map[string]Example `json:"examples,omitempty"`

	// A map containing the representations for the parameter. The key is the media type and the value describes it. The map MUST only contain one entry.
	Content map[string]MediaType `json:"content,omitempty"`
}

// Describes a single request body.
type RequestBody struct {
	// A brief description of the request body. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// REQUIRED. The content of the request body. The key is a media type or media type range and the value describes it. For requests that match multiple keys, only the most specific key is applicable. e.g. "text/plain" overrides "text/*"
	Content map[string]MediaType `json:"content"`
	// Determines if the request body is required in the request. Defaults to false.
	Required bool `json:"required,omitempty"`
}

// Each Media Type Object provides schema and examples for the media type identified by its key.
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
