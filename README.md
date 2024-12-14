# nobitex-go
Unofficial Golang SDK for Nobitex.ir

## Examples
Check out the examples in the [examples](examples/) directory.

## Getting Started
```golang
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/S4eedb/nobitex-go/sdk"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

func main() {
	basicAuth, err := securityprovider.NewSecurityProviderBearerToken("my_token")
	if err != nil {
		log.Fatal(err)
	}

	client, err := sdk.NewClient("https://api.nobitex.ir/", sdk.WithRequestEditorFn(basicAuth.Intercept))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.GetV2OrderbookSymbol(context.TODO(), "BTCIRT")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Decode and pretty print the JSON response
	var jsonResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(jsonResponse, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON to pretty format: %v", err)
	}

	fmt.Println(string(prettyJSON))
}

```