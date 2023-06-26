package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationJwtUseCase struct {
	UsersGateway gateway.UsersGateway
	JwtSecretKey string
}

func NewAuthenticationJwtUseCase(usersGateway gateway.UsersGateway, jwtSecretKey string) *AuthenticationJwtUseCase {
	return &AuthenticationJwtUseCase{
		UsersGateway: usersGateway,
		JwtSecretKey: jwtSecretKey,
	}
}

// Execute authenticates the user with the provided credentials and generates a JWT token.
//
// ctx - The context in which the function is executed.
// input - The input data transfer object containing user credentials.
// Returns an output data transfer object containing a JWT token and/or an error in case of failure.
func (uc *AuthenticationJwtUseCase) Execute(ctx context.Context, input dto.AuthenticateJwtUserInputDTO) (*dto.AuthenticateJwtUserOutputDTO, error) {
	// Lógica de autenticação do usuário
	// Verifique as credenciais fornecidas e, se forem válidas, gere um token JWT
	user, err := uc.UsersGateway.FindUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.New("user not found -> " + err.Error())
	}
	// Password do usuário vem sempre em hash, então usamos bcrypt para verificar a senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid password -> " + err.Error())
	}
	exp_date := time.Now().Add(time.Hour * 24 * 365)

	// Gera a assinatura JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"exp":     exp_date.Unix(),
	})
	tokenString, err := token.SignedString([]byte(uc.JwtSecretKey))
	if err != nil {
		return nil, errors.New("failed to generate token -> " + err.Error())
	}

	return &dto.AuthenticateJwtUserOutputDTO{
		Token: tokenString,
	}, nil

}
