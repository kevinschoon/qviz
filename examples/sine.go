package main

import qviz "github.com/kevinschoon/qviz/pkg"

func New() (*qviz.Viz, error) {
	return &qviz.Viz{
		Labels: []string{"fuu"},
	}, nil
}
