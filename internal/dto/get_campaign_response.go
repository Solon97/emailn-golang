package dto

type GetCampaignResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	SendStatus string `json:"send_status"`
}
