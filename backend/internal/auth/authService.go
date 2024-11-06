package auth

import (
	"errors"
	"time"

	"github.com/codeableorg/weekend-challenge-13-GabNatali/internal/base/crypto"
	"github.com/codeableorg/weekend-challenge-13-GabNatali/internal/user"
)

type AuthService interface {
	Login(dto LoginUserDto) (LoggedUserDto, error)
	VerifyAccessToken(accessToken string) (uint, error)
	ParseAccessToken(accessToken string) (uint, error)
}

type AuthServiceOpts struct {
	UserRepository user.UserRepository
	Config         string
	Crypto         crypto.Crypto
}

func NewAuthService(opts AuthServiceOpts) AuthService {
	return &authService{
		UserRepository: opts.UserRepository,
		Config:         opts.Config,
		Crypto:         opts.Crypto,
	}
}

type authService struct {
	user.UserRepository
	Config string
	crypto.Crypto
}

func (u *authService) Login(in LoginUserDto) (out LoggedUserDto, err error) {

	user, err := u.UserRepository.GetByEmail(in.Email)
	if err != nil {
		return out, errors.New("invalid Credentials")
	}

	if !user.ComparePassword(in.Password) {
		return out, errors.New("invalid Credentials")
	}

	token, err := u.generateAccessToken(int64(user.Id))
	if err != nil {
		return out, err
	}

	return out.MapFromModel(user, token), nil

}

func (u *authService) VerifyAccessToken(accessToken string) (uint, error) {
	payload, err := u.ParseAndValidateJWT(accessToken, u.Config)
	if err != nil {
		return 0, errors.New("invalid token")
	}

	userId, ok := payload["userId"].(uint)
	if !ok {
		return 0, errors.New("unauthorized")
	}

	return userId, nil
}

func (u *authService) ParseAccessToken(accessToken string) (uint, error) {
	payload, err := u.ParseJWT(accessToken, u.Config)
	if err != nil {
		return 0, errors.New("invalid token")
	}

	userId, ok := payload["userId"].(uint)
	if !ok {
		return 0, errors.New("unauthorized")
	}

	return userId, nil
}

func (u *authService) generateAccessToken(userId int64) (string, error) {
	payload := map[string]interface{}{
		"userId": userId,
	}
	exp := time.Now().Add(time.Hour * 1)
	return u.GenerateJWT(
		payload,
		u.Config,
		exp,
	)
}
