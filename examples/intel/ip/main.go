// intel domain lookup is an example of how to use the lookup method
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pangeacyber/go-pangea/pangea"
	"github.com/pangeacyber/go-pangea/service/ip_intel"
)

func main() {
	token := os.Getenv("INTEL_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}

	configID := os.Getenv("INTEL_IP_CONFIG_ID")
	if token == "" {
		log.Fatal("Configuration: No config ID present")
	}

	intelcli := ip_intel.New(&pangea.Config{
		Token:    token,
		Domain:   os.Getenv("PANGEA_DOMAIN"),
		ConfigID: configID,
	})

	ctx := context.Background()
	input := &ip_intel.IpLookupInput{
		Ip:       "93.231.182.110",
		Raw:      true,
		Verbose:  true,
		Provider: "crowdstrike",
	}

	response, err := intelcli.Lookup(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pangea.Stringify(response.Result))
}
