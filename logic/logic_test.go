package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Nil(t *testing.T) {
	type args struct {
		ptr int
	}
	tests := map[string]struct {
		args args
		want int
	}{
		"test": {
			args: args{
				ptr: 10,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := failOnNil(&tt.args.ptr)
			assert.Equal(t, tt.want, got)
		})
	}
}
