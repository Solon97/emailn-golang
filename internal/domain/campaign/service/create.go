package campaign_service

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/dto"
	internalerrors "emailn/internal/errors"

	"github.com/sirupsen/logrus"
)

func (s *CampaignService) Create(newCampaign *dto.NewCampaign) (string, error) {
	campaign, err := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Contacts)
	if err != nil {
		return "", err
	}
	// TODO: Criar erros personalizados Repository Errors poderiam ser considerados InternalErros (500) e Domain Errors poderiam ser considerados 400
	//* Usar Wrap para envelopar mensagem de erros.
	if err := s.repository.Create(campaign); err != nil {
		logrus.WithError(err).Error("error saving campaign")
		return "", internalerrors.ErrInternalServer
	}
	return campaign.ID, nil
}
