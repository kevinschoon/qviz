// Code generated by 'goexports gonum.org/v1/gonum/mathext/prng'. DO NOT EDIT.

// +build go1.14,!go1.15

package symbols

import (
	"gonum.org/v1/gonum/mathext/prng"
	"reflect"
)

func init() {
	Symbols["gonum.org/v1/gonum/mathext/prng"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewMT19937":            reflect.ValueOf(prng.NewMT19937),
		"NewMT19937_64":         reflect.ValueOf(prng.NewMT19937_64),
		"NewSplitMix64":         reflect.ValueOf(prng.NewSplitMix64),
		"NewXoshiro256plus":     reflect.ValueOf(prng.NewXoshiro256plus),
		"NewXoshiro256plusplus": reflect.ValueOf(prng.NewXoshiro256plusplus),
		"NewXoshiro256starstar": reflect.ValueOf(prng.NewXoshiro256starstar),

		// type definitions
		"MT19937":            reflect.ValueOf((*prng.MT19937)(nil)),
		"MT19937_64":         reflect.ValueOf((*prng.MT19937_64)(nil)),
		"SplitMix64":         reflect.ValueOf((*prng.SplitMix64)(nil)),
		"Xoshiro256plus":     reflect.ValueOf((*prng.Xoshiro256plus)(nil)),
		"Xoshiro256plusplus": reflect.ValueOf((*prng.Xoshiro256plusplus)(nil)),
		"Xoshiro256starstar": reflect.ValueOf((*prng.Xoshiro256starstar)(nil)),
	}
}
