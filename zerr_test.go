package zerr

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestWrap(t *testing.T) {
	err := fmt.Errorf("text")

	type args struct {
		err error
		msg []string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "formal",
			args: args{
				err: err,
				msg: []string{"message"},
			},
			want: fmt.Errorf("zerr_test.go/zerr.TestWrap.func1():62: message: %w", err),
		}, {
			name: "formal",
			args: args{
				err: err,
				msg: []string{"message1", "message2"},
			},
			want: fmt.Errorf("zerr_test.go/zerr.TestWrap.func1():62: message1: message2: %w", err),
		}, {
			name: "nil",
			args: args{
				err: nil,
				msg: nil,
			},
			want: nil,
		}, {
			name: "nil",
			args: args{
				err: err,
				msg: nil,
			},
			want: fmt.Errorf("zerr_test.go/zerr.TestWrap.func1():62: %w", err),
		}, {
			name: "nil",
			args: args{
				err: nil,
				msg: []string{"message"},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.args.err, tt.args.msg...); !(errors.Is(got, tt.args.err) && reflect.DeepEqual(tt.want, got)) {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
