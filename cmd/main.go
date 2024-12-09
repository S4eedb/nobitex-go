package main

import (
	"context"
	"fmt"
	"log"

	"github.com/S4eedb/nobitex-go/internal/api"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

func main() {
	basicAuth, err := securityprovider.NewSecurityProviderBasicAuth("my_user", "my_pass")
	if err != nil {
		log.Fatal(err)
	}

	client, err := api.NewClient("https://api.nobitex.ir/", api.WithRequestEditorFn(basicAuth.Intercept))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.GetV2OrderbookSymbol(context.TODO(), "BTCIRT")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp.StatusCode: %v\n", resp.StatusCode)
}
