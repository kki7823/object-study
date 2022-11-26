package ticket

type Bag struct {
	Amount    int64
	Invitaion *Invitaion
	ticket    *Ticket
}

func (b *Bag) HasInvitation() bool {
	return b.Invitaion != nil
}

func (b *Bag) HasTicket() bool {
	return b.ticket != nil
}

func (b *Bag) SetTicket(t *Ticket) {
	b.ticket = t
}

func (b *Bag) MinusAmount(amount int64) {
	b.Amount -= amount
}

func (b *Bag) PlusAmount(amount int64) {
	b.Amount += amount
}
