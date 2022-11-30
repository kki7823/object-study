package ticket

type ticketSeller struct {
	ticketOffice TicketOffice
}

func NewTiketSeller(ticketOffice TicketOffice) *ticketSeller {
	return &ticketSeller{
		ticketOffice: ticketOffice,
	}
}

func (s *ticketSeller) SellTo(a Audience) {
	s.ticketOffice.SellTicketTo(a)
}
