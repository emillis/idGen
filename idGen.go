package idGen

import (
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

//===========[STATIC]============================================================================================

//Every time Random() function is invoked, this counter is incremented
var randCounter int64

//Regular expression for stripping all non-alphanumeric chars
var alphanumericOnly *regexp.Regexp
var alphaOnly *regexp.Regexp
var numericOnly *regexp.Regexp

//AllowedCase is used to define whether the ID should have only upper case or lower case characters
type AllowedCase func(string) string

//UpperOnly will convert the supplied string into all upper case letters
var UpperOnly AllowedCase = func(s string) string { return strings.ToUpper(s) }

//LowerOnly will convert the supplied string into all lower case letters
var LowerOnly AllowedCase = func(s string) string { return strings.ToLower(s) }

//Composition defines what the ID should be composed of.
type Composition func(string) string

//AlphaOnly will strip everything but alpha characters from the ID
var AlphaOnly Composition = func(s string) string { return alphaOnly.ReplaceAllString(s, "") }

//NumericOnly will strip everything but numeric characters from the ID
var NumericOnly Composition = func(s string) string { return numericOnly.ReplaceAllString(s, "") }

//AlphanumericOnly will strip everything but alphanumeric characters from the ID
var AlphanumericOnly Composition = func(s string) string { return alphanumericOnly.ReplaceAllString(s, "") }

//===========[FUNCTIONALITY]============================================================================================

//Static always generates same ID from value supplied
func Static(val string, r *Requirements) string {
	return generateStaticString(val, makeRequirementsReasonable(r))
}

//Random generates random ID every time it's invoked
func Random(r *Requirements) string {
	return generateRandomString(makeRequirementsReasonable(r))
}

//generateStaticString always generates same ID from value supplied, but it does no checks for invalid values supplied
func generateStaticString(val string, r *Requirements) string {
	return applyRequirements(r.Encode(r.Salt+val+r.Salt), r)
}

//generateRandomString generates random ID every time it's invoked, but it does no checks for invalid values supplied
func generateRandomString(r *Requirements) string {
	return applyRequirements(r.Encode(strconv.Itoa(int(atomic.AddInt64(&randCounter, 1)+time.Now().UnixNano()))), r)
}

//===========[INITIALIZATION]============================================================================================

func init() {
	alphanumericOnly, _ = regexp.Compile("[^a-zA-Z0-9]+")
	alphaOnly, _ = regexp.Compile("[^a-zA-Z]+")
	numericOnly, _ = regexp.Compile("[^0-9]+")
}
