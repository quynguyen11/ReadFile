package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/pkcs12"
)

var (
	// ErrInvalidPrivateKey p12 file invalid
	ErrInvalidPrivateKey = errors.New("invalid file")
)

// Keypair :
type Keypair struct {
	PrivateKey []byte
	PublicKey  []byte
}

//RSAGenerateKey : generate publickey (DER-encoded PKIX format), privatekey (PKCS#8 encoded form, see RFC 5208).
func RSAGenerateKey() (kp *Keypair, err error) {
	priKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	privateKey, err := x509.MarshalPKCS8PrivateKey(priKey)
	if err != nil {
		return nil, err
	}
	publicKey, err := x509.MarshalPKIXPublicKey(&priKey.PublicKey)
	if err != nil {
		return nil, err
	}
	kp = &Keypair{privateKey, publicKey}
	return kp, nil
}

//RSAEncryptOAEP : encrypts the given message with RSA-OAEP, publickey in DER-encoded PKIX format.
func RSAEncryptOAEP(publicKey, data []byte) (output []byte, err error) {

	pub, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		log.Println("EncryptRSA", err)
		return output, err
	}
	pubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return output, fmt.Errorf("pubkey isn't a rsa pubkey")
	}

	label := []byte("")
	output, err = rsa.EncryptOAEP(sha1.New(), rand.Reader, pubKey, data, label)
	if err != nil {
		return output, err
	}
	return output, nil
}

//RSADecryptOAEP decrypts ciphertext using RSA-OAEP, privateKey in PKCS#8 encoded form.
func RSADecryptOAEP(ciphertext, privateKey []byte) (output []byte, err error) {

	priKey, err := x509.ParsePKCS8PrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	privkey, ok := priKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("privkey wrong type")
	}
	label := []byte("")
	output, err = rsa.DecryptOAEP(sha1.New(), rand.Reader, privkey, ciphertext, label)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return output, nil
}

// RSADecodePKCS12PriKey load rsa private key from p12 file
func RSADecodePKCS12PriKey(privateKeypath string, password string) (*rsa.PrivateKey, error) {
	pfxData, err := ioutil.ReadFile(privateKeypath)
	if err != nil {
		return nil, err
	}
	priKey, _, err := pkcs12.Decode(pfxData, password)
	if err != nil {
		return nil, err
	}
	rsaPriKey, ok := priKey.(*rsa.PrivateKey)
	if !ok {
		return nil, ErrInvalidPrivateKey
	}
	return rsaPriKey, nil
}

// RSASignMessage sign message, return base64 signature
func RSASignMessage(message []byte, signKey []byte) (signature string, err error) {
	priKey, err := x509.ParsePKCS8PrivateKey([]byte(signKey))
	if err != nil {
		return "", err
	}
	privkey, ok := priKey.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("privkey wrong type")
	}
	h := sha256.New()
	h.Write(message)
	d := h.Sum(nil)
	sig, err := rsa.SignPKCS1v15(rand.Reader, privkey, crypto.SHA256, d)
	if err != nil {
		return "", err
	}
	signature = base64.StdEncoding.EncodeToString(sig)
	return signature, nil
}

// RSAVerifySignature veriify signature, return nil is signature valid, sig is signature base64
// sig base64 encoding, message raw messsage
func RSAVerifySignature(sig string, message []byte, verifyKey []byte) error {
	pub, err := x509.ParsePKIXPublicKey(verifyKey)
	if err != nil {
		log.Println("EncryptRSA", err)
		return err
	}
	pubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("pubkey isn't a rsa pubkey")
	}
	h := sha256.New()
	h.Write(message)
	d := h.Sum(nil)
	sigBytes, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, d, sigBytes)
}

