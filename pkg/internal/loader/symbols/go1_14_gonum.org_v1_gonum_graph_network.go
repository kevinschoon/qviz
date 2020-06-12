// Code generated by 'goexports gonum.org/v1/gonum/graph/network'. DO NOT EDIT.

// +build go1.14,!go1.15

package symbols

import (
	"gonum.org/v1/gonum/graph/network"
	"reflect"
)

func init() {
	Symbols["gonum.org/v1/gonum/graph/network"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Betweenness":             reflect.ValueOf(network.Betweenness),
		"BetweennessWeighted":     reflect.ValueOf(network.BetweennessWeighted),
		"Closeness":               reflect.ValueOf(network.Closeness),
		"Diffuse":                 reflect.ValueOf(network.Diffuse),
		"DiffuseToEquilibrium":    reflect.ValueOf(network.DiffuseToEquilibrium),
		"EdgeBetweenness":         reflect.ValueOf(network.EdgeBetweenness),
		"EdgeBetweennessWeighted": reflect.ValueOf(network.EdgeBetweennessWeighted),
		"Farness":                 reflect.ValueOf(network.Farness),
		"HITS":                    reflect.ValueOf(network.HITS),
		"Harmonic":                reflect.ValueOf(network.Harmonic),
		"PageRank":                reflect.ValueOf(network.PageRank),
		"PageRankSparse":          reflect.ValueOf(network.PageRankSparse),
		"Residual":                reflect.ValueOf(network.Residual),

		// type definitions
		"HubAuthority": reflect.ValueOf((*network.HubAuthority)(nil)),
	}
}
