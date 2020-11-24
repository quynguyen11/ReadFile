package crypto

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

//StringToSha256 :
func StringToSha256(text string) string {
	textSha256 := sha256.Sum256([]byte(text))
	return string(textSha256[:])

}

//HashPwd : string format sha256 pwd
func HashPwd(text string) string {
	textSha256 := sha256.Sum256([]byte(text))
	// return base64.URLEncoding.EncodeToString(textSha256[:])
	return hex.EncodeToString(textSha256[:])
}

//HashString : return hash string
func HashString(data []byte) string {
	h := sha1.New()
	h.Write(data)
	hByte := h.Sum(nil)

	return fmt.Sprintf("%x", hByte)
}

//HashByte : return hash string
func HashByte(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hByte := h.Sum(nil)
	return hByte
}
