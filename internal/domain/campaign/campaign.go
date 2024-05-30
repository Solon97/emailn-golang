package campaign

import (
	"emailn/internal/validator"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"required,email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	CreadetAt time.Time `validate:"required"`
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
	if err := validator.ValidateStruct(newCampaign); err != nil {
		return nil, err
	}
	return newCampaign, nil
}
