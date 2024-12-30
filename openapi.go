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

type Path struct {
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
