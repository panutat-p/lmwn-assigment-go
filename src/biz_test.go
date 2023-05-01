package src

import (
	"reflect"
	"testing"
)

func TestGenerateSummary_one_person(t *testing.T) {
	var summary = NewSummary()
	summary.Province["Leyndell"] = 42
	summary.AgeGroup["60+"] = 1
	want := summary

	morrgot := Person{ProvinceID: 42, ProvinceEn: "Leyndell", Age: 110}
	got := GenerateSummary([]*Person{&morrgot})
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerateSummary_empty(t *testing.T) {
	want := NewSummary()
	got := GenerateSummary([]*Person{})
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerateFilteredSummary_one_person(t *testing.T) {
	want := NewSummary()
	morrgot := Person{ProvinceID: 42, ProvinceEn: "Leyndell", Age: 110}
	got := GenerateFilteredSummary([]*Person{&morrgot}, NewSet())
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerateFilteredSummary_not_in_filter(t *testing.T) {
	var summary = NewSummary()
	summary.Province["Leyndell"] = 42
	summary.AgeGroup["60+"] = 1
	want := summary

	morrgot := Person{ProvinceID: 42, ProvinceEn: "Leyndell", Age: 110}
	got := GenerateFilteredSummary([]*Person{&morrgot}, NewSetFromString("Leyndell"))
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerateFilteredSummary_empty(t *testing.T) {
	want := NewSummary()
	got := GenerateFilteredSummary([]*Person{}, NewSet())
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestDetermineAge_child(t *testing.T) {
	want := "0-30"
	got := DetermineAge(15)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestDetermineAge_mid(t *testing.T) {
	want := "31-60"
	got := DetermineAge(45)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestDetermineAge_aged(t *testing.T) {
	want := "60+"
	got := DetermineAge(80)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestDetermineAge_null(t *testing.T) {
	want := "N/A"
	got := DetermineAge(-1)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
