package handler

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"
)

const appleOAuthTicketTTL = 10 * time.Minute

type appleOAuthTicketPayload struct {
	Code    string
	IDToken string
	State   string
	Expires time.Time
}

var (
	appleOAuthTicketsMu sync.Mutex
	appleOAuthTickets   = map[string]*appleOAuthTicketPayload{}
)

func createAppleOAuthTicket(code, idToken, state string) string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	ticket := hex.EncodeToString(b)

	appleOAuthTicketsMu.Lock()
	appleOAuthTickets[ticket] = &appleOAuthTicketPayload{
		Code:    code,
		IDToken: idToken,
		State:   state,
		Expires: time.Now().Add(appleOAuthTicketTTL),
	}
	pruneExpiredAppleOAuthTicketsLocked()
	appleOAuthTicketsMu.Unlock()
	return ticket
}

func consumeAppleOAuthTicket(ticket string) (*appleOAuthTicketPayload, bool) {
	appleOAuthTicketsMu.Lock()
	defer appleOAuthTicketsMu.Unlock()
	p, ok := appleOAuthTickets[ticket]
	if !ok {
		return nil, false
	}
	delete(appleOAuthTickets, ticket)
	if time.Now().After(p.Expires) {
		return nil, false
	}
	return p, true
}

func pruneExpiredAppleOAuthTicketsLocked() {
	if len(appleOAuthTickets) < 512 {
		return
	}
	now := time.Now()
	for k, v := range appleOAuthTickets {
		if now.After(v.Expires) {
			delete(appleOAuthTickets, k)
		}
	}
}
