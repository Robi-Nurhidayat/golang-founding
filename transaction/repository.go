package transaction

import "gorm.io/gorm"

type TransactionRepository interface {
	GetByCampaigndId(campaignId int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &repository{db: db}
}

func (r *repository) GetByCampaigndId(campaignId int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignId).Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
