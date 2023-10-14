package campaign

import "gorm.io/gorm"

type RepositoryCampaign interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
}

type repositoryCampaign struct {
	db *gorm.DB
}

func NewRepositoryCampaign(db *gorm.DB) RepositoryCampaign {
	return &repositoryCampaign{
		db: db,
	}
}

func (r *repositoryCampaign) FindAll() ([]Campaign, error) {

	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repositoryCampaign) FindByUserId(userId int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
