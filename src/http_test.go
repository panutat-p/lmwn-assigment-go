package src

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestReport_GetCovidReport_status_OK(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}))
	defer srv.Close()

	report := NewReport(srv.URL, time.Second)
	_, err := report.GetCovidReport()
	if err != nil {
		t.Error("want err to be nil")
	}
}

func TestReport_GetCovidReport_status_error(t *testing.T) {
	report := NewReport("", time.Second)
	_, err := report.GetCovidReport()
	fmt.Println("err", err)
	if err.Error() != `Get "": unsupported protocol scheme ""` {
		t.Error("want unsupported protocol scheme error")
	}
}

func TestReport_GetCovidReport_timeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusRequestTimeout)
		time.Sleep(1 * time.Second)
	}))
	defer srv.Close()
	report := NewReport(srv.URL, time.Second)
	_, err := report.GetCovidReport()
	if !os.IsTimeout(err) {
		t.Error("want timeout error")
	}
}
