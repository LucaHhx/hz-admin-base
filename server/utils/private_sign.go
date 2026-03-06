package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"sort"
	"strings"
)

func GenerateRSAKeys() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func StringToSHA256(s string) string {
	hash := sha256.Sum256([]byte(s))   // 计算哈希值
	return hex.EncodeToString(hash[:]) // 转换为十六进制字符串
}

func MapToSignStr(data map[string]interface{}) string {
	keys := make([]string, 0)
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	list := make([]string, 0)
	for _, k := range keys {
		value := StringMust(data[k])
		if value == "" {
			continue
		}
		list = append(list, k+"="+StringMust(data[k]))
	}
	joinStr := strings.Join(list, "&")
	fmt.Println("joinStr:", joinStr)
	return StringToSHA256(joinStr)
}

func DataToMsg(data any) string {
	switch data.(type) {
	case string:
		return StringToSHA256(data.(string))
	case []byte:
		return StringToSHA256(string(data.([]byte)))
	case map[string]interface{}:
		return MapToSignStr(data.(map[string]interface{}))
	default:
		sMap := StructToSpreadMap(data)
		return MapToSignStr(sMap)
	}
}

func ExportPrivateKey(privateKey *rsa.PrivateKey) string {
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)
	privPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privDER})
	return string(privPEM)
}

func ExportPublicKey(publicKey *rsa.PublicKey) string {
	pubDER, _ := x509.MarshalPKIXPublicKey(publicKey)
	pubPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubDER})
	return string(pubPEM)
}

func FmtPrivateKey(privBase64 string) (string, error) {
	if len(privBase64) < 5 {
		return "", errors.New("failed to parse public key PEM")
	}
	if privBase64[0:5] != "-----" {
		privBase64 = "-----BEGIN RSA PRIVATE KEY-----\n" + privBase64 + "\n-----END RSA PRIVATE KEY-----"
	}
	block, _ := pem.Decode([]byte(privBase64))
	if block == nil {
		return "", errors.New("failed to parse private key PEM")
	}
	return privBase64, nil
}

func ImportPrivateKey(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse private key PEM")
	}

	switch block.Type {
	case "RSA PRIVATE KEY": // PKCS#1
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY": // PKCS#8
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaKey, ok := key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("not an RSA private key")
		}
		return rsaKey, nil
	default:
		return nil, errors.New("unknown private key format")
	}
}

func ImportPublicKey(pubBase64 string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubBase64))
	if block == nil {
		return nil, errors.New("failed to parse public key PEM")
	}
	switch block.Type {
	case "PUBLIC KEY":
		pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaPubKey, ok := pubKey.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("not an RSA public key")
		}
		return rsaPubKey, nil
	case "RSA PUBLIC KEY":
		return x509.ParsePKCS1PublicKey(block.Bytes)
	default:
		return nil, errors.New("unknown public key format")
	}
}

// SignWithPrivateKey 使用私钥对消息签名
func SignWithPrivateKey(privateKey *rsa.PrivateKey, message string) (string, error) {
	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifyWithPublicKey 使用公钥验证签名
func VerifyWithPublicKey(publicKey *rsa.PublicKey, message string, signatureBase64 string) error {
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return err
	}
	fmt.Println("Go message:", message)
	hashed := sha256.Sum256([]byte(message))
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
}

// EncryptWithPublicKey 使用公钥加密数据
func EncryptWithPublicKey(publicKey *rsa.PublicKey, message string) (string, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptWithPrivateKey 使用私钥解密数据
func DecryptWithPrivateKey(privateKey *rsa.PrivateKey, ciphertextBase64 string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return "", err
	}
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
func equalHashes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
