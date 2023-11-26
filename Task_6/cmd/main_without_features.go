package main

import (
	"fmt"
	"log"
	"time"
)

// Concert Ticket's structure
type ConcertTicket struct {
	ID     int
	Status string // "available" or "sold"
}

// Ticket System
type TicketSystem struct {
	tickets []*ConcertTicket
}

var logger = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime)

// This function emulates ticket selling
func (ts *TicketSystem) SellTicket(ticketID int) {

	if ticketID < 1 || ticketID > len(ts.tickets) {
		logger.Printf("incorrect ticket ID: %d\n", ticketID)
		return
	}

	ticket := ts.tickets[ticketID-1]

	if ticket.Status == "available" {
		ticket.Status = "sold"
		logger.Printf("ticket %d has been sold out.\n", ticket.ID)

		// Checking: all tickets have been sold out
		allSold := true
		for _, t := range ts.tickets {
			if t.Status == "available" {
				allSold = false
				break
			}
		}

		// If all tickets have been gone then print message
		if allSold {
			fmt.Println("All tickets have been sold out.")
		}
	} else {
		logger.Printf("ticket %d has been sold out or unavailable.\n", ticket.ID)
	}
}

func main() {
	// Creating Ticket System
	ticketSystem := &TicketSystem{
		tickets: make([]*ConcertTicket, 10),
	}

	// Tickets initializing
	for i := 1; i <= 10; i++ {
		ticketSystem.tickets[i-1] = &ConcertTicket{
			ID:     i,
			Status: "available",
		}
	}

	// Processing some goroutines for selling tickets
	for i := 0; i < 15; i++ {
		go func(id int) {
			// Delaying for emulating some operations
			time.Sleep(time.Millisecond * 500)
			ticketSystem.SellTicket(id%10 + 1)
		}(i)
	}

	// Adding a delay to let all goroutines finish
	time.Sleep(time.Second)
}
