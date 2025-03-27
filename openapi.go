package openapi

// This is the root object of the OpenAPI Description.
type OpenAPI struct {
	// REQUIRED. This string MUST be the version number of the OpenAPI Specification that the OpenAPI Document uses. The openapi field SHOULD be used by tooling to interpret the OpenAPI Document. This is not related to the API info.version string.
	Version string `json:"openapi"`
	// REQUIRED. Provides metadata about the API. The metadata MAY be used by tooling as required.
	Info Info `json:"info"`
	// The default value for the $schema keyword within Schema Objects contained within this OAS document. This MUST be in the form of a URI.
	JSONSchemaDialect string `json:"jsonSchemaDialect,omitzero"`
	// An array of Server Objects, which provide connectivity information to a target server. If the servers field is not provided, or is an empty array, the default value would be a Server Object with a url value of /.
	Servers []Server `json:"servers,omitzero"`
	// The available paths and operations for the API.
	Paths Paths `json:"paths,omitzero"`
	// The incoming webhooks that MAY be received as part of this API and that the API consumer MAY choose to implement. Closely related to the callbacks feature, this section describes requests initiated other than by an API call, for example by an out of band registration. The key name is a unique string to refer to each webhook, while the (optionally referenced) Path Item Object describes a request that may be initiated by the API provider and the expected responses. An example is available.
	Webhooks map[string]PathItem `json:"webhooks,omitzero"`
	// An element to hold various Objects for the OpenAPI Description.
	Components Components `json:"components,omitzero"`
	// A declaration of which security mechanisms can be used across the API. The list of values includes alternative Security Requirement Objects that can be used. Only one of the Security Requirement Objects need to be satisfied to authorize a request. Individual operations can override this definition. The list can be incomplete, up to being empty or absent. To make security explicitly optional, an empty security requirement ({}) can be included in the array.
	Security SecurityRequirement `json:"security,omitzero"`
	// A list of tags used by the OpenAPI Description with additional metadata. The order of the tags can be used to reflect on their order by the parsing tools. Not all tags that are used by the Operation Object must be declared. The tags that are not declared MAY be organized randomly or based on the tools' logic. Each tag name in the list MUST be unique.
	Tags []Tag `json:"tags,omitzero"`
	// Additional external documentation.
	ExternalDocs []ExternalDoc `json:"externalDocs,omitzero"`
}

type Info struct {
	// REQUIRED. The title of the API.
	Title string `json:"title"`
	// A short summary of the API.
	Summary string `json:"summary,omitzero"`
	// A description of the API. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
	// A URI for the Terms of Service for the API. This MUST be in the form of a URI.
	TermsOfService string `json:"termsOfService,omitzero"`
	// The contact information for the exposed API.
	Contact Contact `json:"contact,omitzero"`
	// The license information for the exposed API.
	License License `json:"license,omitzero"`
	// REQUIRED. The version of the OpenAPI Document (which is distinct from the OpenAPI Specification version or the version of the API being described or the version of the OpenAPI Description).
	Version string `json:"version,omitzero"`
}

type Contact struct {
	// The identifying name of the contact person/organization.
	Name string `json:"name,omitzero"`
	// The URI for the contact information. This MUST be in the form of a URI.
	URL string `json:"url,omitzero"`
	// The email address of the contact person/organization. This MUST be in the form of an email address.
	Email string `json:"email,omitzero"`
}

type License struct {
	// REQUIRED. The license name used for the API.
	Name string `json:"name"`
	// An SPDX license expression for the API. The identifier field is mutually exclusive of the url field.
	Identifier string `json:"identifier,omitzero"`
	// A URI for the license used for the API. This MUST be in the form of a URI. The url field is mutually exclusive of the identifier field.
	URL string `json:"url,omitzero"`
}

// An object representing a Server.
type Server struct {
	// REQUIRED. A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the document containing the Server Object is being served. Variable substitutions will be made when a variable is named in {braces}.
	URL string `json:"url"`
	// An optional string describing the host designated by the URL. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
	// A map between a variable name and its value. The value is used for substitution in the server's URL template.
	Variables map[string]ServerVariable `json:"variables,omitzero"`
}

