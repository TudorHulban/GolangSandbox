package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

const organization = "xxxx"

type simpleHandler struct{}

func (h *simpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "alive")
}

func main() {
	bigInteger := new(big.Int).Lsh(big.NewInt(1), 256)
	log.Println("bigInteger:", bigInteger)

	randomNumber, _ := rand.Int(rand.Reader, bigInteger) // TODO: no seed needed?
	log.Println("serialNumber:", randomNumber)

	pkiSubject := pkix.Name{
		Country:      []string{"Romania"},
		Organization: []string{organization},
	}

	template := x509.Certificate{
		SerialNumber: randomNumber,
		Subject:      pkiSubject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	bytesDER, _ := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)

	certifPEM, _ := os.Create("cert.pem")
	pem.Encode(certifPEM, &pem.Block{Type: "CERTIFICATE", Bytes: bytesDER})
	certifPEM.Close()

	certifKEY, _ := os.Create("key.pem")
	pem.Encode(certifKEY, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	certifKEY.Close()

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &simpleHandler{},
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
