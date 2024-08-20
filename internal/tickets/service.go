package tickets

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketRepository interface {
	GetUsersTickets(primitive.ObjectID) ([]Ticket, error)
	GetTicket(primitive.ObjectID) (Ticket, error)
	CreateTicket(NewTicket) (primitive.ObjectID, error)
	UseTicket(primitive.ObjectID) error
}

type TicketService struct {
	repo TicketRepository
}

func NewTicketService(r TicketRepository) *TicketService {
	return &TicketService{repo: r}
}

func (s *TicketService) GetUsersTickets(userID primitive.ObjectID) ([]Ticket, error) {
	return s.repo.GetUsersTickets(userID)
}

func (s *TicketService) CreateTicket(newTicket NewTicket) (primitive.ObjectID, error) {
	return s.repo.CreateTicket(newTicket)
}

func (s *TicketService) UseTicket(ticketID primitive.ObjectID) TicketStatus {
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
