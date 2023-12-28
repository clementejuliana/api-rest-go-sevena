package authentication

import (
    "errors"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"

    jwt "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

func CriarToken(usuarioID uint64) (string, error) {
    permissoes := jwt.MapClaims{}
    permissoes["authorized"] = true
    permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
    permissoes["usuarioId"] = usuarioID
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
    return token.SignedString([]byte(SECRET_KEY))
}

func ValidarToken() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := extrairToken(c.Request)
        token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
        if erro != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": erro.Error()})
            c.Abort()
            return
        }

        if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Next()
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            c.Abort()
        }
    }
}

func ExtrairUsuarioID(c *gin.Context) (uint64, error) {
    tokenString := extrairToken(c.Request)
    token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
    if erro != nil {
        return 0, erro
    }

    if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
        if erro != nil {
            return 0, erro
        }

        return usuarioID, nil
    }

    return 0, errors.New("Token inválido")
}

func extrairToken(r *http.Request) string {
    token := r.Header.Get("Authorization")

    if len(strings.Split(token, " ")) == 2 {
        return strings.Split(token, " ")[1]
    }

    return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
    }

    return SECRET_KEY, nil
}
