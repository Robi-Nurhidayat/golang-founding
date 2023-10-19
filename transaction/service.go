package transaction

import (
	"bwa/golang/campaign"
	"errors"
)

type TransactionService interface {
	GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error)
}

type service struct {
	repository         TransactionRepository
	campaignRepository campaign.RepositoryCampaign
}

func NewTransactionService(repository TransactionRepository, campaignRepository campaign.RepositoryCampaign) TransactionService {
	return &service{repository: repository, campaignRepository: campaignRepository}
}

func (s *service) GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error) {

	campaignById, err := s.campaignRepository.FindById(input.Id)

	if err != nil {
		return []Transaction{}, err
	}

	if campaignById.UserId != input.User.Id {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}
	transactions, err := s.repository.GetByCampaigndId(input.Id)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