// An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	// An enumeration of string values to be used if the substitution options are from a limited set. The array MUST NOT be empty.
	Enum []string `json:"enum,omitzero"`
	// REQUIRED. The default value to use for substitution, which SHALL be sent if an alternate value is not supplied. If the enum is defined, the value MUST exist in the enum's values. Note that this behavior is different from the Schema Object's default keyword, which documents the receiver's behavior rather than inserting the value into the data.
	Default string `json:"default"`
	// An optional description for the server variable. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
}

// Holds a set of reusable objects for different aspects of the OAS. All objects defined within the Components Object will have no effect on the API unless they are explicitly referenced from outside the Components Object.
type Components struct {
	Schemas         map[string]Schema         `json:"schemas,omitzero"`
	Responses       map[string]Response       `json:"responses,omitzero"`
	Parameters      map[string]Parameter      `json:"parameters,omitzero"`
	Examples        map[string]Example        `json:"examples,omitzero"`
	RequestBodies   map[string]RequestBody    `json:"requestBodies,omitzero"`
	Headers         map[string]Header         `json:"headers,omitzero"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitzero"`
	Links           map[string]Link           `json:"links,omitzero"`
	Callbacks       map[string]Callback       `json:"callbacks,omitzero"`
	PathItems       map[string]PathItem       `json:"pathItems,omitzero"`
}

// Holds the relative paths to the individual endpoints and their operations. The path is appended to the URL from the Server Object in order to construct the full URL. The Paths Object MAY be empty, due to Access Control List (ACL) constraints.
type Paths map[string]PathItem

// Describes the operations available on a single path. A Path Item MAY be empty, due to ACL constraints. The path itself is still exposed to the documentation viewer but they will not know which operations and parameters are available.
type PathItem struct {
	// Allows for a referenced definition of this path item. The value MUST be in the form of a URI, and the referenced structure MUST be in the form of a Path Item Object. In case a Path Item Object field appears both in the defined object and the referenced object, the behavior is undefined. See the rules for resolving Relative References.
	Ref string `json:"$ref,omitzero"`
	// An optional string summary, intended to apply to all operations in this path.
	Summary string `json:"summary,omitzero"`
	// An optional string description, intended to apply to all operations in this path. CommonMark syntax MAY be used for rich text representation.
	Description string    `json:"description,omitzero"`
	Get         Operation `json:"get,omitzero"`
	Put         Operation `json:"put,omitzero"`
	Post        Operation `json:"post,omitzero"`
	Delete      Operation `json:"delete,omitzero"`
	Options     Operation `json:"options,omitzero"`
	Head        Operation `json:"head,omitzero"`
	Patch       Operation `json:"patch,omitzero"`
	Trace       Operation `json:"trace,omitzero"`
	// An alternative servers array to service all operations in this path. If a servers array is specified at the OpenAPI Object level, it will be overridden by this value.
	Servers []Server `json:"servers,omitzero"`
	// A list of parameters that are applicable for all the operations described under this path. These parameters can be overridden at the operation level, but cannot be removed there. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined in the OpenAPI Object's components.parameters.
	Parameters []Parameter `json:"parameters,omitzero"`
}

