package service

import "github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/repository"

type Service struct{}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
