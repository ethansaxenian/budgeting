package util

import (
	"context"

	"github.com/ethansaxenian/budgeting/types"
)

type contextKey string

const (
	ContextKeySortDir         = contextKey("sort")
	ContextKeyCurrMonth       = contextKey("currMonth")
	ContextKeyTransactionType = contextKey("transactionType")
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

func GetTransactionTypeCtx(ctx context.Context) string {
	transactionType, ok := ctx.Value(ContextKeyTransactionType).(string)
	if !ok {
		return string(types.INCOME)
	}

	return transactionType
}

func WithTransactionTypeCtx(ctx context.Context, transactionType string) context.Context {
	return context.WithValue(ctx, ContextKeyTransactionType, transactionType)
}
