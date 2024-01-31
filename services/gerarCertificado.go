package services

import (
	"bytes"
	"io/ioutil"
	"net/smtp"

	"github.com/jung-kurt/gofpdf"
)

// Gerar um certificado PDF simples
func GerarCertificadoPDF(nome string, atividade string) (*bytes.Buffer, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Certificado de Participação")
	pdf.Ln(20)
	pdf.Cell(40, 10, "Este certificado é concedido a "+nome)
	pdf.Ln(10)
	pdf.Cell(40, 10, "por participar da atividade "+atividade)
	pdf.Ln(10)
	pdf.Cell(40, 10, "Parabéns!")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

// Salvar o certificado em um local temporário
func SalvarCertificado(certificado *bytes.Buffer, caminho string) error {
	return ioutil.WriteFile(caminho, certificado.Bytes(), 0644)
}

// Enviar o certificado por e-mail
func EnviarCertificadoPorEmail(caminho string, email string) error {
	from := "seu-email@gmail.com"
	pass := "sua-senha"
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Seu Certificado\n\n" +
		"Aqui está o seu certificado."

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	return err
}