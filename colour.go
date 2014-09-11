// Package colour provides Quake-style colour formatting [1] for Unix terminals.
//
// eg.
//
//     Printf("^0black ^1red ^2green ^3yellow ^4blue ^5magenta ^6cyan ^7white^R\n")
//
// The following sequences are converted to their equivalent ANSI colours:
//
//     ^0 = Black
//     ^1 = Red
//     ^2 = Green
//     ^3 = Yellow
//     ^4 = Blue
//     ^5 = Cyan (light blue)
//     ^6 = Magenta (purple)
//     ^7 = White
//     ^8 = Black Background
//     ^9 = Red Background
//     ^a = Green Background
//     ^b = Yellow Background
//     ^c = Blue Background
//     ^d = Cyan (light blue) Background
//     ^e = Magenta (purple) Background
//     ^f = White Background
//     ^R = Reset
//     ^U = Underline
//     ^B = Bold
//
// [1] http://www.holysh1t.net/quake-live-colors-nickname/
package colour

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
)

var (
	extract = regexp.MustCompile(`(\^[0-9a-fRUB])|(\^\^)|([^^]+)`)
	colours = map[byte]string{
		'0': "\033[30m",
		'1': "\033[31m",
		'2': "\033[32m",
		'3': "\033[33m",
		'4': "\033[34m",
		'5': "\033[35m",
		'6': "\033[36m",
		'7': "\033[37m",

		'8': "\033[40m",
		'9': "\033[41m",
		'a': "\033[42m",
		'b': "\033[43m",
		'c': "\033[44m",
		'd': "\033[45m",
		'e': "\033[46m",
		'f': "\033[47m",

		'R': "\033[0m",
		'U': "\033[4m",
		'B': "\033[1m",
	}
)

func formatString(s string) string {
	out := &bytes.Buffer{}
	for _, match := range extract.FindAllStringSubmatch(s, -1) {
		if match[1] != "" {
			n := match[1][1]
			out.WriteString(colours[n])
		} else if match[2] != "" {
			out.WriteString("^")
		} else {
			out.WriteString(match[3])
		}
	}
	return out.String()
}

func Fprintln(w io.Writer, args ...interface{}) (n int, err error) {
	for i, arg := range args {
		if s, ok := arg.(string); ok {
			args[i] = formatString(s)
		}
	}
	return fmt.Fprintln(w, args...)
}

func Println(args ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, args...)
}

func Fprint(w io.Writer, args ...interface{}) (n int, err error) {
	for i, arg := range args {
		if s, ok := arg.(string); ok {
			args[i] = formatString(s)
		}
	}
	return fmt.Fprint(w, args...)
}

func Print(args ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, args...)
}

func Fprintf(w io.Writer, format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, formatString(format), args...)
}

func Printf(format string, args ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, args...)
}

func Sprintf(format string, args ...interface{}) string {
	w := &bytes.Buffer{}
	Fprintf(w, format, args...)
	return w.String()
}
