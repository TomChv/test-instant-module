package main

import (
	"context"
	"dagger/test-instant/internal/dagger"
)

type TestInstant struct{}

func (t *TestInstant) ModuleRuntime(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.Container, error) {
	return dag.Container().From("ttl.sh/dagger/instant-module/v0@sha256:0574ea0dd7af650b370862173f168ddd58b209c116ca7492a1b5a334d0f3f222"), nil
}

func (t *TestInstant) Codegen(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.GeneratedCode, error) {
	return dag.GeneratedCode(dag.Directory()), nil
}
