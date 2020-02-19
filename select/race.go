package race

import (
	"net/http"
	"time"
)

const (
	ErrRaceTimeOut = RaceError("Time out")
)

type RaceError string

func (e RaceError) Error() string { return string(e) }

func Racer(x, y string) (winner string, err error) {
	return ConfigurableRacer(x, y, 10*time.Second)
	// startX := time.Now()
	// http.Get(x)
	// xDuration := time.Since(startX)

	// startY := time.Now()
	// http.Get(y)
	// yDuration := time.Since(startY)

	// if winner = x; xDuration > yDuration {
	// 	winner = y
	// }
	// return winner

}

func ConfigurableRacer(x, y string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(x):
		return x, nil
	case <-ping(y):
		return y, nil
	case <-time.After(timeout):
		return "", ErrRaceTimeOut
	}
}

// Ping acccesses a website and returns
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func(x string) {
		http.Get(x)
		close(ch)
	}(url)
	return ch
}
