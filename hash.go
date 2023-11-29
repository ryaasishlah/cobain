package cobain

import (
	"github.com/whatsauth/watoken"
	"golang.org/x/crypto/bcrypt"
)

func HashPass(password string) (string, error) {
	bytess, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytess), err
}

func CompareHashPass(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TokenEncoder(email, privatekey string) string {
	resp := new(ResponseEncode)
	encode, err := watoken.Encode(email, privatekey)
	if err != nil {
		resp.Message = "Gagal Encode" + err.Error()
	} else {
		resp.Token = encode
		resp.Message = "Welcome iyaszz"
	}

	return ReturnStringStruct(resp)
}
