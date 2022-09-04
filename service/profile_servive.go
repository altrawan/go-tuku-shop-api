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
	FindByPK(profileID uint64) entity.Profile
	FindByUserID(profileID uint64) entity.Profile
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

func (s *iProfileService) FindByPK(profileID uint64) entity.Profile {
	return s.repository.FindByPK(profileID)
}

func (s *iProfileService) FindByUserID(userID uint64) entity.Profile {
	return s.repository.FindByUserID(userID)
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
		"name":   profile.Name,
		"email":  user.Email,
		"phone":  profile.Phone,
		"gender": profile.Gender,
	}

	return res
}

func (s *iProfileService) ChangePassword(p dto.ProfileChangePasswordDTO) interface{} {
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

	user, _ := s.repository.ChangePassword(entityUser, entityProfile)

	res := map[string]string{
		"password": user.Password,
	}

	return res
}

func (s *iProfileService) IsAllowedToEdit(userID uint64, profileID uint64) bool {
	p := s.repository.FindByUserID(profileID)
	return userID == p.UserID
}
