// Code generated by 'goexports gonum.org/v1/gonum/graph/encoding/graphql'. DO NOT EDIT.

// +build go1.14,!go1.15

package symbols

import (
	"gonum.org/v1/gonum/graph/encoding/graphql"
	"reflect"
)

func init() {
	Symbols["gonum.org/v1/gonum/graph/encoding/graphql"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Unmarshal": reflect.ValueOf(graphql.Unmarshal),

		// type definitions
		"LabelSetter":    reflect.ValueOf((*graphql.LabelSetter)(nil)),
		"StringIDSetter": reflect.ValueOf((*graphql.StringIDSetter)(nil)),

		// interface wrapper definitions
		"_LabelSetter":    reflect.ValueOf((*_gonum_org_v1_gonum_graph_encoding_graphql_LabelSetter)(nil)),
		"_StringIDSetter": reflect.ValueOf((*_gonum_org_v1_gonum_graph_encoding_graphql_StringIDSetter)(nil)),
	}
}

// _gonum_org_v1_gonum_graph_encoding_graphql_LabelSetter is an interface wrapper for LabelSetter type
type _gonum_org_v1_gonum_graph_encoding_graphql_LabelSetter struct {
	WSetLabel func(a0 string)
}

func (W _gonum_org_v1_gonum_graph_encoding_graphql_LabelSetter) SetLabel(a0 string) { W.WSetLabel(a0) }

// _gonum_org_v1_gonum_graph_encoding_graphql_StringIDSetter is an interface wrapper for StringIDSetter type
type _gonum_org_v1_gonum_graph_encoding_graphql_StringIDSetter struct {
	WSetIDFromString func(uid string) error
}

func (W _gonum_org_v1_gonum_graph_encoding_graphql_StringIDSetter) SetIDFromString(uid string) error {
	return W.WSetIDFromString(uid)
}
