package ticket

type Audience struct {
	bag *Bag
}

func (a *Audience) Buy(ticket *Ticket) int64 {
	return a.bag.Hold(ticket)
}
