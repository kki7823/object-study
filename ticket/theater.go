package ticket

type Theater struct {
	ticketSeller ticketSeller
}

func (t *Theater) Enter(a Audience) {
	t.ticketSeller.SellTo(a)
}
