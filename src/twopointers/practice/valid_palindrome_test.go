package practice

import "testing"

func Test_isAlphanumericCharacter(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{c: '5'}, want: true},
		{name: "test2", args: args{c: 'b'}, want: true},
		{name: "test3", args: args{c: 'B'}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlphanumericCharacter(tt.args.c); got != tt.want {
				t.Errorf("isAlphanumericCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{s: "A man, a plan, a canal: Panama"}, want: true},
		{name: "test2", args: args{s: " "}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
