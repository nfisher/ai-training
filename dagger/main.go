package main

import (
	"context"
	"dagger.io/dagger"
	"os"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	_, err = client.Host().
		Directory(".").
		DockerBuild(dagger.DirectoryDockerBuildOpts{Platform: "linux/amd64"}).
		Publish(ctx, "nrfisher/pytorch-cuda:latest")
	if err != nil {
		panic(err)
	}
}
