package tddbc8

import "testing"

func Test虚部に３を渡したときに虚部３を持つ純虚数が生成できること(t *testing.T) {
	sut := NewPurelyImaginaryNumber(3)
	if sut.imaginaryPart != 3 {
		t.Fatalf("虚部に３を渡したときに虚部３を持つ純虚数が生成できない")
	}
}

func Test虚部に10を渡したときに虚部10を持つ純虚数が生成できること(t *testing.T) {
	sut := NewPurelyImaginaryNumber(10)
	if sut.imaginaryPart != 10 {
		t.Fatalf("虚部に10を渡したときに虚部10を持つ純虚数が生成できない")
	}
}