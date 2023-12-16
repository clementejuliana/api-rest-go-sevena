package models

import "time"

type Login struct {
    ID        int    `json:"id"`
    UsuarioEmail string `json:"usuario_email"`
    Status    string `json:"status"`
    ExpirationData time.Time `json:"expiration_data"`
    Email string `json:"email"`
    Senha string `json:"senha"`
}

