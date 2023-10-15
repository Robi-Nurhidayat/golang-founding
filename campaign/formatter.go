package campaign

import "strings"

type CompaignFormatter struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserId           int    `json:"user_id"`
	Slug             string `json:"slug"`
}

func FormatterCampaign(campaign Campaign) CompaignFormatter {

	formatter := CompaignFormatter{
		Id:               campaign.Id,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserId:           campaign.UserId,
		Slug:             campaign.Slug,
	}

	formatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return formatter
}

type CampaignDetailFormatter struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Description      string   `json:"description"`
	ImageUrl         string   `json:"image_url"`
	GoalAmount       int      `json:"goal_amount"`
	CurrentAmount    int      `json:"current_amount"`
	UserId           int      `json:"user_id"`
	Slug             string   `json:"slug"`
	Perks            []string `json:"perks"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {

	campaignDetail := CampaignDetailFormatter{
		Id:               campaign.Id,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserId:           campaign.UserId,
		Slug:             campaign.Slug,
	}

	campaignDetail.ImageUrl = ""
	if len(campaign.CampaignImages) > 0 {
		campaignDetail.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetail.Perks = perks
	return campaignDetail
}
