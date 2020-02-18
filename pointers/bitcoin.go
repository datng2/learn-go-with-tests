package wallet

import "fmt"

// Bitcoin is a Cryptocurrency
type Bitcoin int

// Stringer is a formatter for displaying bitcoin
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
