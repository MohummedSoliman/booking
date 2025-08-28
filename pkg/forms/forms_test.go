package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/anyUrl", nil)

	form := New(r.PostForm)

	valid := form.Valid()
	if !valid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form show valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("form show not valid when it should be valid")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("a")
	if has {
		t.Error("form shows has when it does not exist")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.MinLength("x", 10)

	if form.Valid() {
		t.Error("from shows min length for non-existent field")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)
	form.MinLength("a", 1)

	if !form.Valid() {
		t.Error("the form field min length should be true")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.IsEmail("email")

	if form.Valid() {
		t.Error("Should be not valid form email")
	}

	postedData = url.Values{}
	postedData.Add("email", "mo@gmail.com")
	form = New(postedData)
	form.IsEmail("email")

	if !form.Valid() {
		t.Error("should be valid email")
	}
}
