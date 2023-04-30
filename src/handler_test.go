package src

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestHandler_GetReport_ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockReporter(ctrl)
	m.
		EXPECT().
		GetCovidReport().
		Return([]*Person{}, nil)
	handler := NewHandler(m)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetReport(c)

	want := 200
	got := w.Code
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestHandler_GetReport_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockReporter(ctrl)
	m.
		EXPECT().
		GetCovidReport().
		Return(nil, errors.New("not found"))
	handler := NewHandler(m)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetReport(c)

	want := 500
	got := w.Code
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestHandler_GetSummary_ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockReporter(ctrl)
	m.
		EXPECT().
		GetCovidReport().
		Return([]*Person{}, nil)
	handler := NewHandler(m)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetSummary(c)
	want := 200
	got := w.Code
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestHandler_GetSummary_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockReporter(ctrl)
	m.
		EXPECT().
		GetCovidReport().
		Return(nil, errors.New("not found"))
	handler := NewHandler(m)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetSummary(c)

	want := 500
	got := w.Code
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
