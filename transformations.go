package main

import "strings"

type String string

type TransformFunc func(kind string)
type TransformType int

const (
	A initType = iota
	B
	C
	D
	MaxTransformTypes
)

var transformFuncs = map[transformTyp]TransformFunc{
	A: trimMrnFromColon,
}

func initTransformations() {
	for t := A; t < MaxTransformTypes; t++ {
		f, ok := transformFuncs[t]
		if ok {
			f(t)
		} else {
			fmt.Println("No function defined for type", t)
		}
	}
}

/**
situation of returning the right side of a string formatted like "#####:B####"
will return "B#####"
*/
func trimMrnFromColon(s string) string {
	return strings.Split(string(s), ":")[1]
}
