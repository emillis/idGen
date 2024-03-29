package idGen

//===========[STATIC]====================================================================================================

//defaultRequirements defines Requirements that will be used by default
var defaultRequirements = Requirements{
	Length:      64,
	Salt:        "[156vs/d1ce5_35c=t+RF&^£$fDFS5RV45_;_31dfv1r4w5(}]})e6f1d1\\5sdD@5we(fE\",fe3s5EF]",
	AllowedCase: func(s string) string { return s },
	Composition: func(s string) string { return s },
	Encoder:     newSha512Encoder(),
}

//===========[STRUCTS]====================================================================================================

//Requirements defines ID configuration
type Requirements struct {
	//You can define the desired ID length here
	Length int

	//This is used for Static ID generation when the ID is used as a password.
	//You can read more about salting here: https://en.wikipedia.org/wiki/Salt_(cryptography)
	Salt string

	//If marked as true, the salt is not used by default
	OmitSalt bool

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

//adjustLength adjusts string length based on Requirements
func (r *Requirements) adjustLength(s string) string {
	if len(s) < r.Length {
		return r.adjustLength(s + r.Composition(r.Encode(s)))
	}

	return s[:r.Length]
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

	if r.Salt == "" && !r.OmitSalt {
		r.Salt = defaultRequirements.Salt
	}

	if r.AllowedCase == nil {
		r.AllowedCase = defaultRequirements.AllowedCase
	}

	if r.Composition == nil {
		r.Composition = defaultRequirements.Composition
	}

	if r.Encoder == nil {
		r.Encoder = defaultRequirements.Encoder
	}

	return r
}

//applyRequirements applies Requirements supplied to the string supplied
func applyRequirements(s string, r *Requirements) string {
	return r.AllowedCase(r.adjustLength(r.Composition(s)))
}
