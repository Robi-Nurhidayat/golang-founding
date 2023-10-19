package transaction

import "time"

type TransactionCampaignFormatter struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func TransactionFormatter(transaction Transaction) TransactionCampaignFormatter {

	formatter := TransactionCampaignFormatter{}
	formatter.Id = transaction.Id
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func TransactionsFormatter(transactions []Transaction) []TransactionCampaignFormatter {

	if len(transactions) == 0 {
		return []TransactionCampaignFormatter{}
	}

	var transactionsFormatter []TransactionCampaignFormatter

	for _, transaction := range transactions {
		data := TransactionFormatter(transaction)
		transactionsFormatter = append(transactionsFormatter, data)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	Id        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  campaignFormatter `json:"campaign"`
}

type campaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {

	formatter := UserTransactionFormatter{}
	formatter.Id = transaction.Id
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaign := campaignFormatter{}
	campaign.Name = transaction.Campaign.Name
	campaign.ImageUrl = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaign.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaign

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {

	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		data := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, data)
	}

	return transactionsFormatter
}
