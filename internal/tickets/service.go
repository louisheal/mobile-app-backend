package tickets

import (
	"mobile-app-backend/internal/users"
)

type TicketRepository interface {
	GetUsersTickets(users.UserID) ([]Ticket, error)
	GetTicket(TicketID) (Ticket, error)
	CreateTicket(TicketInput) (TicketID, error)
	UseTicket(TicketID) error
}

type TicketService struct {
	repo TicketRepository
}

func NewTicketService(r TicketRepository) *TicketService {
	return &TicketService{repo: r}
}

func (s *TicketService) GetUsersTickets(userID users.UserID) ([]Ticket, error) {
	return s.repo.GetUsersTickets(userID)
}

func (s *TicketService) CreateTicket(newTicket TicketInput) (TicketID, error) {
	return s.repo.CreateTicket(newTicket)
}

func (s *TicketService) UseTicket(ticketID TicketID) TicketStatus {
	if ticket, err := s.repo.GetTicket(ticketID); err != nil {
		return Invalid
	} else if ticket.Used {
		return Used
	}

	if err := s.repo.UseTicket(ticketID); err != nil {
		panic(err)
	}

	return Valid
}