// Describes a single API operation on a path.
type Operation struct {
	// A list of tags for API documentation control. Tags can be used for logical grouping of operations by resources or any other qualifier.
	Tags []string `json:"tags,omitzero"`
	// A short summary of what the operation does.
	Summary string `json:"summary,omitzero"`
	// A verbose explanation of the operation behavior. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
	// Additional external documentation for this operation.
	ExternalDocs ExternalDoc `json:"externalDocs,omitzero"`
	// Unique string used to identify the operation. The id MUST be unique among all operations described in the API. The operationId value is case-sensitive. Tools and libraries MAY use the operationId to uniquely identify an operation, therefore, it is RECOMMENDED to follow common programming naming conventions.
	OperationID string `json:"operationId,omitzero"`
	// A list of parameters that are applicable for this operation. If a parameter is already defined at the Path Item, the new definition will override it but can never remove it. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined in the OpenAPI Object's components.parameters.
	Parameters Parameter `json:"parameters,omitzero"`
	// The request body applicable for this operation. The requestBody is fully supported in HTTP methods where the HTTP 1.1 specification RFC7231 has explicitly defined semantics for request bodies. In other cases where the HTTP spec is vague (such as GET, HEAD and DELETE), requestBody is permitted but does not have well-defined semantics and SHOULD be avoided if possible.
	RequestBody RequestBody `json:"requestBody,omitzero"`
	// The list of possible responses as they are returned from executing this operation.
	Responses Responses `json:"responses,omitzero"`
	// A map of possible out-of band callbacks related to the parent operation. The key is a unique identifier for the Callback Object. Each value in the map is a Callback Object that describes a request that may be initiated by the API provider and the expected responses.
	Callbacks map[string]Callback `json:"callbacks,omitzero"`
	// Declares this operation to be deprecated. Consumers SHOULD refrain from usage of the declared operation. Default value is false.
	Deprecated bool `json:"deprecated,omitzero"`
	// A declaration of which security mechanisms can be used for this operation. The list of values includes alternative Security Requirement Objects that can be used. Only one of the Security Requirement Objects need to be satisfied to authorize a request. To make security optional, an empty security requirement ({}) can be included in the array. This definition overrides any declared top-level security. To remove a top-level security declaration, an empty array can be used.
	Security []SecurityRequirement `json:"security,omitzero"`
	// An alternative servers array to service this operation. If a servers array is specified at the Path Item Object or OpenAPI Object level, it will be overridden by this value.
	Servers []Server `json:"servers,omitzero"`
}

// Allows referencing an external resource for extended documentation.
type ExternalDoc struct {
	// REQUIRED. The URI for the target documentation. This MUST be in the form of a URI.
	URL string `json:"url"`
	// A description of the target documentation. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
}

// Describes a single operation parameter.
type Parameter struct {
	// REQUIRED. The name of the parameter. Parameter names are case sensitive.
	Name string `json:"name"`
	// REQUIRED. The location of the parameter. Possible values are "query", "header", "path" or "cookie".
	In string `json:"in"`
	// A brief description of the parameter. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
	// Determines whether this parameter is mandatory. If the parameter location is "path", this field is REQUIRED and its value MUST be true. Otherwise, the field MAY be included and its default value is false.
	Required bool `json:"required,omitzero"`
	// Specifies that a parameter is deprecated and SHOULD be transitioned out of usage. Default value is false.
	Deprecated bool `json:"deprecated,omitzero"`
	// If true, clients MAY pass a zero-length string value in place of parameters that would otherwise be omitted entirely, which the server SHOULD interpret as the parameter being unused. Default value is false. If style is used, and if behavior is n/a (cannot be serialized), the value of allowEmptyValue SHALL be ignored. Interactions between this field and the parameter's Schema Object are implementation-defined. This field is valid only for query parameters. Use of this field is NOT RECOMMENDED, and it is likely to be removed in a later revision.
	AllowEmptyValue bool `json:"allowEmptyValue,omitzero"`

	// Describes how the parameter value will be serialized depending on the type of the parameter value. Default values (based on value of in): for "query" - "form"; for "path" - "simple"; for "header" - "simple"; for "cookie" - "form".
	Style string `json:"style,omitzero"`
	// When this is true, parameter values of type array or object generate separate parameters for each value of the array or key-value pair of the map. For other types of parameters this field has no effect. When style is "form", the default value is true. For all other styles, the default value is false. Note that despite false being the default for deepObject, the combination of false with deepObject is undefined.
	Explode bool `json:"explode,omitzero"`
	// When this is true, parameter values are serialized using reserved expansion, as defined by RFC6570, which allows RFC3986's reserved character set, as well as percent-encoded triples, to pass through unchanged, while still percent-encoding all other disallowed characters (including % outside of percent-encoded triples). Applications are still responsible for percent-encoding reserved characters that are not allowed in the query string ([, ], #), or have a special meaning in application/x-www-form-urlencoded (-, &, +); see Appendices C and E for details. This field only applies to parameters with an in value of query. The default value is false.
	AllowReserved bool `json:"allowReserved,omitzero"`
	// The schema defining the type used for the parameter.
	Schema Schema `json:"schema,omitzero"`
	// Example of the parameter's potential value; see Working With Examples.
	Example any `json:"example,omitzero"`
	// Examples of the parameter's potential value; see Working With Examples.
	Examples map[string]Example `json:"examples,omitzero"`

	// A map containing the representations for the parameter. The key is the media type and the value describes it. The map MUST only contain one entry.
	Content map[string]MediaType `json:"content,omitzero"`
}

