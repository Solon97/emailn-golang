package campaign

type Repository interface {
	Create(campaign *Campaign) error
	GetAll() ([]Campaign, error)
}
