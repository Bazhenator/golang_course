package main

import (
	"fmt"
	"log"
	"sync"
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
	mu      sync.Mutex
}

var logger = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime)

// This function emulates ticket selling
func (ts *TicketSystem) SellTicket(ticketID int, wg *sync.WaitGroup, soldOutChan chan bool, logChan chan string) {
	defer wg.Done()

	ts.mu.Lock()
	defer ts.mu.Unlock()

	if ticketID < 1 || ticketID > len(ts.tickets) {
		logChan <- fmt.Sprintf("incorrect ticket ID: %d", ticketID)
		return
	}

	ticket := ts.tickets[ticketID-1]

	if ticket.Status == "available" {
		ticket.Status = "sold"
		logChan <- fmt.Sprintf("ticket %d has been sold out.", ticket.ID)

		// Checking: all tickets have been sold out
		allSold := true
		for _, t := range ts.tickets {
			if t.Status == "available" {
				allSold = false
				break
			}
		}

		// If all tickets have been gone then send signal to BoolChan
		if allSold {
			soldOutChan <- true
		}
	} else {
		logChan <- fmt.Sprintf("ticket %d has been sold out or unavailable.", ticket.ID)
	}
}

func main() {
	// Creating Ticket System
	ticketSystem := &TicketSystem{
		tickets: make([]*ConcertTicket, 10),
	}

	// Tickets initiallizing
	for i := 1; i <= 10; i++ {
		ticketSystem.tickets[i-1] = &ConcertTicket{
			ID:     i,
			Status: "available",
		}
	}

	// Channel for recieving logs
	logChannel := make(chan string)

	// BoolChannel for stop signal
	soldOutChannel := make(chan bool)

	// Channel for waiting all goroutines
	var wg sync.WaitGroup

	// Processing some goroutines for selling tickets
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			//defer wg.Done()

			ticketSystem.SellTicket(id%10+1, &wg, soldOutChannel, logChannel)

			// Delaying for emulating some operations
			time.Sleep(time.Millisecond * 500)
		}(i)
	}

	// Expecting finishing all goroutines using logChannel
	go func() {
		wg.Wait()
		close(logChannel)
		close(soldOutChannel)
	}()

	// Reading logs and waiting for stop signal
	for {
		select {
		case logEntry, ok := <-logChannel:
			if !ok {
				logChannel = nil
			} else {
				logger.Println(logEntry)
			}
		case <-soldOutChannel:
			fmt.Println("All tickets have been sold out.")
			return
		}
	}
}
