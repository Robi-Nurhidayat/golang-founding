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

type TransactionNotificationAmountInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderId           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
