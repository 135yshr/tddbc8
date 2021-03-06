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
		{
			name: "虚部に-3を渡したときに虚部-3を持つ純虚数が生成できること",
			args: args{
				i: -3,
			},
			want: &PurelyImaginaryNumber{-3},
		},
		{
			name: "虚部に-1を渡したときに虚部-1を持つ純虚数が生成できること",
			args: args{
				i: -1,
			},
			want: &PurelyImaginaryNumber{-1},
		},
		{
			name: "虚部に1を渡したときに虚部1を持つ純虚数が生成できること",
			args: args{
				i: 1,
			},
			want: &PurelyImaginaryNumber{1},
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
		{
			name: "虚部に１を持つ純虚数の文字列表記がiになること",
			args: args{1},
			want: "i",
		},
		{
			name: "虚部に-１を持つ純虚数の文字列表記が-iになること",
			args: args{-1},
			want: "-i",
		},
		{
			name: "虚部に-3を持つ純虚数の文字列表記が-3iになること",
			args: args{-3},
			want: "-3i",
		},
		{
			name: "虚部に2を持つ純虚数の文字列表記が2iになること",
			args: args{2},
			want: "2i",
		},
		{
			name: "虚部に-2を持つ純虚数の文字列表記が-2iになること",
			args: args{-2},
			want: "-2i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut, _ := NewPurelyImaginaryNumber(tt.args.i)
			if sut.Notation() != tt.want {
				t.Errorf("NewPurelyImaginaryNumber() = %s, want = %s", tt.want, sut.Notation())
			}
		})
	}
}
