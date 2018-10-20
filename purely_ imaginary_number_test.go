package tddbc8

import (
	"reflect"
	"testing"
)

func Test純虚数が生成されることを確認するテスト(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    *PurelyImaginaryNumber
		wantErr bool
	}{
		{
			name: "虚部に３を渡したときに虚部３を持つ純虚数が生成できること",
			args: args{
				i: 3,
			},
			want: &PurelyImaginaryNumber{3},
		},
		{
			name: "虚部に10を渡したときに虚部10を持つ純虚数が生成できること",
			args: args{
				i: 10,
			},
			want: &PurelyImaginaryNumber{10},
		},
		{
			name: "虚部に0渡したときにエラーを返すこと",
			args: args{
				i: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPurelyImaginaryNumber(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPurelyImaginaryNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPurelyImaginaryNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test純虚数の文字列表記を確認するテスト(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "虚部に3を持つ純虚数の文字列表記が3iになること",
			args: args{3},
			want: "3i",
		},
		{
			name: "虚部に10を持つ純虚数の文字列表記が10iになること",
			args: args{10},
			want: "10i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut, _ := NewPurelyImaginaryNumber(tt.args.i)
			if sut.Notation() != tt.want {
				t.Fatalf("文字列表記が期待値と一致していない expected=%s actual=%s", tt.want, sut.Notation())
			}
		})
	}
}
