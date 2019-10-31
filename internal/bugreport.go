package wtfd

import (
	"errors"
	"fmt"
	"time"
	"strconv"
	"net/smtp"
)


var (
	/* Runtime Parameter */        /* Defaults */
	BRServiceDeskAddress         = "mail@example.com" // Server recieving service desk mails
	BRServiceDeskEnabled         = false              // Is service desk support enabled
	BRRateLimitInterval  float64 = 180                // 3 Minutes
	BRRateLimitReports           = 2                  // 2 Reports during interval before beeing rate limited
	BRSMTPPort                   = 25                 // server to server smtp port
	BRSMTPUser                   = "sender"           // user used for sending mails
	BRSMTPPassword               = "passwd"           // password for user
	BRSMTPHost                   = "example.com"      // host where the send stmp server runns at

	userAccess map[string]access = make(map[string]access)
)

type access struct {
	lastBlock  time.Time // Currently unused
	lastAccess []time.Time
}

/**
 * Check if user is rate limited
 */
func BRIsUserRateLimited(u *User) bool {
	record, ok := userAccess[u.Name]
	if !ok {
		return false
	}

	/* Ok if no critical ammount of records */
	if len(record.lastAccess) < BRRateLimitReports {
		return false
	}

	/* Check if earliest record is in interval, then block */
	if time.Since(record.lastAccess[0]).Seconds() < BRRateLimitInterval {
		return true
	}

	return false
}

/**
 * Register a user access
 */
func registerUserAccess(u *User) {
	record, ok := userAccess[u.Name]

	if !ok {
		/* New record */
		record = access{
			lastBlock:  time.Time{},
			lastAccess: []time.Time{time.Now()},
		}
	} else if len(record.lastAccess) < BRRateLimitReports {
		/* No critical ammount of records */
		record.lastAccess = append(record.lastAccess, time.Now())
	} else if !BRIsUserRateLimited(u) {
		/* Cycle access */
		record.lastAccess = record.lastAccess[1:]
		record.lastAccess = append(record.lastAccess, time.Now())
	}
	userAccess[u.Name] = record
}

/**
 * Send bugreport
 */
func BRDispatchBugreport(u *User, subject string, content string) error {
	if !BRServiceDeskEnabled {
		return errors.New("Service Desk is disabled")
	}

	recipient := BRServiceDeskAddress
	recipients := []string{recipient}
	formatContent := fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nReply-To: %s\r\n\r\n%s\r\n",
		recipient, BRSMTPUser+"@"+BRSMTPHost,  subject, u.Name, content)

	auth := smtp.PlainAuth("", BRSMTPUser+"@"+BRSMTPHost, BRSMTPPassword, BRSMTPHost)

	err := smtp.SendMail(BRSMTPHost+":"+strconv.Itoa(BRSMTPPort),
		auth, BRSMTPUser+"@"+BRSMTPHost, recipients, []byte(formatContent))
	if err == nil {
		registerUserAccess(u)
	}
	return err
}
