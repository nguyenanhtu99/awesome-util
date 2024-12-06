package clientx

import (
	"fmt"

	"go.temporal.io/sdk/client"
)

var cli client.Client

func GetClient() client.Client {
	return cli
}

func Init() (client.Client, error) {
	var err error
	cli, err = client.Dial(client.Options{
		HostPort: "127.0.0.1:7233",
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create Temporal client: %w", err)
	}

	return cli, nil
}
