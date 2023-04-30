package src

import (
	"reflect"
	"testing"
)

func TestNewSetFromSlice(t *testing.T) {
	var set = NewSet()
	set.Add("apple")
	set.Add("banana")
	want := set
	got := NewSetFromSlice([]string{"apple", "banana"})
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestNewSetFromString(t *testing.T) {
	want := NewSetFromSlice([]string{"apple", "banana"})
	got := NewSetFromString("apple,banana")
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestSet_Add(t *testing.T) {
	set := NewSet()
	set.Add("apple")
	set.Add("banana")
	got := set
	want := NewSetFromString("apple,banana")
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestSet_Has(t *testing.T) {
	set := NewSetFromString("apple,banana")
	if !set.Has("apple") {
		t.Error("want", true)
	}
	if !set.Has("banana") {
		t.Error("want", true)
	}
	if set.Has("carrot") {
		t.Error("want", false)
	}
}
