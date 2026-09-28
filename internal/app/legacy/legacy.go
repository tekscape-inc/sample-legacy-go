// Package legacy holds the untested business logic of the legacy sample.
package legacy

import "fmt"

// Banner returns the greeting line printed by the legacy binary.
func Banner(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("legacy says hello, %s", name)
}
