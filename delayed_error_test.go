package gobencherrors

import (
	"testing"
)

func chainWrapDelayedErrorFloat64(v int, n int) error {
	var err error
	for i := 0; i < n; i++ {
		err = DelayedPrefixError{Prefix: "some message", V: (float64(n%17) + 1) * 1001, Err: err}
	}
	return err
}

func BenchmarkWrap_DelayedError_float64_Message_2wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapDelayedErrorFloat64(n, 2)
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_DelayedError_float64_2wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapDelayedErrorFloat64(n, 2)
		if err == nil {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_DelayedError_float64_Message_5wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapDelayedErrorFloat64(n, 5)
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_DelayedError_float64_5wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapDelayedErrorFloat64(n, 5)
		if err == nil {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_DelayedError_float64_Message_10wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapDelayedErrorFloat64(n, 10)
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_DelayedError_float64_10wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapDelayedErrorFloat64(n, 10)
		if err == nil {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_DelayedError_float64_Message_50wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapDelayedErrorFloat64(n, 50)
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_DelayedError_float64_50wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapDelayedErrorFloat64(n, 50)
		if err == nil {
			b.FailNow()
		}
	}
}
