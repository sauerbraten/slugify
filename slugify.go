package slugify

import (
	"regexp"
	"strings"
)

var replacements = map[rune]string{
	'á':  "a",
	'à':  "a",
	'ä':  "a",
	'â':  "a",
	'ã':  "a",
	'æ':  "ae",
	'é':  "e",
	'è':  "e",
	'ë':  "e",
	'ê':  "e",
	'ẽ':  "e",
	'í':  "i",
	'ì':  "i",
	'ï':  "i",
	'î':  "i",
	'ĩ':  "i",
	'ó':  "o",
	'ò':  "o",
	'ö':  "o",
	'ő':  "o",
	'ô':  "o",
	'ø':  "o",
	'õ':  "o",
	'ú':  "u",
	'ù':  "u",
	'ü':  "u",
	'ű':  "u",
	'û':  "u",
	'ũ':  "u",
	'ç':  "c",
	'ñ':  "n",
	'ß':  "ss",
	'\'': "", // collapses "adam's" to "adams"
}

// S creates a slug from the str input string. The slug only contains a-z, 0-9 and '-', making it safe to use in URLs.
func S(str string) string {
	// trim all whitespace and special characters at beginning and end, then convert to lower case
	str = strings.ToLower(strings.TrimSpace(str))
	newstr := ""

	// iterate over string, for each rune check if there is a replacement, if so, append it to newstr, if not, append original value to newstr
	for _, char := range str {
		if replacement, ok := replacements[char]; ok {
			newstr += replacement
		} else {
			newstr += string(char)
		}
	}

	// replace invalid chars (spaces are invalid) with dashes
	newstr = regexp.MustCompile("[^a-z0-9-]").ReplaceAllString(newstr, "-")

	// collapse dashes ('---' → '-')
	newstr = regexp.MustCompile("-+").ReplaceAllString(newstr, "-")

	// trim dashes from beginning and end; return the result
	return strings.Trim(newstr, "-")
}
