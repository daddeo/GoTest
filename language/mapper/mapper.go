package mapper

// need to run:
// go get github.com/rylans/getlang
// in order to pull down the library
import (
	"github.com/rylans/getlang"
)

func Greets(s string) string {
	info := getlang.FromString(s)
	switch info.LanguageCode() {
	case "en":
		return "Hello"
	case "de":
		return "Guten Tag"
	case "fr":
		return "Bonjour"
	default:
		return "I don't know your language yet."
	}
}
