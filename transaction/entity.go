package transaction

import (
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
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
