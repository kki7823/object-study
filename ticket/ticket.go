package ticket

import "time"

type Invitaion struct {
	when time.Time
}

type Ticket struct {
	fee int64
}

func (t *Ticket) GetFee() int64 {
	return t.fee
}
