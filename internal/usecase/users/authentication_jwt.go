package usecase

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/firerplayer/hexagonal-arch-go/internal/domain/gateway"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase/dto"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationJwtUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewAuthenticationJwtUseCase(usersGateway gateway.UsersGateway) *AuthenticationJwtUseCase {
	return &AuthenticationJwtUseCase{
		UsersGateway: usersGateway,
	}
}

func (uc *AuthenticationJwtUseCase) Execute(ctx context.Context, input dto.AuthenticateJwtUserInputDTO) (*dto.AuthenticateJwtUserOutputDTO, error) {
	// Lógica de autenticação do usuário
	// Verifique as credenciais fornecidas e, se forem válidas, gere um token JWT
	user, err := uc.UsersGateway.FindUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.New("user not found " + err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("Invalid password " + err.Error())
	}

	// Exemplo de geração de um token JWT com uma assinatura simples
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = input.Email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	// Defina outras informações relevantes no token, como tempo de expiração
	tokenString, err := token.SignedString([]byte(os.Getenv("jtw_secret_key")))
	if err != nil {
		return nil, errors.New("failed to generate token " + err.Error())
	}

	return &dto.AuthenticateJwtUserOutputDTO{
		Token: tokenString,
	}, nil

}
