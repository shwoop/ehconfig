package main

import (
	"reflect"
	"testing"
)

// noError test helper function to check an error is nil.
func noError(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}

// noError test helper function to check an error is not nil.
func expectError(t *testing.T, e error) {
	if e == nil {
		t.Error(e)
	}
}

// TestCheckType__valid checks valid type entries do not raise errors.
func TestCheckType__valid(t *testing.T) {
	var o string
	var e error

	for _, v := range []string{"guest", "container", "drive", "folder"} {
		o, e = checkType(v)
		noError(t, e)
		if o != v {
			t.Error("Expected %s, got %s.", v, e)
		}
	}
}

// TestCheckType__valid checks invalid type raise errors.
func TestCheckType__invalid(t *testing.T) {
	var e error

	_, e = checkType("invalid")
	expectError(t, e)
}

// TestJsonMapping__add checks that we can add to a json mapping.
func TestJsonMapping__add(t *testing.T) {
	i := make(JsonMapping)
	updateInfo(i, "arse", "face")
	expected := make(JsonMapping)
	expected["arse"] = "face"
	if !reflect.DeepEqual(i, expected) {
		t.Error("Failed to add to json mapping")
	}
}

// TestJsonMapping__update checks that we can update a json mapping.
func TestJsonMapping__update(t *testing.T) {
	i := make(JsonMapping)
	i["arse"] = "not face"
	updateInfo(i, "arse", "face")
	expected := make(JsonMapping)
	expected["arse"] = "face"
	if !reflect.DeepEqual(i, expected) {
		t.Error("Failed to update json mapping")
	}
}

// TestJsonMapping__remove checks that we can clear a value from a json mapping
func TestJsonMapping__remove(t *testing.T) {
	i := make(JsonMapping)
	i["arse"] = "face"
	updateInfo(i, "arse", "")
	expected := make(JsonMapping)
	if !reflect.DeepEqual(i, expected) {
		t.Error("Failed to remove from json mapping")
	}
}
