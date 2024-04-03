package practice

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{numbers: []int{2, 7, 11, 15}, target: 9}, want: []int{1, 2}},
		{name: "test2", args: args{numbers: []int{2, 3, 4}, target: 6}, want: []int{1, 3}},
		{name: "test3", args: args{numbers: []int{-1, 0}, target: -1}, want: []int{1, 2}},
		{name: "test4", args: args{numbers: []int{5, 25, 75}, target: 100}, want: []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
