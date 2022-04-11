package idGen

import (
	"crypto/sha512"
	"encoding/base64"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

//===========[STATIC]============================================================================================

//If no configuration is passed to ID generating functions, this is used as default Config
var defaultCfg = Config{
	Length: 64,
	Salt:   "[156vs/d1ce5_35c=t+RF&^£$fDFS5RV45_;_31dfv1r4w5(}]})e6f1d1\\5sdD@5we(fE\",fe3s5EF]",
}

//Every time Random() function is invoked, this counter is incremented
var randCounter int64

//Regular expression for stripping all non-alphanumeric chars
var alphanumericOnly *regexp.Regexp
var alphaOnly *regexp.Regexp
var numericOnly *regexp.Regexp

//AllowedCase is used to define whether the ID should have only upper case or lower case characters
type AllowedCase string

const (
	//UpperOnly will convert the whole ID to upper case characters
	UpperOnly AllowedCase = "upper_only"

	//LowerOnly will convert the whole ID to lower case characters
	LowerOnly AllowedCase = "lower_only"
)

//Composition defines what the ID should be composed of.
type Composition string

const (
	//AlphaOnly will strip everything but alpha characters from the ID
	AlphaOnly Composition = "char_only"

	//NumericOnly will strip everything but numeric characters from the ID
	NumericOnly Composition = "digit_only"

	//AlphanumericOnly will strip everything but alphanumeric characters from the ID
	AlphanumericOnly Composition = "alphanumeric_only"
)

//===========[STRUCTS]============================================================================================

//Config defines ID configuration
type Config struct {
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
}

func (c *Config) toLowerCase(val *string) {
	*val = strings.ToLower(*val)
}
func (c *Config) toUpperCase(val *string) {
	*val = strings.ToUpper(*val)
}
func (c *Config) adjustLength(val *string) {
	if len(*val) > c.Length {
		*val = (*val)[:c.Length]
	} else if len(*val) < c.Length {
		ns := encode(*val)
		c.adjustComposition(&ns)
		c.adjustCase(&ns)
		*val = *val + ns
		c.adjustLength(val)
	}
}
func (c *Config) adjustCase(val *string) {
	if c.Composition == NumericOnly {
		return
	}

	if c.AllowedCase == UpperOnly {
		c.toUpperCase(val)
	} else if c.AllowedCase == LowerOnly {
		c.toLowerCase(val)
	}
}
func (c *Config) adjustComposition(val *string) {
	if c.Composition == AlphanumericOnly {
		*val = alphanumericOnly.ReplaceAllString(*val, "")
	} else if c.Composition == AlphaOnly {
		*val = alphaOnly.ReplaceAllString(*val, "")
	} else if c.Composition == NumericOnly {
		*val = numericOnly.ReplaceAllString(*val, "")
	}
}

//Applies this Config to the string supplied
func (c *Config) apply(val *string) {
	c.adjustComposition(val)

	c.adjustLength(val)

	c.adjustCase(val)
}

//===========[FUNCTIONALITY]============================================================================================

//encode encodes the supplied value
func encode(val string) string {
	sha := sha512.New()
	sha.Write([]byte(val))
	return base64.URLEncoding.EncodeToString(sha.Sum(nil))
}

//Static always generates same ID from value supplied
func Static(val string, cfg *Config) string {
	if cfg == nil {
		cfg = &defaultCfg
	}

	id := encode(cfg.Salt + val + cfg.Salt)
	cfg.apply(&id)

	return id
}

//Random generates random ID every time it's invoked
func Random(cfg *Config) string {
	if cfg == nil {
		cfg = &defaultCfg
	}

	id := encode(strconv.Itoa(int(atomic.AddInt64(&randCounter, 1) + time.Now().UnixNano())))
	cfg.apply(&id)

	return id
}

//===========[INITIALIZATION]============================================================================================

func init() {
	alphanumericOnly, _ = regexp.Compile("[^a-zA-Z0-9]+")
	alphaOnly, _ = regexp.Compile("[^a-zA-Z]+")
	numericOnly, _ = regexp.Compile("[^0-9]+")
}
