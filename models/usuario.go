package models

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Status         string        `json:"status,omitempty"`
	Nome           string        `json:"nome,omitempty"`
	CPF            string        `json:"cpf,omitempty"`
	RG             string        `json:"rg,omitempty"`
	Genero         string        `json:"genero"`
	DataNascimento time.Time     `json:"data_nascimento"`
	Email          string        `json:"email,omitempty"`
	Senha          string        `json:"senha,omitempty"`
	Telefone       string        `json:"telefone,omitempty"`
	Escolaridade   string        `json:"escolaridade"`
	Profissao      string        `json:"profissao"`
	FotoPerfil     string        `json:"foto_perfil,omitempty"`
	TipoUsuarioID  int           `json:"tipo_usuario_id,omitempty"`
	InstituicaoID  int           `json:"instituicao_id,omitempty"`
	CidadeID       int           `json:"cidadeid,omitempty"`
	Instituicoes   []Instituicao `gorm:"many2many:instituicao_usuarios;"`
}

func (usuario *Usuario) Preparar() error {
	// Chama a função ValidarUsuario()
	err := usuario.ValidarUsuario()
	// Verifica se houve erros
	if err != nil {
		return err
	}
	// Retorna nil se não houver erros
	return nil
}

func (usuario Usuario) ValidarUsuario() error {
	cpfRegexp := regexp.MustCompile("[0-9]{3}.[0-9]{3}.[0-9]{3}-[0-9]{2}")
	rgRegexp := regexp.MustCompile("[0-9]{2}.?[0-9]{3}.?[0-9]{3}-?[0-9]{1}")
	nomere := regexp.MustCompile("^[a-zA-Záàâãéêíóôõúçñ ]+$")
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	// Valida o status
	if usuario.Status != "ativo" && usuario.Status != "inativo" {
		return errors.New("status inválido")
	}

	// Valida o nome
	if !nomere.MatchString(usuario.Nome) {
		return errors.New("nome inválido")
	}

	// Valida o CPF
	if !cpfRegexp.MatchString(usuario.CPF) {
		return errors.New("CPF inválido")
	}

	//Valida o RG

	if !rgRegexp.MatchString(usuario.RG) {
		return errors.New("RG inválido")
	}

	// Valida o gênero
	if usuario.Genero != "masculino" && usuario.Genero != "feminino" {
		return errors.New("gênero inválido")
	}

	// Valida a data de nascimento
	if usuario.DataNascimento.IsZero() {
		return errors.New("data de nascimento inválida")
	}

	// Valida o e-mail
	if !emailRegex.MatchString(usuario.Email) {
		return errors.New("e-mail inválido")

	}

	if len(usuario.Senha) > 20 {
        return errors.New("senha deve ter no máximo 20 caracteres")
    }

	// Valida o telefone
	if len(usuario.Telefone) < 8 {
		return errors.New("telefone deve ter pelo menos 8 caracteres")
	}

	// Valida a escolaridade
	if usuario.Escolaridade != "fundamental" && usuario.Escolaridade != "médio" && usuario.Escolaridade != "superior" {
		return errors.New("escolaridade inválida")
	}

	// Valida a profissão
	if len(usuario.Profissao) < 3 {
		return errors.New("profissão deve ter pelo menos 3 caracteres")
	}

	// Valida a foto de perfil
	if len(usuario.FotoPerfil) > 100 {
		return errors.New("foto de perfil deve ter no máximo 100 caracteres")
	}

	// Valida o ID do tipo de usuário
	if usuario.TipoUsuarioID < 1 {
		return errors.New("ID do tipo de usuário inválido")
	}

	// Valida o ID da instituição
	if usuario.InstituicaoID < 1 {
		return errors.New("ID da instituição inválido")
	}

	// Valida o ID da cidade
	if usuario.CidadeID < 1 {
		return errors.New("ID da cidade inválido")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.Trim(usuario.Nome, "")
	usuario.CPF = strings.Trim(usuario.CPF, "")
	usuario.Email = strings.Trim(usuario.Email, "")
	usuario.RG = strings.Trim(usuario.RG, "")
}
