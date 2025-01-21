package security

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"filippo.io/age"
	"filippo.io/age/armor"
)

// Vault encapsulates encryption and decryption logic
type AgeKeyPair struct {
	privateKey age.Identity
	publicKey  age.Recipient
}

// NewVault initializes a new Vault with the given private and public keys
func NewAgeKeyPair(privateKeySource, publicKeyPath string) (*AgeKeyPair, error) {
	privateKey, err := loadPrivateKey(privateKeySource)
	if err != nil {
		return nil, err
	}

	publicKey, err := loadPublicKey(publicKeyPath)
	if err != nil {
		return nil, err
	}

	return &AgeKeyPair{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// Encrypt encrypts data using the public key
func (m *AgeKeyPair) Encrypt(data []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	armorWriter := armor.NewWriter(buf)
	w, err := age.Encrypt(armorWriter, m.publicKey)
	if err != nil {
		return nil, err
	}

	_, err = w.Write(data)
	if err != nil {
		return nil, err
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}

	err = armorWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decrypt decrypts encrypted data using the private key
func (m *AgeKeyPair) Decrypt(encryptedData []byte) ([]byte, error) {
	buf := bytes.NewReader(encryptedData)
	armorReader := armor.NewReader(buf)
	rc, err := age.Decrypt(armorReader, m.privateKey)
	if err != nil {
		return nil, err
	}
	decryptedData, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil
}

// loadPrivateKey loads the private key from a file, command-line argument, or environment variable
func loadPrivateKey(source string) (age.Identity, error) {
	var keyData []byte
	var err error

	// Check if source is a file
	if _, err = os.Stat(source); err == nil {
		keyData, err = os.ReadFile(source)
		//fmt.Println(string(keyData))
		if err != nil {
			return nil, fmt.Errorf("failed to read private key from file: %w", err)
		}
	} //else if value, found := os.LookupEnv(source); found {
	//	// Check if source is an environment variable
	//	keyData = []byte(value)
	//} else {
	//	// Fallback: treat as direct key input
	//	keyData = []byte(source)
	//}

	return age.ParseX25519Identity(string(keyData))
}

// loadPublicKey loads the public key from the file
func loadPublicKey(path string) (age.Recipient, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key from file: %w", err)
	}
	return age.ParseX25519Recipient(string(keyData))
}
