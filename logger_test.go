package logger

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	type args struct {
		isFile bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "file-logger",
			args: args{
				isFile: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLogger(tt.args.isFile)
		})
	}
}
