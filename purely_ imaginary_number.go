package tddbc8

import (
	"errors"
	"strconv"
)

type PurelyImaginaryNumber struct {
	imaginaryPart int
}

func NewPurelyImaginaryNumber(i int) (*PurelyImaginaryNumber, error) {
	if i == 0 {
		return nil, errors.New("虚数に０を設定できません")
	}
	return &PurelyImaginaryNumber{i}, nil
}

func (p *PurelyImaginaryNumber) Notation() string {
	if p.imaginaryPart == 1 {
		return "i"
	} else if p.imaginaryPart == (-1) {
		return "-i"
	}

	return strconv.Itoa(p.imaginaryPart) + "i"
}
