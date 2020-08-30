package k8sFirstSteps

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Data struct {
	Payload     string
	GeneratedAt time.Time
}

func NewEncoder() Encoder {
	return Encoder{block: getBlockCipher()}
}

func NewDecoder() Decoder {
	return Decoder{block: getBlockCipher()}
}

func getBlockCipher() cipher.Block {
	secret := getEncodingSecret()
	block, err := aes.NewCipher(secret)
	if err != nil {
		panic(err)
	}

	return block
}

func getEncodingSecret() []byte {
	f, err := os.Open(os.Getenv("ENCRYPTION_SECRET_PATH"))
	if err != nil {
		log.Println("WARNING: Unable to open secret file: " + err.Error())
	}
	keyString, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("WARNING: Unable to read secret file: " + err.Error())
	}

	// This should never be in the code
	key, err := hex.DecodeString(string(keyString))
	if err != nil || len(key) == 0 {
		log.Println("WARNING: Unable to use secret - using default")
		key, _ = hex.DecodeString("6368616e676520746869732070617373")
	}

	return key
}

type Encoder struct {
	block cipher.Block
}

func (e Encoder) Encode(value string) (string, error) {
	plaintext := []byte(value)

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(e.block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	return fmt.Sprintf("%x\n", ciphertext), nil
}

type Decoder struct {
	block cipher.Block
}

func (d Decoder) Decode(value string) (string, error) {
	ciphertext, _ := hex.DecodeString(value)

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(d.block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext), nil
}
