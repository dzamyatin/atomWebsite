package usecase

import (
	"github.com/dzamyatin/atomWebsite/internal/dto"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/dzamyatin/atomWebsite/internal/validator"
)

type RegistrationUseCase struct {
	userRepository  repository.IUserRepository
	passwordEncoder entity.PasswordEncoder
	validator       validator.IRegistrationValidator
}

func NewRegistrationUseCase(userRepository repository.IUserRepository) *RegistrationUseCase {
	return &RegistrationUseCase{userRepository: userRepository}
}

func (r *RegistrationUseCase) Execute(request dto.RegistrationRequest) error {
	if err := r.validator.Validate(request); err != nil {
		return err
	}

	user := entity.NewUserEntity(request.Email, request.Phone)

	if request.Password != "" {
		err := user.AddPassword(request.Password, r.passwordEncoder)
		if err != nil {
			return err
		}
	}

	err := r.userRepository.AddUser(user)

	if err != nil {
		return err
	}

	return nil
}
