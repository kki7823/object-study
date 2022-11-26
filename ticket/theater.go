package ticket

type Theater struct {
	ticketSeller TicketSeller
}

func (t *Theater) Enter(a Audience) {
	if a.GetBag().HasInvitation() {
		ticket := t.ticketSeller.GetTicketOffice().GetTicket()
		a.GetBag().SetTicket(ticket)
	} else {
		ticket := t.ticketSeller.GetTicketOffice().GetTicket()
		a.GetBag().MinusAmount(ticket.GetFee())
		t.ticketSeller.GetTicketOffice().PlusAmount(ticket.GetFee())
		a.GetBag().SetTicket(ticket)
	}
}
