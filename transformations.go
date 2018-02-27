/**NOTE: look back in commit history for vestigial attempt to add a pattern of transformations
to be able to inject based on flags or args
*/
package main

import "strings"

type String string

/**
situation of returning the right side of a string formatted like "#####:B####"
will return "B#####"
*/
func trimMrnFromColon(s string) string {
	return strings.Split(string(s), ":")[1]
}
