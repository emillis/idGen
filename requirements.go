package idGen

import (
	"strings"
)

//===========[STATIC]====================================================================================================

var defaultRequirements = Requirements{
	Length:  64,
	Salt:    "[156vs/d1ce5_35c=t+RF&^£$fDFS5RV45_;_31dfv1r4w5(}]})e6f1d1\\5sdD@5we(fE\",fe3s5EF]",
	Encoder: newSha512Encoder(),
}

//===========[STRUCTS]====================================================================================================

//Requirements defines ID configuration
type Requirements struct {
	//You can define the desired ID length here
	Length int

	//This is used for Static ID generation when the ID is used as a password.
	//You can read more about salting here: https://en.wikipedia.org/wiki/Salt_(cryptography)
	Salt string

	//ID generated can be all upper case or all lower case if needed. By default,
	//the ID returned will be a mixture of upper and lower case characters.
	AllowedCase AllowedCase

	//Here you can choose your ID to be numbers only, alpha only or alphanumeric only.
	//If left undefined, the ID returned will consist off all three plus some special characters.
	Composition Composition

	//This defines encoder used in process of turning a human-readable string into a hashed one.
	//If nothing is defines, default is used.
	Encoder
}

func (c *Requirements) adjustLength(s string) string {
	if len(s) < c.Length {
		return c.adjustLength(s + c.applyComposition(enc.Encode(s)))
	}

	return s[:c.Length]
}
func (c *Requirements) applyAllowedCase(s string) string {
	switch c.AllowedCase {
	case UpperOnly:
		return strings.ToUpper(s)
	case LowerOnly:
		return strings.ToLower(s)
	}

	return s
}
func (c *Requirements) applyComposition(s string) string {

	switch c.Composition {
	case AlphaOnly:
		return alphaOnly.ReplaceAllString(s, "")
	case AlphanumericOnly:
		return alphanumericOnly.ReplaceAllString(s, "")
	case NumericOnly:
		return numericOnly.ReplaceAllString(s, "")
	}

	return s
}

//===========[FUNCTIONALITY]====================================================================================================

//makeRequirementsReasonable checks the Requirements supplied and adds default values if the ones supplied don't make sense
func makeRequirementsReasonable(r *Requirements) *Requirements {
	if r == nil {
		tmpReq := defaultRequirements
		return &tmpReq
	}

	if r.Length < 1 {
		r.Length = defaultRequirements.Length
	}

	if r.Salt == "" {
		r.Salt = defaultRequirements.Salt
	}

	if r.Encoder == nil {
		r.Encoder = defaultRequirements.Encoder
	}

	return r
}

//applyRequirements applies Requirements supplied to the string supplied
func applyRequirements(s string, r *Requirements) string {
	return r.applyAllowedCase(r.adjustLength(r.applyComposition(s)))
}
