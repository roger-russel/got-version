package tag

import (
	"bufio"
	"os"
	"testing"
)

func TestGetLastTag(t *testing.T) {
	type args struct {
		pipe *bufio.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple Prefixed",
			args: args{
				pipe: getSimplePrefixed(),
			},
			want: "v123.3.2",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLastTag(tt.args.pipe); got != tt.want {
				t.Errorf("GetLastTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getSimplePrefixed() *bufio.Reader {
	file, err := os.Open("../../tests/fixture/prefixed.list.txt")

	if err != nil {
		panic(err)
	}

	return bufio.NewReader(file)
}
