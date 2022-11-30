package ticket

type Bag struct {
	Amount    int64
	Invitaion *Invitaion
	ticket    *Ticket
}

func (b *Bag) Hold(ticket *Ticket) (fee int64) {
	if b.hasInvitation() {
		fee = 0
	} else {
		fee = ticket.GetFee()
		b.minusAmount(ticket.GetFee())
	}
	b.setTicket(ticket)
	return fee
}

func (b *Bag) hasInvitation() bool {
	return b.Invitaion != nil
}

func (b *Bag) HasTicket() bool {
	return b.ticket != nil
}

func (b *Bag) setTicket(t *Ticket) {
	b.ticket = t
}

func (b *Bag) minusAmount(amount int64) {
	b.Amount -= amount
}

func (b *Bag) PlusAmount(amount int64) {
	b.Amount += amount
}
