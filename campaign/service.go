package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type CampaignService interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaign(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type campaignService struct {
	repository RepositoryCampaign
}

func NewCampaignService(repository RepositoryCampaign) CampaignService {
	return &campaignService{
		repository: repository,
	}
}

func (c *campaignService) GetCampaigns(userId int) ([]Campaign, error) {

	if userId != 0 {
		campaigns, err := c.repository.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := c.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (c *campaignService) GetCampaign(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := c.repository.FindById(input.Id)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (c *campaignService) CreateCampaign(input CreateCampaignInput) (Campaign, error) {

	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserId = input.User.Id

	//pembuatan slug gunakan library

	stringSlug := fmt.Sprintf("%s %d", input.Name, input.User.Id)
	campaign.Slug = slug.Make(stringSlug)
	newCampaign, err := c.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
