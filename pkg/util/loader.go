package util

import (
	"fmt"
	"io/ioutil"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

func Read(path string) error {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = i.Eval(string(raw))
	if err != nil {
		return err
	}
	v, err := i.Eval("New")
	if err != nil {
		return err
	}
	fn := v.Interface().(func() (string, error))
	fmt.Println(fn())
	return nil
}
