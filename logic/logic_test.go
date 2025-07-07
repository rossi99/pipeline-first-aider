package logic

import (
	"testing"
	"time"

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

func Test_timeout(t *testing.T) {
	tests := map[string]struct {
		wait int
	}{
		"times out": {
			wait: 10,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			time.Sleep(time.Duration(tt.wait * int(time.Second)))
		})
	}
}
