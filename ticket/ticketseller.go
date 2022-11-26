package ticket

type TicketSeller struct {
	TicketOffice *TicketOffice
}

func (s *TicketSeller) GetTicketOffice() *TicketOffice {
	return s.TicketOffice
}
