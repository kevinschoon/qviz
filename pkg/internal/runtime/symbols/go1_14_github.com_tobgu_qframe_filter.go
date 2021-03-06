// Code generated by 'goexports github.com/tobgu/qframe/filter'. DO NOT EDIT.

// +build go1.14,!go1.15

package symbols

import (
	"github.com/tobgu/qframe/filter"
	"reflect"
)

func init() {
	Symbols["github.com/tobgu/qframe/filter"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Eq":        reflect.ValueOf(filter.Eq),
		"Gt":        reflect.ValueOf(filter.Gt),
		"Gte":       reflect.ValueOf(filter.Gte),
		"In":        reflect.ValueOf(filter.In),
		"Inverse":   reflect.ValueOf(&filter.Inverse).Elem(),
		"IsNotNull": reflect.ValueOf(filter.IsNotNull),
		"IsNull":    reflect.ValueOf(filter.IsNull),
		"Lt":        reflect.ValueOf(filter.Lt),
		"Lte":       reflect.ValueOf(filter.Lte),
		"Neq":       reflect.ValueOf(filter.Neq),
		"Nin":       reflect.ValueOf(filter.Nin),

		// type definitions
		"Filter": reflect.ValueOf((*filter.Filter)(nil)),
	}
}
