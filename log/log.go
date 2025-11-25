package log

import (
	"fmt"
	"os"
)

const (
	_VERSION = "0.3.0807"
)

func Version() string {
	return _VERSION
}

var (
	Cprintf func(uint8, string, ...interface{}) (int, error) = func(color uint8, format string, v ...interface{}) (int, error) {
		return fmt.Printf(format, v...)
	}
	Cprintln func(uint8, ...interface{}) (int, error) = func(color uint8, v ...interface{}) (int, error) {
		return fmt.Println(v...)
	}
)

var (
	Debug bool = true
)

func Warnf(format string, values ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}
	bytesWritten, err = Cprintf(YELLOW, format, values...)
	return bytesWritten, err
}

func Yellowf(format string, values ...interface{}) (bytesWritten int, err error) {
	return Warnf(format, values...)
}

func Dangerf(format string, values ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}
	bytesWritten, err = Cprintf(RED, format, values...)
	return bytesWritten, err
}

func Redf(format string, values ...interface{}) (bytesWritten int, err error) {
	return Dangerf(format, values...)
}

func Finef(format string, values ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}
	bytesWritten, err = Cprintf(GREEN, format, values...)
	return bytesWritten, err
}

func Greenf(format string, values ...interface{}) (bytesWritten int, err error) {
	return Finef(format, values...)
}

func Bluef(format string, values ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}
	bytesWritten, err = Cprintf(BLUE, format, values...)
	return bytesWritten, err
}

func Infof(format string, values ...interface{}) (bytesWritten int, err error) {
	return Bluef(format, values...)
}

func Pinkf(format string, values ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}
	bytesWritten, err = Cprintf(PINK, format, values...)
	return bytesWritten, err
}

// line

func Warnln(args ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}
	bytesWritten, err = Cprintln(YELLOW, args...)
	return bytesWritten, err
}

func Yellowln(args ...interface{}) (bytesWritten int, err error) {
	return Warnln(args...)
}

func Dangerln(args ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}

	bytesWritten, err = Cprintln(RED, args...)
	return bytesWritten, err
}

func Redln(args ...interface{}) (bytesWritten int, err error) {
	return Dangerln(args...)
}

func Fineln(args ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}

	bytesWritten, err = Cprintln(GREEN, args...)
	return bytesWritten, err
}

func Greenln(args ...interface{}) (bytesWritten int, err error) {
	return Fineln(args...)
}

func Blueln(args ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}

	bytesWritten, err = Cprintln(BLUE, args...)
	return bytesWritten, err
}

func Infoln(args ...interface{}) (bytesWritten int, err error) {
	return Blueln(args...)
}

func Pinkln(args ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}

	bytesWritten, err = Cprintln(PINK, args...)
	return bytesWritten, err
}

// fatal

func Fatalf(format string, values ...interface{}) {
	Dangerf(format, values...)
	os.Exit(1)
}

func Fatal(values ...interface{}) {
	Dangerln(values...)
	os.Exit(1)
}

func Fatalln(values ...interface{}) {
	Dangerln(values...)
	os.Exit(1)
}

func Println(args ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}

	return fmt.Println(args...)
}

func Printf(format string, values ...interface{}) (bytesWritten int, err error) {
	if !Debug {
		return 0, nil
	}

	return fmt.Printf(format, values...)
}
