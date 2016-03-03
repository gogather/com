package log

/*
#if defined _WIN32

#include <stdio.h>
#include <stdlib.h>
#include <windows.h>

BOOL set_console_color(WORD wAttributes)
{
    HANDLE hConsole = GetStdHandle(STD_OUTPUT_HANDLE);
    if (hConsole == INVALID_HANDLE_VALUE)
        return FALSE;

    return SetConsoleTextAttribute(hConsole, wAttributes);
}

// red
void console_color_red(){
	set_console_color(FOREGROUND_RED | FOREGROUND_INTENSITY | FOREGROUND_INTENSITY);
}

// green
void console_color_green(){
	set_console_color(FOREGROUND_INTENSITY | FOREGROUND_GREEN | FOREGROUND_INTENSITY);
}

// blue
void console_color_blue(){
	set_console_color(FOREGROUND_INTENSITY | FOREGROUND_INTENSITY | FOREGROUND_BLUE);
}

// yellow
void console_color_yellow(){
	set_console_color(FOREGROUND_RED | FOREGROUND_GREEN | FOREGROUND_INTENSITY);
}

// pink
void console_color_pink(){
	set_console_color(0x0d);
}

// reset -- is white
void console_color_reset(){
	set_console_color(FOREGROUND_RED | FOREGROUND_GREEN | FOREGROUND_BLUE);
}

#else

void console_color_red(){}
void console_color_green(){}
void console_color_blue(){}
void console_color_yellow(){}
void console_color_pink(){}
void console_color_reset(){}

#endif

*/
import "C"

import (
	"fmt"
	"os"
	"runtime"
)

// format

func Warnf(format string, v ...interface{}) (n int, err error) {
	if "windows" == runtime.GOOS {
		C.console_color_yellow()
		n, err = fmt.Printf(format, v...)
		C.console_color_reset()
	} else {
		fmt.Printf("\033[1;33m") // yellow
		n, err = fmt.Printf(format, v...)
		fmt.Printf("\033[0m")
	}

	return n, err
}

func Yellowf(format string, v ...interface{}) (n int, err error) {
	return Warnf(format, v...)
}

func Dangerf(format string, v ...interface{}) (n int, err error) {
	if "windows" == runtime.GOOS {
		C.console_color_red()
		n, err = fmt.Printf(format, v...)
		C.console_color_reset()
	} else {
		fmt.Printf("\033[0;31m") // red
		n, err = fmt.Printf(format, v...)
		fmt.Printf("\033[0m")
	}

	return n, err
}

func Redf(format string, v ...interface{}) (n int, err error) {
	return Dangerf(format, v...)
}

func Finef(format string, v ...interface{}) (n int, err error) {
	if "windows" == runtime.GOOS {
		C.console_color_green()
		n, err = fmt.Printf(format, v...)
		C.console_color_reset()
	} else {
		fmt.Printf("\033[0;32m") // green
		n, err = fmt.Printf(format, v...)
		fmt.Printf("\033[0m")
	}

	return n, err
}

func Greenf(format string, v ...interface{}) (n int, err error) {
	return Finef(format, v...)
}

func Bluef(format string, v ...interface{}) (n int, err error) {
	if "windows" == runtime.GOOS {
		C.console_color_blue()
		n, err = fmt.Printf(format, v...)
		C.console_color_reset()
	} else {
		fmt.Printf("\033[0;34m") // blue
		n, err = fmt.Printf(format, v...)
		fmt.Printf("\033[0m")
	}

	return n, err
}

func Infof(format string, v ...interface{}) (n int, err error) {
	return Bluef(format, v...)
}

func Pinkf(format string, v ...interface{}) (n int, err error) {
	if "windows" == runtime.GOOS {
		C.console_color_pink()
		n, err = fmt.Printf(format, v...)
		C.console_color_reset()
	} else {
		fmt.Printf("\033[0;35m") // blue
		n, err = fmt.Printf(format, v...)
		fmt.Printf("\033[0m")
	}

	return n, err
}

// line

func Warnln(a ...interface{}) (n int, err error) {
	if "windows" == runtime.GOOS {
		C.console_color_yellow()
		n, err = fmt.Println(a...)
		C.console_color_reset()
	} else {
		fmt.Printf("\033[1;33m") // yellow
		n, err = fmt.Println(a...)
		fmt.Printf("\033[0m")
	}

	return n, err
}

func Yellowln(a ...interface{}) (n int, err error) {
	return Warnln(a...)
}

func Dangerln(a ...interface{}) (n int, err error) {
	if "windows" == runtime.GOOS {
		C.console_color_red()
		n, err = fmt.Println(a...)
		C.console_color_reset()
	} else {
		fmt.Printf("\033[0;31m") // red
		n, err = fmt.Println(a...)
		fmt.Printf("\033[0m")
	}

	return n, err
}

func Redln(a ...interface{}) (n int, err error) {
	return Dangerln(a...)
}

func Fineln(a ...interface{}) (n int, err error) {
	if "windows" == runtime.GOOS {
		C.console_color_green()
		n, err = fmt.Println(a...)
		C.console_color_reset()
	} else {
		fmt.Printf("\033[0;32m") // green
		n, err = fmt.Println(a...)
		fmt.Printf("\033[0m")
	}

	return n, err
}

func Greenln(a ...interface{}) (n int, err error) {
	return Fineln(a...)
}

func Blueln(a ...interface{}) (n int, err error) {
	if "windows" == runtime.GOOS {
		C.console_color_blue()
		n, err = fmt.Println(a...)
		C.console_color_reset()
	} else {
		fmt.Printf("\033[0;34m") // blue
		n, err = fmt.Println(a...)
		fmt.Printf("\033[0m")
	}

	return n, err
}

func Infoln(a ...interface{}) (n int, err error) {
	return Blueln(a...)
}

func Pinkln(a ...interface{}) (n int, err error) {
	if "windows" == runtime.GOOS {
		C.console_color_pink()
		n, err = fmt.Println(a...)
		C.console_color_reset()
	} else {
		fmt.Printf("\033[0;35m") // blue
		n, err = fmt.Println(a...)
		fmt.Printf("\033[0m")
	}

	return n, err
}

// fatal

func Fatalf(format string, v ...interface{}) {
	Dangerf(format, v...)
	os.Exit(1)
}

func Fatal(v ...interface{}) {
	Dangerln(v...)
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	Dangerln(v...)
	os.Exit(1)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func Printf(format string, v ...interface{}) (n int, err error) {
	return fmt.Printf(format, v...)
}