// Describes a single request body.
type RequestBody struct {
	// A brief description of the request body. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
	// REQUIRED. The content of the request body. The key is a media type or media type range and the value describes it. For requests that match multiple keys, only the most specific key is applicable. e.g. "text/plain" overrides "text/*"
	Content map[string]MediaType `json:"content"`
	// Determines if the request body is required in the request. Defaults to false.
	Required bool `json:"required,omitzero"`
}

// Each Media Type Object provides schema and examples for the media type identified by its key.
type MediaType struct {
	// The schema defining the content of the request, response, parameter, or header.
	Schema Schema `json:"schema,omitzero"`
	// Example of the media type.
	Example any `json:"example,omitzero"`
	// Examples of the media type.
	Examples map[string]Example `json:"examples,omitzero"`
	// A map between a property name and its encoding information. The key, being the property name, MUST exist in the schema as a property. The encoding field SHALL only apply to Request Body Objects, and only when the media type is multipart or application/x-www-form-urlencoded. If no Encoding Object is provided for a property, the behavior is determined by the default values documented for the Encoding Object.
	Encoding map[string]Encoding `json:"encoding,omitzero"`
}

// A single encoding definition applied to a single schema property.
type Encoding struct {
	// The Content-Type for encoding a specific property. The value is a comma-separated list, each element of which is either a specific media type (e.g. image/png) or a wildcard media type (e.g. image/*). Default value depends on the property type as shown in the table below.
	ContentType string `json:"contentType,omitzero"`
	// A map allowing additional information to be provided as headers. Content-Type is described separately and SHALL be ignored in this section. This field SHALL be ignored if the request body media type is not a multipart.
	Headers map[string]Header `json:"headers,omitzero"`
	// Describes how a specific property value will be serialized depending on its type. See Parameter Object for details on the style field. The behavior follows the same values as query parameters, including default values. Note that the initial ? used in query strings is not used in application/x-www-form-urlencoded message bodies, and MUST be removed (if using an RFC6570 implementation) or simply not added (if constructing the string manually). This field SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	Style string `json:"style,omitzero"`
	// When this is true, property values of type array or object generate separate parameters for each value of the array, or key-value-pair of the map. For other types of properties this field has no effect. When style is "form", the default value is true. For all other styles, the default value is false. Note that despite false being the default for deepObject, the combination of false with deepObject is undefined. This field SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	Explode bool `json:"explode,omitzero"`
	// When this is true, parameter values are serialized using reserved expansion, as defined by RFC6570, which allows RFC3986's reserved character set, as well as percent-encoded triples, to pass through unchanged, while still percent-encoding all other disallowed characters (including % outside of percent-encoded triples). Applications are still responsible for percent-encoding reserved characters that are not allowed in the query string ([, ], #), or have a special meaning in application/x-www-form-urlencoded (-, &, +); see Appendices C and E for details. The default value is false. This field SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	AllowReserved bool `json:"allowReserved,omitzero"`
}

