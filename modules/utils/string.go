package utils

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strings"
	"unicode"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// Slug ...
func Slug(s string) (string, error) {
	str := []byte(strings.ToLower(s))

	// convert all spaces to dash
	regE := regexp.MustCompile("[[:space:]]")
	str = regE.ReplaceAll(str, []byte("-"))

	// remove all blanks such as tab
	regE = regexp.MustCompile("[[:blank:]]")
	str = regE.ReplaceAll(str, []byte(""))

	// remove all punctuations with the exception of dash

	regE = regexp.MustCompile("[!/:-@[-`{-~]")
	str = regE.ReplaceAll(str, []byte(""))

	regE = regexp.MustCompile("/[^\x20-\x7F]/")
	str = regE.ReplaceAll(str, []byte(""))

	regE = regexp.MustCompile("`&(amp;)?#?[a-z0-9]+;`i")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("`&([a-z])(acute|uml|circ|grave|ring|cedil|slash|tilde|caron|lig|quot|rsquo);`i")
	str = regE.ReplaceAll(str, []byte("\\1"))

	regE = regexp.MustCompile("`[^a-z0-9]`i")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("`[-]+`")
	str = regE.ReplaceAll(str, []byte("-"))

	strReplaced := strings.Replace(string(str), "&", "", -1)
	strReplaced = strings.Replace(strReplaced, `"`, "", -1)
	strReplaced = strings.Replace(strReplaced, "&", "-", -1)
	strReplaced = strings.Replace(strReplaced, "--", "-", -1)

	if strings.HasPrefix(strReplaced, "-") || strings.HasPrefix(strReplaced, "--") {
		strReplaced = strings.TrimPrefix(strReplaced, "-")
		strReplaced = strings.TrimPrefix(strReplaced, "--")
	}

	if strings.HasSuffix(strReplaced, "-") || strings.HasSuffix(strReplaced, "--") {
		strReplaced = strings.TrimSuffix(strReplaced, "-")
		strReplaced = strings.TrimSuffix(strReplaced, "--")
	}

	// normalize unicode strings and remove all diacritical/accents marks
	// see https://www.socketloop.com/tutorials/golang-normalize-unicode-strings-for-comparison-purpose

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	slug, _, err := transform.String(t, strReplaced)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(slug), nil
}
