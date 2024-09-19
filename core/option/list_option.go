package option

import "github.com/hongyuxuan/bkcmdb-sdk-go/types"

type ListOptionFunc func(*types.ListOption)

func WithStart(start int) ListOptionFunc {
	return func(l *types.ListOption) {
		l.Page.Start = start
	}
}

func WithLimit(limit int) ListOptionFunc {
	return func(l *types.ListOption) {
		l.Page.Limit = limit
	}
}

func WithSort(sort string) ListOptionFunc {
	return func(l *types.ListOption) {
		l.Page.Sort = sort
	}
}

func WithFields(fields []string) ListOptionFunc {
	return func(l *types.ListOption) {
		l.Fields = fields
	}
}

func WithCondition(cond *types.Conditions) ListOptionFunc {
	return func(l *types.ListOption) {
		l.Conditions = cond
	}
}
