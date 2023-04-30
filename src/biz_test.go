package src

import (
	"reflect"
	"testing"
)

func TestGenerateSummary_one_person(t *testing.T) {
	var summary = NewSummary()
	summary.Province["Leyndell"] = 42
	summary.AgeGroup["61+"] = 1
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
	summary.AgeGroup["61+"] = 1
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
