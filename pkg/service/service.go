package service

import (
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/entity"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/repository"
)

type Service struct {
	repos *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{repos}
}

func (s *Service) GetAllBanks() ([]entity.Bank, error) {
	return s.repos.GetAllBanks()
}

func (s *Service) CreateBank(bank entity.Bank) error {
	return s.repos.CreateBank(bank)
}

func (s *Service) UpdateBank(bankId int, bank *entity.Bank) error {
	return s.repos.UpdateBank(bankId, bank)
}

func (s *Service) GetBankByID(bankId int) (*entity.Bank, error) {
	return s.repos.GetBankByID(bankId)
}

func (s *Service) DeleteBank(bankId int) (string, error) {
	return s.repos.DeleteBank(bankId)
}
