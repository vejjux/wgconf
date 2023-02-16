package exx

import "errors"

func C(err error) {
	if err != nil {
		panic(err)
	}
}

func CA[T any](val T, err error) T {
	C(err)

	return val
}

func CA2[T1, T2 any](val1 T1, val2 T2, err error) (T1, T2) {
	C(err)

	return val1, val2
}

func CI(cond bool, msg string) {
	if cond {
		C(errors.New(msg))
	}
}
