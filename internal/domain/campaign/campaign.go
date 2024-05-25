package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreadetAt time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for idx, email := range emails {
		contacts[idx].Email = email
	}
	newCampaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreadetAt: time.Now(),
		Contacts:  contacts,
	}
	if err := newCampaign.validate(); err != nil {
		return nil, err
	}
	return newCampaign, nil
}

func (c *Campaign) validate() error {
	if c.Name == "" {
		return errors.New(ErrNameRequired)
	}
	if c.Content == "" {
		return errors.New(ErrContentRequired)
	}
	if len(c.Contacts) == 0 {
		return errors.New(ErrContactsRequired)
	}
	return nil
}
