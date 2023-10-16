package campaign

import "gorm.io/gorm"

type RepositoryCampaign interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
	FindById(id int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(image CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(id int) (bool, error)
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

func (r *repositoryCampaign) FindById(id int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ? ", id).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repositoryCampaign) Save(campaign Campaign) (Campaign, error) {

	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repositoryCampaign) Update(campaign Campaign) (Campaign, error) {

	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repositoryCampaign) CreateImage(image CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&image).Error

	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *repositoryCampaign) MarkAllImagesAsNonPrimary(id int) (bool, error) {
	//update campaign_images set is_primary = false where campain_id = 1
	//	terjemah di gorm

	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", id).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
