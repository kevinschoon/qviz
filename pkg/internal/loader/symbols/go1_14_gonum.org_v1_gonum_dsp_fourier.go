// Code generated by 'goexports gonum.org/v1/gonum/dsp/fourier'. DO NOT EDIT.

// +build go1.14,!go1.15

package symbols

import (
	"gonum.org/v1/gonum/dsp/fourier"
	"reflect"
)

func init() {
	Symbols["gonum.org/v1/gonum/dsp/fourier"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewCmplxFFT":       reflect.ValueOf(fourier.NewCmplxFFT),
		"NewDCT":            reflect.ValueOf(fourier.NewDCT),
		"NewDST":            reflect.ValueOf(fourier.NewDST),
		"NewFFT":            reflect.ValueOf(fourier.NewFFT),
		"NewQuarterWaveFFT": reflect.ValueOf(fourier.NewQuarterWaveFFT),

		// type definitions
		"CmplxFFT":       reflect.ValueOf((*fourier.CmplxFFT)(nil)),
		"DCT":            reflect.ValueOf((*fourier.DCT)(nil)),
		"DST":            reflect.ValueOf((*fourier.DST)(nil)),
		"FFT":            reflect.ValueOf((*fourier.FFT)(nil)),
		"QuarterWaveFFT": reflect.ValueOf((*fourier.QuarterWaveFFT)(nil)),
	}
}
