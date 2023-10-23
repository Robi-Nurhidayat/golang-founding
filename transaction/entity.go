package transaction

import (
	"bwa/golang/campaign"
	"bwa/golang/user"
	"time"
)

type Transaction struct {
	Id         int
	CampaignID int
	UserId     int
	Amount     int
	Status     string
	Code       string
	Campaign   campaign.Campaign
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PaymentURL string
}
