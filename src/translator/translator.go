package translator

import "regexp"

var rx = regexp.MustCompile(`[a-zA-Z]+`)

// TranslateBytes returns copy of src, replacing alphabetical words with the given lexeme.
func TranslateBytes(src, lexeme []byte) []byte {
	return rx.ReplaceAll(src, lexeme)
}