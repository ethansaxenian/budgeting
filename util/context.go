package util

import (
	"context"
)

type contextKey string

const ContextKeySortDir = contextKey("sort")

const (
	ContextValueSortDirAsc  = "Asc"
	ContextValueSortDirDesc = "Desc"
)

func GetNextSortCtx(ctx context.Context) string {
	dir, ok := ctx.Value(ContextKeySortDir).(string)
	if !ok {
		return "Asc"
	}

	if dir == "Desc" {
		return "Asc"
	}

	return "Desc"
}

func WithNextSortCtx(ctx context.Context, dir string) context.Context {
	return context.WithValue(ctx, ContextKeySortDir, dir)
}
