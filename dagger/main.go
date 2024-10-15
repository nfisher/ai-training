package main

import (
	"context"
	"dagger.io/dagger"
	"os"
	"time"
)

func main() {
	d := time.Now().Add(time.Hour * 2)
	ctx, _ := context.WithDeadline(context.Background(), d)

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