// A container for the expected responses of an operation. The container maps a HTTP response code to the expected response.
type Responses struct {
	// The documentation of responses other than the ones declared for specific HTTP response codes. Use this field to cover undeclared responses.
	Default                       Response `json:"default,omitzero"`
	Continue                      Response `json:"100,omitzero"`
	SwitchingProtocols            Response `json:"101,omitzero"`
	Processing                    Response `json:"102,omitzero"`
	EarlyHints                    Response `json:"103,omitzero"`
	OK                            Response `json:"200,omitzero"`
	Created                       Response `json:"201,omitzero"`
	Accepted                      Response `json:"202,omitzero"`
	NonAuthoritativeInfo          Response `json:"203,omitzero"`
	NoContent                     Response `json:"204,omitzero"`
	ResetContent                  Response `json:"205,omitzero"`
	PartialContent                Response `json:"206,omitzero"`
	MultiStatus                   Response `json:"207,omitzero"`
	AlreadyReported               Response `json:"208,omitzero"`
	IMUsed                        Response `json:"226,omitzero"`
	MultipleChoices               Response `json:"300,omitzero"`
	MovedPermanently              Response `json:"301,omitzero"`
	Found                         Response `json:"302,omitzero"`
	SeeOther                      Response `json:"303,omitzero"`
	NotModified                   Response `json:"304,omitzero"`
	UseProxy                      Response `json:"305,omitzero"`
	TemporaryRedirect             Response `json:"307,omitzero"`
	PermanentRedirect             Response `json:"308,omitzero"`
	BadRequest                    Response `json:"400,omitzero"`
	Unauthorized                  Response `json:"401,omitzero"`
	PaymentRequired               Response `json:"402,omitzero"`
	Forbidden                     Response `json:"403,omitzero"`
	NotFound                      Response `json:"404,omitzero"`
	MethodNotAllowed              Response `json:"405,omitzero"`
	NotAcceptable                 Response `json:"406,omitzero"`
	ProxyAuthRequired             Response `json:"407,omitzero"`
	RequestTimeout                Response `json:"408,omitzero"`
	Conflict                      Response `json:"409,omitzero"`
	Gone                          Response `json:"410,omitzero"`
	LengthRequired                Response `json:"411,omitzero"`
	PreconditionFailed            Response `json:"412,omitzero"`
	RequestEntityTooLarge         Response `json:"413,omitzero"`
	RequestURITooLong             Response `json:"414,omitzero"`
	UnsupportedMediaType          Response `json:"415,omitzero"`
	RequestedRangeNotSatisfiable  Response `json:"416,omitzero"`
	ExpectationFailed             Response `json:"417,omitzero"`
	Teapot                        Response `json:"418,omitzero"`
	MisdirectedRequest            Response `json:"421,omitzero"`
	UnprocessableEntity           Response `json:"422,omitzero"`
	Locked                        Response `json:"423,omitzero"`
	FailedDependency              Response `json:"424,omitzero"`
	TooEarly                      Response `json:"425,omitzero"`
	UpgradeRequired               Response `json:"426,omitzero"`
	PreconditionRequired          Response `json:"428,omitzero"`
	TooManyRequests               Response `json:"429,omitzero"`
	RequestHeaderFieldsTooLarge   Response `json:"431,omitzero"`
	UnavailableForLegalReasons    Response `json:"451,omitzero"`
	InternalServerError           Response `json:"500,omitzero"`
	NotImplemented                Response `json:"501,omitzero"`
	BadGateway                    Response `json:"502,omitzero"`
	ServiceUnavailable            Response `json:"503,omitzero"`
	GatewayTimeout                Response `json:"504,omitzero"`
	HTTPVersionNotSupported       Response `json:"505,omitzero"`
	VariantAlsoNegotiates         Response `json:"506,omitzero"`
	InsufficientStorage           Response `json:"507,omitzero"`
	LoopDetected                  Response `json:"508,omitzero"`
	NotExtended                   Response `json:"510,omitzero"`
	NetworkAuthenticationRequired Response `json:"511,omitzero"`
}

