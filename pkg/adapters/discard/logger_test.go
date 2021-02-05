package discard

import (
	"errors"
	"testing"

	"go.opentelemetry.io/otel/label"
)

// func TestDiscard(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want Logger
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Discard(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Discard() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestLogger_Error(t *testing.T) {
// 	type args struct {
// 		in0 error
// 		in1 string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := Logger{}
// 		})
// 	}
// }
//
// func TestLogger_ErrorWithContext(t *testing.T) {
// 	type args struct {
// 		in0 error
// 		in1 string
// 		in2 logs.Context
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := Logger{}
// 		})
// 	}
// }
//
// func TestLogger_Log(t *testing.T) {
// 	type args struct {
// 		in0 string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := Logger{}
// 		})
// 	}
// }
//
// func TestLogger_LogWithContext(t *testing.T) {
// 	type args struct {
// 		in0 string
// 		in1 logs.Context
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := Logger{}
// 		})
// 	}
// }
//
// func TestLogger_Verbose(t *testing.T) {
// 	type args struct {
// 		in0 string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := Logger{}
// 		})
// 	}
// }
//
// func TestLogger_VerboseWithContext(t *testing.T) {
// 	type args struct {
// 		in0 string
// 		in1 logs.Context
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := Logger{}
// 		})
// 	}
// }
//
// func TestLogger_WithBaseContext(t *testing.T) {
// 	type args struct {
// 		in0 logs.Context
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want logs.Logger
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := Logger{}
// 			if got := l.WithBaseContext(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("WithBaseContext() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestLogger_WithConfiguration(t *testing.T) {
// 	type args struct {
// 		in0 logs.Configuration
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want logs.Logger
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := Logger{}
// 			if got := l.WithConfiguration(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("WithConfiguration() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func BenchmarkLogger_Error(b *testing.B) {
	logger := Discard()
	err := errors.New("example error")

	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		logger.ErrorWithLabels(err, "test message", label.Any("foo", "bar"))
	}
}
