package tddbc8

import "testing"

func Test虚部に３を渡したときに虚部３を持つ純虚数が生成できること(t *testing.T) {
	sut, err := NewPurelyImaginaryNumber(3)
	if err != nil {
		t.Fatalf("エラーが返ってきている err = %v", err)
	}
	if sut.imaginaryPart != 3 {
		t.Fatalf("虚部に３を渡したときに虚部３を持つ純虚数が生成できない")
	}
}

func Test虚部に10を渡したときに虚部10を持つ純虚数が生成できること(t *testing.T) {
	sut, err := NewPurelyImaginaryNumber(10)
	if err != nil {
		t.Fatalf("エラーが返ってきている err = %v", err)
	}
	if sut.imaginaryPart != 10 {
		t.Fatalf("虚部に10を渡したときに虚部10を持つ純虚数が生成できない")
	}
}

func Test虚部に0渡したときにエラーを返すこと(t *testing.T) {
	_, err := NewPurelyImaginaryNumber(0)
	if err == nil {
		t.Fatalf("エラーが返ってこない")
	}
}
