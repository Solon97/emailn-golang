package campaign

import (
	"emailn/internal/validator"
	"time"

	"github.com/rs/xid"
)

type SendStatus string

const (
	SendStatusPending SendStatus = "pending"
	SendStatusStarted SendStatus = "started"
	SendstatusDone    SendStatus = "done"
	SendStatusFailed  SendStatus = "failed"
)

type Contact struct {
	ID         string `validate:"required"`
	Email      string `validate:"required,email"`
	CampaignID string
}

type Campaign struct {
	ID         string     `validate:"required"`
	Name       string     `validate:"min=5,max=24"`
	Content    string     `validate:"min=5,max=1024"`
	Contacts   []Contact  `validate:"min=1,dive"`
	SendStatus SendStatus `validate:"oneof=pending started done failed"`
	CreadetAt  time.Time  `validate:"required"`
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for idx, email := range emails {
		contacts[idx].ID = xid.New().String()
		contacts[idx].Email = email
	}
	newCampaign := &Campaign{
		ID:         xid.New().String(),
		Name:       name,
		Content:    content,
		CreadetAt:  time.Now(),
		Contacts:   contacts,
		SendStatus: SendStatusPending,
	}
	if err := validator.ValidateStruct(newCampaign); err != nil {
		return nil, err
	}
	return newCampaign, nil
}
