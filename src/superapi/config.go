package superapi

import "github.com/rafaelcn/jeolgyu"

var (
	logger *jeolgyu.Jeolgyu
)

// SetToken sets the default token used on all web requests to the API
func SetToken(t string) {
	if len(token) == 0 {
		token = t
	}
}

// SetLogger sets the logger to be used within this module
func SetLogger(l *jeolgyu.Jeolgyu) {
	logger = l
}
