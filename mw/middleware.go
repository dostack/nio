package mw

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/dostack/nio"
)

type (
	// Skipper defines a function to skip middleware. Returning true skips processing
	// the middleware.
	Skipper func(nio.Context) bool

	// BeforeFunc defines a function which is executed just before the middleware.
	BeforeFunc func(nio.Context)
)

func captureTokens(pattern *regexp.Regexp, input string) *strings.Replacer {
	groups := pattern.FindAllStringSubmatch(input, -1)
	if groups == nil {
		return nil
	}
	values := groups[0][1:]
	replace := make([]string, 2*len(values))
	for i, v := range values {
		j := 2 * i
		replace[j] = "$" + strconv.Itoa(i+1)
		replace[j+1] = v
	}
	return strings.NewReplacer(replace...)
}

// DefaultSkipper returns false which processes the middleware.
func DefaultSkipper(nio.Context) bool {
	return false
}