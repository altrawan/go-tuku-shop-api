package service

import (
	"log"

	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/repository"

	"github.com/mashingan/smapping"
)

type ProfileService interface {
	List() []entity.Profile
	FindByID(profileID uint64) entity.Profile
	Update(p dto.ProfileUpdateDTO) interface{}
	ChangePassword(p dto.ProfileChangePasswordDTO) interface{}
	IsAllowedToEdit(userID uint64, profileID uint64) bool
}

type iProfileService struct {
	repository repository.ProfileRepository
}

func NewProfileService(r repository.ProfileRepository) ProfileService {
	return &iProfileService{r}
}

func (s *iProfileService) List() []entity.Profile {
	return s.repository.List()
}

func (s *iProfileService) FindByID(profileID uint64) entity.Profile {
	return s.repository.FindByID(profileID)
}

func (s *iProfileService) Update(p dto.ProfileUpdateDTO) interface{} {
	entityUser := entity.User{}
	errUser := smapping.FillStruct(&entityUser, smapping.MapFields(&p))
	if errUser != nil {
		log.Fatalf("Failed map %v", errUser)
	}

	entityProfile := entity.Profile{}
	err := smapping.FillStruct(&entityProfile, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	user, profile := s.repository.Update(entityUser, entityProfile)

	res := map[string]string{
		"name":  profile.Name,
		"email": user.Email,
	}

	return res
}

func (s *iProfileService) ChangePassword(p dto.ProfileChangePasswordDTO) interface{} {
	entityUser := entity.User{}
	entityUser.Password = p.NewPassword
	errUser := smapping.FillStruct(&entityUser, smapping.MapFields(&p))
	if errUser != nil {
		log.Fatalf("Failed map %v", errUser)
	}
	res := s.repository.ChangePassword(entityUser)
	return res
}

func (s *iProfileService) IsAllowedToEdit(userID uint64, profileID uint64) bool {
	p := s.repository.FindByID(profileID)
	return userID == p.UserID
}
