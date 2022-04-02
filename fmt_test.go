package gobencherrors

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSingle_fmt_Errorf_float64_Message(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := fmt.Errorf("some error with %f float", (float64(n%17)+1)*1001)
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkSingle_fmt_Errorf_float64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := fmt.Errorf("some error with %f float", (float64(n%17)+1)*1001)
		if err == nil {
			b.FailNow()
		}
	}
}

func BenchmarkSingle_fmt_Errorf_int_Message(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := fmt.Errorf("some error with %d int", (n%17+1)*1001)
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkSingle_fmt_Errorf_int(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := fmt.Errorf("some error with %d int", (n%17+1)*1001)
		if err == nil {
			b.FailNow()
		}
	}
}

func BenchmarkSingle_errors_New_int_Message(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := errors.New("some error with " + strconv.Itoa((n%17+1)*1001) + " int")
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkSingle_errors_New_int(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := errors.New("some error with " + strconv.Itoa((n%17+1)*1001) + " int")
		if err == nil {
			b.FailNow()
		}
	}
}

func chainWrapFmtErrFloat64(v int, n int) error {
	var err error
	for i := 0; i < n; i++ {
		err = fmt.Errorf("some message %f: %w", (float64(n%17)+1)*1001, err)
	}
	return err
}

func BenchmarkWrap_fmt_Errorf_float64_Message_2wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapFmtErrFloat64(n, 2)
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_fmt_Errorf_float64_2wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapFmtErrFloat64(n, 2)
		if err == nil {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_fmt_Errorf_float64_Message_5wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapFmtErrFloat64(n, 5)
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_fmt_Errorf_float64_5wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapFmtErrFloat64(n, 5)
		if err == nil {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_fmt_Errorf_float64_Message_10wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapFmtErrFloat64(n, 10)
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_fmt_Errorf_float64_10wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapFmtErrFloat64(n, 10)
		if err == nil {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_fmt_Errorf_float64_Message_50wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapFmtErrFloat64(n, 50)
		if err.Error() == "" {
			b.FailNow()
		}
	}
}

func BenchmarkWrap_fmt_Errorf_float64_50wrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := chainWrapFmtErrFloat64(n, 50)
		if err == nil {
			b.FailNow()
		}
	}
}
