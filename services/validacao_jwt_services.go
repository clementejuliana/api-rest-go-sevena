package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/smtp"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "secret-key",
		issure:    "api-sevena",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GeraToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidaToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("token inválido")
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}

func GeraTokenRecuperacaoSenha() (string, error) {
    tokenBytes := make([]byte, 32)
    if _, err := rand.Read(tokenBytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(tokenBytes), nil
}

func EnviaEmailRecuperacaoSenha(destinatario, linkRecuperacao string) error {
    // Configurar as informações do servidor de e-mail
    from := "seuemail@gmail.com"
    password := "suasenha"
    host := "smtp.gmail.com"
    port := 587

    // Configurar a mensagem de e-mail
    assunto := "Recuperação de Senha"
    mensagem := fmt.Sprintf("Clique no link para recuperar sua senha: %s", linkRecuperacao)

    // Configurar autenticação
    auth := smtp.PlainAuth("", from, password, host)

    // Configurar e enviar e-mail
    err := smtp.SendMail(fmt.Sprintf("%s:%d", host, port), auth, from, []string{destinatario}, []byte(fmt.Sprintf("Subject: %s\n\n%s", assunto, mensagem)))
    if err != nil {
        return err
    }

    return nil
}
