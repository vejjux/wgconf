package exx

import (
	"fmt"
)

func H(err *error) {
	if e := recover(); e != nil {
		if recovered, ok := e.(error); ok {
			*err = recovered
		} else {
			*err = fmt.Errorf("%v", e)
		}
	}
}

func HL() {
	if e := recover(); e != nil {
		fmt.Println(e)
	}
}
