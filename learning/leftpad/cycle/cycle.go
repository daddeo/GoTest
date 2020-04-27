package cycle

import (
	"test/learning/leftpad"
)

var defaultCharacter = ' '

// FormatDouble blah blah blah
func FormatDouble(s string, i int) string {
	return leftpad.Format(s+s, i)
}
