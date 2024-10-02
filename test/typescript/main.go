package main

import (
	"context"
	"dagger/test-instant/internal/dagger"
)

type TestInstant struct{
		RequiredPaths []string
}

func (t *TestInstant) ModuleRuntime(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.Container, error) {
	return dag.Container().From("ttl.sh/dagger/instant-module/bun/v0@sha256:9e1e90948ab8f0110b84b5b9b6f4f583db73e344d5800dadd09fa22b6d222b5f").WithExec([]string{"ls"}), nil
}

func (t *TestInstant) Codegen(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.GeneratedCode, error) {
	return dag.GeneratedCode(dag.Directory()), nil
}
