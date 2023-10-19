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
