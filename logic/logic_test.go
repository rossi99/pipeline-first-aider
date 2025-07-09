package logic

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringConcat(t *testing.T) {
	type args struct {
		str1, str2 string
	}
	tests := map[string]struct {
		args    args
		want    string
		wantErr error
	}{
		"success: missing string 1": {
			args: args{
				str1: "",
				str2: "World",
			},
			// fixed error
			wantErr: errors.New("error: str1 cannot be empty"),
		},
		"success: missing string 2": {
			args: args{
				str1: "Hello",
				str2: "",
			},
			want: "Hello ",
		},
		"success: concatenates strings": {
			args: args{
				str1: "Hello",
				str2: "World!",
			},
			want: "Hello World!",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := stringConcat(tt.args.str1, tt.args.str2)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
