package auth

import (
	"context"
	"errors"
	"time"

	"github.com/codeableorg/weekend-challenge-13-GabNatali/internal/user"
	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	Login(ctx context.Context, dto LoginUserDto) (LoggedUserDto, error)
	// VerifyAccessToken(accessToken string) (int64, error)
	// ParseAccessToken(accessToken string) (int64, error)
}

type AuthServiceOpts struct {
	UserRepository user.UserRepository
	Config         string
}

func NewAuthService(opts AuthServiceOpts) AuthService {
	return &authService{
		UserRepository: opts.UserRepository,
		Config:         opts.Config,
	}
}

type authService struct {
	user.UserRepository
	Config string
}

func (u *authService) Login(ctx context.Context, in LoginUserDto) (out LoggedUserDto, err error) {

	user, err := u.UserRepository.GetByEmail(ctx, in.Email)
	if err != nil {
		return out, errors.New("invalid Credentials")
	}

	if !user.ComparePassword(in.Password) {
		return out, errors.New("invalid Credentials")
	}

	claims := make(jwt.MapClaims)
	claims["userId"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.Config))
	if err != nil {
		return out, err
	}

	return out.MapFromModel(user, tokenString), nil

}

// func (u *authService) VerifyAccessToken(accessToken string) (int64, error) {
// 	payload, err := u.ParseAndValidateJWT(accessToken, u.AccessTokenSecret())
// 	if err != nil {
// 		return 0, errors.New(errors.UnauthorizedError, "")
// 	}

// 	userId, ok := payload["userId"].(float64)
// 	if !ok {
// 		return 0, errors.New(errors.UnauthorizedError, "")
// 	}

// 	return int64(userId), nil
// }

// func (u *authService) ParseAccessToken(accessToken string) (int64, error) {
// 	payload, err := u.ParseJWT(accessToken, u.AccessTokenSecret())
// 	if err != nil {
// 		return 0, errors.New(errors.UnauthorizedError, "")
// 	}

// 	userId, ok := payload["userId"].(float64)
// 	if !ok {
// 		return 0, errors.New(errors.UnauthorizedError, "")
// 	}

// 	return int64(userId), nil
// }

// func (u *authService) generateAccessToken(userId int64) (string, error) {
// 	payload := map[string]interface{}{"userId": userId}

// 	return u.GenerateJWT(
// 		payload,
// 		u.AccessTokenSecret(),
// 		u.AccessTokenExpiresDate(),
// 	)
// }
