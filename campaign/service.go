package campaign

type CampaignService interface {
	GetCampaigns(userId int) ([]Campaign, error)
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
