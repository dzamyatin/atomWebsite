package servicemail

import (
	"crypto/tls"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

type IMailService interface {
	SendMail(to, subject, body string) error
}

// https://pkg.go.dev/gopkg.in/gomail.v2#example-package
type SmtpMailService struct {
	dialer *gomail.Dialer
	logger *zap.Logger
	sender string
}

func NewSmtpMailService(
	host string,
	port uint32,
	username string,
	password string,
	localName string,
	logger *zap.Logger,
	cert string,
	pk string,
	sender string,
	ssl bool,
) *SmtpMailService {
	d := gomail.NewDialer(
		host,
		int(port),
		username,
		password,
	)

	d.LocalName = localName

	/*
		rsa.PrivateKey{
			PublicKey:   rsa.PublicKey{},
			D:           nil,
			Primes:      nil,
			Precomputed: rsa.PrecomputedValues{},
		}
	*/

	if cert != "" && pk != "" {

		c, err := tls.X509KeyPair(
			[]byte(cert),
			[]byte(pk),
		)
		_ = c
		if err != nil {
			panic(err)
		}

		//b, _ := pem.Decode([]byte(cert))
		////c, err := ParseCertificate(b.Bytes)
		//certA, err := x509.ParseCertificate(b.Bytes)
		//if err != nil {
		//	panic(errors.Wrap(err, "failed to parse certificate"))
		//}

		d.TLSConfig = &tls.Config{
			//ServerName: localName,
			//Certificates: []tls.Certificate{
			//	c,
			//	//{
			//	//	Certificate: [][]byte{
			//	//		[]byte(cert),
			//	//	},
			//	//	//PrivateKey:  nil,
			//	//	//SupportedSignatureAlgorithms: nil,
			//	//	//OCSPStaple:                   nil,
			//	//	//SignedCertificateTimestamps:  nil,
			//	//	//Leaf:                         nil,
			//	//},
			//},
			InsecureSkipVerify: true,
			//VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			//	return nil
			//},
			//VerifyConnection: func(state tls.ConnectionState) error {
			//	return nil
			//},
		}
	}

	d.SSL = ssl

	return &SmtpMailService{
		dialer: d,
		logger: logger,
		sender: sender,
	}
}

func (r *SmtpMailService) SendMail(to, subject, body string) error {

	msg := gomail.NewMessage()
	msg.SetHeader("From", r.sender)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	//msg.SetAddressHeader("Cc", "dan@example.com", "Dan")
	msg.SetBody("text/html", body)
	msg.SetHeader("Mime-Version", "1.0")
	//m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	err := r.dialer.DialAndSend(msg)
	if err != nil {
		return errors.Wrap(err, "send mail")
	}

	return nil
}
