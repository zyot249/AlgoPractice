package practice

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				board: [][]byte{
					{'O', 'X', 'X', 'O', 'X'},
					{'X', 'O', 'O', 'X', 'O'},
					{'X', 'O', 'X', 'O', 'X'},
					{'O', 'X', 'O', 'O', 'O'},
					{'X', 'X', 'O', 'X', 'O'},
				},
			},
		},
		{
			name: "test2",
			args: args{
				board: [][]byte{
					{'O', 'O', 'O', 'O', 'X', 'X'},
					{'O', 'O', 'O', 'O', 'O', 'O'},
					{'O', 'X', 'O', 'X', 'O', 'O'},
					{'O', 'X', 'O', 'O', 'X', 'O'},
					{'O', 'X', 'O', 'X', 'O', 'O'},
					{'O', 'X', 'O', 'O', 'O', 'O'},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
		})
	}
}
