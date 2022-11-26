package ticket

type TicketOffice struct {
	Amount  int64
	Tickets []Ticket
}

func (o *TicketOffice) GetTicket() *Ticket {
	if len(o.Tickets) == 0 {
		return nil
	}

	t := o.Tickets[0]

	if len(o.Tickets) != 1 {
		o.Tickets = o.Tickets[1:]
	}

	return &t
}

func (o *TicketOffice) MinusAmount(amount int64) {
	o.Amount -= amount
}
func (o *TicketOffice) PlusAmount(amount int64) {
	o.Amount += amount
}
