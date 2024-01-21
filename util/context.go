package util

import "context"

type contextKey string

const (
	ContextKeySortDir   = contextKey("sort")
	ContextKeyCurrMonth = contextKey("currMonth")
)

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

func GetCurrMonthCtx(ctx context.Context) string {
	month, ok := ctx.Value(ContextKeyCurrMonth).(string)
	if !ok {
		return GetCurrentMonth()
	}

	return month
}

func WithCurrMonthCtx(ctx context.Context, month string) context.Context {
	return context.WithValue(ctx, ContextKeyCurrMonth, month)
}
