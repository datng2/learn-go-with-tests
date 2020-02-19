package race

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	slowServer := createMockServerWithDelay(10 * time.Millisecond)
	fastServer := createMockServerWithDelay(3 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	t.Run("Race fast and slow", func(t *testing.T) {
		got, _ := Racer(slowServer.URL, fastServer.URL)
		if got != fastServer.URL {
			t.Errorf("got %q, want %q", got, fastServer.URL)
		}
	})

	t.Run("Race Timeout", func(t *testing.T) {
		mockTimeout := 0 * time.Millisecond
		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, mockTimeout)
		if err != ErrRaceTimeOut {
			t.Errorf("got %q, want %q", err, ErrRaceTimeOut)
		}
	})

}

func createMockServerWithDelay(delay time.Duration) *httptest.Server {
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(delay))
		w.WriteHeader(http.StatusOK)
	}
	return httptest.NewServer(http.HandlerFunc(mockHandler))
}
