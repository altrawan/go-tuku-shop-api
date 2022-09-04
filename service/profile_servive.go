package service

import (
	"log"

	"github.com/mashingan/smapping"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/dto"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/entity"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/repository"
)

type ProfileService interface {
	Update(p dto.ProfileUpdateDTO) interface{}
	FindByID(profileID uint64) entity.Profile
	IsAllowedToEdit(userID uint64, profileID uint64) bool
}

type iProfileService struct {
	repository repository.ProfileRepository
}

func NewProfileService(r repository.ProfileRepository) ProfileService {
	return &iProfileService{r}
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

func (s *iProfileService) FindByID(profileID uint64) entity.Profile {
	return s.repository.FindByID(profileID)
}

func (s *iProfileService) IsAllowedToEdit(userID uint64, profileID uint64) bool {
	p := s.repository.FindByID(profileID)
	return userID == p.UserID
}
