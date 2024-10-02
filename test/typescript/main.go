package main

import (
	"context"
	"dagger/test-instant/internal/dagger"
)

type TestInstant struct{
		RequiredPaths []string
}

func (t *TestInstant) ModuleRuntime(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.Container, error) {
	return dag.Container().From("ttl.sh/dagger/instant-module/typescript/v0@sha256:8afa6ad4e99d4449a6d27f90d2867503731259792036558b15e6cffd87eb1add"), nil
}

func (t *TestInstant) Codegen(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.GeneratedCode, error) {
	return dag.GeneratedCode(dag.Directory()), nil
}
