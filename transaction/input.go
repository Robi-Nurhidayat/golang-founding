package transaction

import "bwa/golang/user"

type GetCampaignTransactionInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}
type CreateTransactionInput struct {
	Amount     int       `json:"amount"`
	CampaignId int       `json:"campaign_id"`
	User       user.User `json:"user"`
}
