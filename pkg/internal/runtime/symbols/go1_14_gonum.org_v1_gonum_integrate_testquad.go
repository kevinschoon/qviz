// Code generated by 'goexports gonum.org/v1/gonum/integrate/testquad'. DO NOT EDIT.

// +build go1.14,!go1.15

package symbols

import (
	"gonum.org/v1/gonum/integrate/testquad"
	"reflect"
)

func init() {
	Symbols["gonum.org/v1/gonum/integrate/testquad"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Constant":       reflect.ValueOf(testquad.Constant),
		"ExpOverX2Plus1": reflect.ValueOf(testquad.ExpOverX2Plus1),
		"Poly":           reflect.ValueOf(testquad.Poly),
		"Sin":            reflect.ValueOf(testquad.Sin),
		"Sqrt":           reflect.ValueOf(testquad.Sqrt),
		"XExpMinusX":     reflect.ValueOf(testquad.XExpMinusX),

		// type definitions
		"Integral": reflect.ValueOf((*testquad.Integral)(nil)),
	}
}
