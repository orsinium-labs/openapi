package openapi_test

import (
	"encoding/json"
	"fmt"

	"github.com/orsinium-labs/openapi"
)

func Example() {
	docs := openapi.OpenAPI{
		Version: "3.0.2",
		Info: openapi.Info{
			Title: "Cool service",
		},
		Paths: openapi.Paths{
			"/echo": openapi.PathItem{
				Post: openapi.Operation{
					Summary: "Scream into the void",
				},
			},
		},
	}
	jsonDocs, err := json.Marshal(docs)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonDocs)
}
