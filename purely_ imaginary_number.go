package tddbc8

import "errors"

type PurelyImaginaryNumber struct {
	imaginaryPart int
}

func NewPurelyImaginaryNumber(i int) (*PurelyImaginaryNumber, error) {
	return &PurelyImaginaryNumber{i}, errors.New("dummy")
}