// Describes a single response from an API operation, including design-time, static links to operations based on the response.
type Response struct {
	// REQUIRED. A description of the response. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description"`
	// Maps a header name to its definition. RFC7230 states header names are case insensitive. If a response header is defined with the name "Content-Type", it SHALL be ignored.
	Headers map[string]Header `json:"headers,omitzero"`
	// A map containing descriptions of potential response payloads. The key is a media type or media type range and the value describes it. For responses that match multiple keys, only the most specific key is applicable. e.g. "text/plain" overrides "text/*"
	Content map[string]MediaType `json:"content,omitzero"`
	// A map of operations links that can be followed from the response. The key of the map is a short name for the link, following the naming constraints of the names for Component Objects.
	Links map[string]Link `json:"links,omitzero"`
}

// A map of possible out-of band callbacks related to the parent operation. Each value in the map is a Path Item Object that describes a set of requests that may be initiated by the API provider and the expected responses. The key value used to identify the Path Item Object is an expression, evaluated at runtime, that identifies a URL to use for the callback operation.
type Callback map[string]PathItem

// An object grouping an internal or external example value with basic summary and description metadata. This object is typically used in fields named examples (plural), and is a referenceable alternative to older example (singular) fields that do not support referencing or metadata.
type Example struct {
	// Short description for the example.
	Summary string `json:"summary,omitzero"`
	// Long description for the example. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
	// Embedded literal example. The value field and externalValue field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON or YAML, use a string value to contain the example, escaping where necessary.
	Value any `json:"value,omitzero"`
	// A URI that identifies the literal example. This provides the capability to reference examples that cannot easily be included in JSON or YAML documents. The value field and externalValue field are mutually exclusive. See the rules for resolving Relative References.
	ExternalValue string `json:"externalValue,omitzero"`
}

// The Link Object represents a possible design-time link for a response. The presence of a link does not guarantee the caller's ability to successfully invoke it, rather it provides a known relationship and traversal mechanism between responses and other operations.
type Link struct {
	// A URI reference to an OAS operation. This field is mutually exclusive of the operationId field, and MUST point to an Operation Object. Relative operationRef values MAY be used to locate an existing Operation Object in the OpenAPI Description.
	OperationRef string `json:"operationRef,omitzero"`
	// The name of an existing, resolvable OAS operation, as defined with a unique operationId. This field is mutually exclusive of the operationRef field.
	OperationID string `json:"operationId,omitzero"`
	// A map representing parameters to pass to an operation as specified with operationId or identified via operationRef. The key is the parameter name to be used (optionally qualified with the parameter location, e.g. path.id for an id parameter in the path), whereas the value can be a constant or an expression to be evaluated and passed to the linked operation.
	Parameters map[string]any `json:"parameters,omitzero"`
	// A literal value or {expression} to use as a request body when calling the target operation.
	RequestBody any `json:"requestBody,omitzero"`
	// A description of the link. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
	// A server object to be used by the target operation.
	Server Server `json:"server,omitzero"`
}

// Describes a single header for HTTP responses and for individual parts in multipart representations; see the relevant Response Object and Encoding Object documentation for restrictions on which headers can be described.
type Header struct {
	// A brief description of the header. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
	// Determines whether this header is mandatory. The default value is false.
	Required bool `json:"required,omitzero"`
	// Specifies that the header is deprecated and SHOULD be transitioned out of usage. Default value is false.
	Deprecated bool `json:"deprecated,omitzero"`

	// Describes how the header value will be serialized. The default (and only legal value for headers) is "simple".
	Style string `json:"style,omitzero"`
	// When this is true, header values of type array or object generate a single header whose value is a comma-separated list of the array items or key-value pairs of the map, see Style Examples. For other data types this field has no effect. The default value is false.
	Explode bool `json:"explode,omitzero"`
	// The schema defining the type used for the header.
	Schema Schema `json:"schema,omitzero"`
	// Example of the header's potential value; see Working With Examples.
	Example any `json:"example,omitzero"`
	// Examples of the header's potential value; see Working With Examples.
	Examples map[string]Example `json:"examples,omitzero"`

	// A map containing the representations for the header. The key is the media type and the value describes it. The map MUST only contain one entry.
	Content map[string]MediaType `json:"content,omitzero"`
}

