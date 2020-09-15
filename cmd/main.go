package main

import (
	"context"
	"fmt"
	"github.com/mhickman/go-awair/awair"
	"golang.org/x/oauth2"
	"os"
)

const awareBearerToken = "BEARER_TOKEN"

var bearerToken string

func init() {
	bearerToken = os.Getenv(awareBearerToken)
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: bearerToken,
		},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := awair.NewClient(tc)

	devices, _, err := client.Devices.List(ctx)

	fmt.Printf("%+v\n", devices)
	fmt.Println(err)
}
