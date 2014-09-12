// +build !windows,!linux,!cgo

package colour

import "os"

/*
Examples:

    isTerminal(os.Stdin)
    isTerminal(os.Stdout)
    isTerminal(os.Stderr)
*/
func isTerminal(file *os.File) bool {
	panic("Not implemented")
}
