package main

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

//Default configuration
var defaultCfg = Config{
	Length: 64,
}

//This is used by Random function to create a random output every time it's called
var randCounter int64

//Regular expression for stripping all non-alphanumeric chars
var reg *regexp.Regexp

//===========[STRUCTS]============================================================================================

type Config struct {
	Length            int
	Salt              string
	UpperCaseOnly     bool
	LowerCaseOnly     bool
	AllowSpecialChars bool
}

func (c *Config) toLowerCase(val *string) {
	if !c.LowerCaseOnly {
		return
	}
	*val = strings.ToLower(*val)
}
func (c *Config) toUpperCase(val *string) {
	if !c.UpperCaseOnly {
		return
	}
	*val = strings.ToUpper(*val)
}

//stripSpecialChars strips all non-alphanumeric characters
func (c *Config) stripSpecialChars(val *string) {
	if c.AllowSpecialChars {
		return
	}
	*val = reg.ReplaceAllString(*val, "")
}
func (c *Config) adjustSize(val *string) {
	if len(*val) > c.Length {
		*val = (*val)[:c.Length]
	} else if len(*val) < c.Length {
		ns := encode(*val)
		c.process(&ns)
		*val = *val + ns
		c.adjustSize(val)
	}
}
func (c *Config) process(val *string) {
	c.toLowerCase(val)
	c.toUpperCase(val)
}

//Applies this Config to the string supplied
func (c *Config) apply(val *string) {
	c.stripSpecialChars(val)

	c.adjustSize(val)

	c.process(val)
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

func init() {
	reg, _ = regexp.Compile("[^a-zA-Z0-9]+")
}