// Adds metadata to a single tag that is used by the Operation Object. It is not mandatory to have a Tag Object per tag defined in the Operation Object instances.
type Tag struct {
	// REQUIRED. The name of the tag.
	Name string `json:"name"`
	// A description for the tag. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
	// Additional external documentation for this tag.
	ExternalDocs ExternalDoc `json:"externalDocs,omitzero"`
}

// A simple object to allow referencing other components in the OpenAPI Description, internally and externally.
type Reference struct {
	// A simple object to allow referencing other components in the OpenAPI Description, internally and externally.
	Ref string `json:"$ref,omitzero"`
	// A short summary which by default SHOULD override that of the referenced component. If the referenced object-type does not allow a summary field, then this field has no effect.
	Summary string `json:"summary,omitzero"`
	// A description which by default SHOULD override that of the referenced component. CommonMark syntax MAY be used for rich text representation. If the referenced object-type does not allow a description field, then this field has no effect.
	Description string `json:"description,omitzero"`
}

type Schema = any

// Defines a security scheme that can be used by the operations.
type SecurityScheme struct {
	// REQUIRED. The type of the security scheme. Valid values are "apiKey", "http", "mutualTLS", "oauth2", "openIdConnect".
	Type string `json:"type"`
	// A description for security scheme. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitzero"`
	// REQUIRED. The name of the header, query or cookie parameter to be used.
	Name string `json:"name"`
	// REQUIRED. The location of the API key. Valid values are "query", "header", or "cookie".
	In string `json:"in"`
	// REQUIRED. The name of the HTTP Authentication scheme to be used in the Authorization header as defined in RFC7235. The values used SHOULD be registered in the IANA Authentication Scheme registry. The value is case-insensitive, as defined in RFC7235.
	Scheme string `json:"scheme"`
	// A hint to the client to identify how the bearer token is formatted. Bearer tokens are usually generated by an authorization server, so this information is primarily for documentation purposes.
	BearerFormat string `json:"bearerFormat,omitzero"`
	// REQUIRED. An object containing configuration information for the flow types supported.
	Flows OAuthFlows `json:"flows,omitzero"`
	// REQUIRED. Well-known URL to discover the [[OpenID-Connect-Discovery]] provider metadata.
	OpenIDConnectURL string `json:"openIdConnectUrl,omitzero"`
}

// Allows configuration of the supported OAuth Flows.
type OAuthFlows struct {
	// Configuration for the OAuth Implicit flow
	Implicit OAuthFlow `json:"implicit,omitzero"`
	// Configuration for the OAuth Resource Owner Password flow
	Password OAuthFlow `json:"password,omitzero"`
	// Configuration for the OAuth Client Credentials flow. Previously called application in OpenAPI 2.0.
	ClientCredentials OAuthFlow `json:"clientCredentials,omitzero"`
	// Configuration for the OAuth Authorization Code flow. Previously called accessCode in OpenAPI 2.0.
	AuthorizationCode OAuthFlow `json:"authorizationCode,omitzero"`
}

// Configuration details for a supported OAuth Flow.
type OAuthFlow struct {
	// REQUIRED. The authorization URL to be used for this flow. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
	AuthorizationURL string `json:"authorizationUrl"`
	// REQUIRED. The token URL to be used for this flow. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
	TokenURL string `json:"tokenUrl"`
	// The URL to be used for obtaining refresh tokens. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
	RefreshURL string `json:"refreshUrl,omitzero"`
	// REQUIRED. The available scopes for the OAuth2 security scheme. A map between the scope name and a short description for it. The map MAY be empty.
	Scopes map[string]string `json:"scopes"`
}

// Lists the required security schemes to execute this operation. The name used for each property MUST correspond to a security scheme declared in the Security Schemes under the Components Object.
type SecurityRequirement map[string][]string
