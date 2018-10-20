package tddbc8

import "testing"

func TestNewInstance(t *testing.T) {
	sut := NewPurelyImaginaryNumber(3)
	if sut.imaginaryPart != 3 {
		t.Fatalf("虚部に３を渡したときに虚部３を持つ純虚数が生成できない")
	}
}
