// Package forms for handling form validation
package forms

import (
	"net/http"
	"net/url"
)

// Form create custom form struct
type Form struct {
	url.Values
	errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	return x != ""
}
