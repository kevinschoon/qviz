// Code generated by 'goexports github.com/tobgu/qframe'. DO NOT EDIT.

// +build go1.14,!go1.15

package symbols

import (
	"github.com/tobgu/qframe"
	"reflect"
)

func init() {
	Symbols["github.com/tobgu/qframe"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"And":      reflect.ValueOf(qframe.And),
		"Doc":      reflect.ValueOf(qframe.Doc),
		"Expr":     reflect.ValueOf(qframe.Expr),
		"New":      reflect.ValueOf(qframe.New),
		"Not":      reflect.ValueOf(qframe.Not),
		"Null":     reflect.ValueOf(qframe.Null),
		"Or":       reflect.ValueOf(qframe.Or),
		"ReadCSV":  reflect.ValueOf(qframe.ReadCSV),
		"ReadJSON": reflect.ValueOf(qframe.ReadJSON),
		"ReadSQL":  reflect.ValueOf(qframe.ReadSQL),
		"Val":      reflect.ValueOf(qframe.Val),

		// type definitions
		"Aggregation":  reflect.ValueOf((*qframe.Aggregation)(nil)),
		"AndClause":    reflect.ValueOf((*qframe.AndClause)(nil)),
		"BoolView":     reflect.ValueOf((*qframe.BoolView)(nil)),
		"ConstBool":    reflect.ValueOf((*qframe.ConstBool)(nil)),
		"ConstFloat":   reflect.ValueOf((*qframe.ConstFloat)(nil)),
		"ConstInt":     reflect.ValueOf((*qframe.ConstInt)(nil)),
		"ConstString":  reflect.ValueOf((*qframe.ConstString)(nil)),
		"EnumView":     reflect.ValueOf((*qframe.EnumView)(nil)),
		"Expression":   reflect.ValueOf((*qframe.Expression)(nil)),
		"Filter":       reflect.ValueOf((*qframe.Filter)(nil)),
		"FilterClause": reflect.ValueOf((*qframe.FilterClause)(nil)),
		"FloatView":    reflect.ValueOf((*qframe.FloatView)(nil)),
		"GroupStats":   reflect.ValueOf((*qframe.GroupStats)(nil)),
		"Grouper":      reflect.ValueOf((*qframe.Grouper)(nil)),
		"Instruction":  reflect.ValueOf((*qframe.Instruction)(nil)),
		"IntView":      reflect.ValueOf((*qframe.IntView)(nil)),
		"NotClause":    reflect.ValueOf((*qframe.NotClause)(nil)),
		"NullClause":   reflect.ValueOf((*qframe.NullClause)(nil)),
		"OrClause":     reflect.ValueOf((*qframe.OrClause)(nil)),
		"Order":        reflect.ValueOf((*qframe.Order)(nil)),
		"QFrame":       reflect.ValueOf((*qframe.QFrame)(nil)),
		"StringView":   reflect.ValueOf((*qframe.StringView)(nil)),

		// interface wrapper definitions
		"_Expression":   reflect.ValueOf((*_github_com_tobgu_qframe_Expression)(nil)),
		"_FilterClause": reflect.ValueOf((*_github_com_tobgu_qframe_FilterClause)(nil)),
	}
}

// _github_com_tobgu_qframe_Expression is an interface wrapper for Expression type
type _github_com_tobgu_qframe_Expression struct {
	WErr func() error
}

func (W _github_com_tobgu_qframe_Expression) Err() error { return W.WErr() }

// _github_com_tobgu_qframe_FilterClause is an interface wrapper for FilterClause type
type _github_com_tobgu_qframe_FilterClause struct {
	WErr    func() error
	WString func() string
}

func (W _github_com_tobgu_qframe_FilterClause) Err() error     { return W.WErr() }
func (W _github_com_tobgu_qframe_FilterClause) String() string { return W.WString() }
