package logic

// import (
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// )

// func Test_Nil(t *testing.T) {
// 	type args struct {
// 		ptr int
// 	}
// 	tests := map[string]struct {
// 		args args
// 		want int
// 	}{
// 		"test": {
// 			args: args{
// 				ptr: 10,
// 			},
// 		},
// 	}
// 	for name, tt := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			got := failOnNil(&tt.args.ptr)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }

// func Test_timeout(t *testing.T) {
// 	tests := map[string]struct {
// 		wait int
// 	}{
// 		"times out": {
// 			wait: 10,
// 		},
// 	}
// 	for name, tt := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			time.Sleep(time.Duration(tt.wait * int(time.Second)))
// 		})
// 	}
// }

// //Failing Tests

// func Test_Fail_Equality(t *testing.T) {
// 	assert.Equal(t, 1, 2)
// }

// func Test_Fail_StringComparison(t *testing.T) {
// 	assert.Equal(t, "hello", "world")
// }

// func Test_Fail_NilCheck(t *testing.T) {
// 	var a *int
// 	assert.NotNil(t, a)
// }

// func Test_Fail_NotEqual(t *testing.T) {
// 	assert.NotEqual(t, 3.14, 3.14)
// }

// func Test_Fail_SliceComparison(t *testing.T) {
// 	expected := []int{1, 2, 3}
// 	actual := []int{3, 2, 1}
// 	assert.Equal(t, expected, actual)
// }

// func Test_Fail_MapComparison(t *testing.T) {
// 	expected := map[string]int{"a": 1}
// 	actual := map[string]int{"a": 2}
// 	assert.Equal(t, expected, actual)
// }

// func Test_Fail_Panic(t *testing.T) {
// 	assert.Panics(t, func() {
// 		x := 1 + 1
// 		_ = x
// 	})
// }

// func Test_Fail_True(t *testing.T) {
// 	assert.True(t, false)
// }

// func Test_Fail_False(t *testing.T) {
// 	assert.False(t, true)
// }

// func Test_Fail_Error(t *testing.T) {
// 	err := nil
// 	assert.Error(t, err)
// }

// func Test_Fail_Zero(t *testing.T) {
// 	var x int = 1
// 	assert.Zero(t, x)
// }

// func Test_Fail_Contains(t *testing.T) {
// 	assert.Contains(t, "hello world", "bye")
// }

// // Passing tests
// func Test_Pass_Equality(t *testing.T) {
// 	assert.Equal(t, 42, 42)
// }

// func Test_Pass_StringComparison(t *testing.T) {
// 	assert.Equal(t, "hello", "hello")
// }

// func Test_Pass_NotEqual(t *testing.T) {
// 	assert.NotEqual(t, 100, 200)
// }

// func Test_Pass_NilCheck(t *testing.T) {
// 	var ptr *int = nil
// 	assert.Nil(t, ptr)
// }

// func Test_Pass_NotNil(t *testing.T) {
// 	x := 5
// 	assert.NotNil(t, &x)
// }

// func Test_Pass_SliceEquality(t *testing.T) {
// 	assert.Equal(t, []string{"a", "b", "c"}, []string{"a", "b", "c"})
// }

// func Test_Pass_MapEquality(t *testing.T) {
// 	expected := map[string]int{"foo": 1, "bar": 2}
// 	actual := map[string]int{"foo": 1, "bar": 2}
// 	assert.Equal(t, expected, actual)
// }

// func Test_Pass_Panic(t *testing.T) {
// 	assert.Panics(t, func() {
// 		panic("something went wrong")
// 	})
// }

// func Test_Pass_True(t *testing.T) {
// 	assert.True(t, 1 < 2)
// }

// func Test_Pass_False(t *testing.T) {
// 	assert.False(t, 10 == 5)
// }

// func Test_Pass_Zero(t *testing.T) {
// 	var count int
// 	assert.Zero(t, count)
// }

// func Test_Pass_Contains(t *testing.T) {
// 	assert.Contains(t, "golang is awesome", "awesome")
// }
