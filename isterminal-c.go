// +build !windows,!linux,cgo

package colour

/*
#include <unistd.h>
*/
import "C"

import "os"

/*
Examples:

    isTerminal(os.Stdin)
    isTerminal(os.Stdout)
    isTerminal(os.Stderr)
*/
func isTerminal(file *os.File) bool {
	return int(C.isatty(C.int(file.Fd()))) != 0
}
