package middleware

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"io"
	"strconv"
	"strings"
)

type DecryptionIdentifier struct {
	Token string
}

func (decryption DecryptionIdentifier) Decrypt() []byte {
	tokenDecode, err := base64.StdEncoding.DecodeString(decryption.Token)
	if err != nil {
		xtremeres.ErrXtremeUnauthenticated("Unable to decode token!!")
	}

	iv := tokenDecode[0:16]
	tokenExplode := strings.Split(string(tokenDecode[16:]), "-:-")

	secretLength, _ := strconv.Atoi(tokenExplode[0])
	secret := []byte(tokenExplode[1][0:secretLength])
	identifierData := []byte(tokenExplode[1][secretLength:])

	block, err := aes.NewCipher(secret)
	if err != nil {
		xtremeres.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decode token!! %s", err))
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(identifierData, identifierData)
	identifierData = unpadPKCS7(identifierData)

	identifierData, err = decompressZlib(identifierData)
	if err != nil {
		xtremeres.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decompress data: %s", err))
	}

	return identifierData
}

func decompressZlib(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	resData = resB.Bytes()

	return
}

func unpadPKCS7(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
