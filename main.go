// A generated module for InstantModule functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/instant-module/internal/dagger"
)

type InstantModule struct{}

// Returns a container that echoes whatever string argument is provided
func (m *InstantModule) CompileGo(ctx context.Context, source *dagger.Directory) (string, error) {
	bin := dag.Container().
		From("golang:1.23.1-alpine").
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"go", "mod", "download"}).
		WithExec([]string{"go", "build", "-o", "/bin/entrypoint"}).
		File("/bin/entrypoint")

	return dag.Container().
		From("alpine").
		WithFile("/bin/entrypoint", bin).
		WithEntrypoint([]string{"/bin/entrypoint"}).
		Publish(ctx, "ttl.sh/dagger/instant-module/v0")
}


func (m *InstantModule) CompileTS(ctx context.Context, source *dagger.Directory) (string, error) {
	return dag.Container().
		From("node:20-alpine").
		WithExec([]string{"npm", "install", "-g", "tsx@4.15.6"}).
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"yarn"}).
		WithEntrypoint([]string{"tsx", "./src/entrypoint.ts"}).
		Publish(ctx, "ttl.sh/dagger/instant-module/typescript/v0")
}

func (m *InstantModule) CompileBun(ctx context.Context, source *dagger.Directory) (string, error) {
	return dag.Container().
		From("oven/bun:1.1.26-alpine").
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"bun", "install"}).
		WithEntrypoint([]string{"bun", "src/entrypoint.ts"}).
		Publish(ctx, "ttl.sh/dagger/instant-module/bun/v0")
}