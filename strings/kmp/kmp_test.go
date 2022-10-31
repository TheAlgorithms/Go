package kmp

import (
	"reflect"
	"testing"
)

func TestKmp(t *testing.T) {
	type args struct {
		word         string
		text         string
		patternTable []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				word:         "ab",
				text:         "ababacaab",
				patternTable: table("ababacaab"),
			},
			want: []int{0, 2, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kmp(tt.args.word, tt.args.text, tt.args.patternTable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable(t *testing.T) {
	type args struct {
		w string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				w: "ababacaab",
			},
			want: []int{-1, 0, 0, 1, 2, 3, 0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := table(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table() = %v, want %v", got, tt.want)
			}
		})
	}
}
