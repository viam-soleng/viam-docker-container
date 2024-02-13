package main

import (
	"context"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/utils/rpc"
)

func main() {
	logger := logging.NewDebugLogger("client")
	robot, err := client.New(
		context.Background(),
		"<SMART-MACHINE-URL>",
		logger,
		client.WithDialOptions(rpc.WithEntityCredentials(
			// Replace "<API-KEY-ID>" (including brackets) with your machine's api key id
			"<API-KEY-ID>",
			rpc.Credentials{
				Type: rpc.CredentialsTypeAPIKey,
				// Replace "<API-KEY>" (including brackets) with your machine's api key
				Payload: "<API-KEY>",
			})),
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer robot.Close(context.Background())
	logger.Info("Resources:")
	logger.Info(robot.ResourceNames())

}
