package usecase

import (
	"errors"
	"os"

	"github.com/VinncentWong/DiverBE/app/diver/repository"
	"github.com/VinncentWong/DiverBE/domain"
	_diverDto "github.com/VinncentWong/DiverBE/domain/dto/diver"

	"github.com/VinncentWong/DiverBE/util"
)

type IDiverUsecase interface {
	CreateUsernameIndex() error
	CreateDiver() (domain.Diver, error)
	LoginDiver(dto _diverDto.LoginDto) (domain.Diver, error)
}

type DiverUsecase struct {
	r repository.IDiverRepository
}

func NewDiverUsecase(r repository.IDiverRepository) IDiverUsecase {
	return &DiverUsecase{
		r: r,
	}
}

func (u *DiverUsecase) CreateUsernameIndex() error {
	err := u.r.CreateUsernameIndex()
	return err
}

func (u *DiverUsecase) CreateDiver() (domain.Diver, error) {
	result, err := u.r.GetDiver(os.Getenv("USER_USERNAME"))
	if err != nil {
		if err.Error() != "diver doesn't exist" {
			return domain.Diver{}, err
		}
	}
	if result.Email != "" && result.Password != "" && result.Username != "" {
		return domain.Diver{}, errors.New("diver already exist")
	}
	result, err = u.r.CreateDiver()
	if err != nil {
		return domain.Diver{}, err
	}
	return result, nil
}

func (u *DiverUsecase) LoginDiver(dto _diverDto.LoginDto) (domain.Diver, error) {
	result, err := u.r.LoginDiver(dto)
	if err != nil {
		return domain.Diver{}, err
	}
	err = util.ComparePassword(dto.Password, result.Password)
	if result.Username != os.Getenv("USER_USERNAME") || err != nil {
		return domain.Diver{}, err
	}
	return result, nil
}
