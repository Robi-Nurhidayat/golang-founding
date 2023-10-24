package transaction

import (
	"bwa/golang/campaign"
	"bwa/golang/payment"
	"errors"
	"strconv"
)

type TransactionService interface {
	GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationAmountInput) error
}

type service struct {
	repository         TransactionRepository
	campaignRepository campaign.RepositoryCampaign
	paymentService     payment.PaymentService
}

func NewTransactionService(repository TransactionRepository, campaignRepository campaign.RepositoryCampaign, paymentService payment.PaymentService) TransactionService {
	return &service{repository: repository, campaignRepository: campaignRepository, paymentService: paymentService}
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

func (s *service) GetTransactionByUserId(userId int) ([]Transaction, error) {

	transactions, err := s.repository.GetByUserId(userId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignId
	transaction.Amount = input.Amount
	transaction.UserId = input.User.Id
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.TransactionPayment{
		Id:     newTransaction.Id,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetURLPayment(paymentTransaction, input.User)

	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationAmountInput) error {
	transactionId, _ := strconv.Atoi(input.OrderId)

	transactionById, err := s.repository.GetById(transactionId)

	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transactionById.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transactionById.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transactionById.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transactionById)
	if err != nil {
		return err
	}

	campaignById, err := s.campaignRepository.FindById(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaignById.BackerCount = campaignById.BackerCount + 1
		campaignById.CurrentAmount = campaignById.CurrentAmount + updatedTransaction.Amount

		_, err := s.campaignRepository.Update(campaignById)
		if err != nil {
			return err
		}
	}

	return nil
}
