package cobain

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/whatsauth/wa"
	"github.com/whatsauth/watoken"
)

func Login(Privatekey, MongoEnv, dbname, Colname string, r *http.Request) string {
	var resp Credential
	mconn := SetConnection(MongoEnv, dbname)
	var dataadmin Admin
	err := json.NewDecoder(r.Body).Decode(&dataadmin)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, Colname, dataadmin) {
			tokenstring, err := watoken.Encode(dataadmin.Email, os.Getenv(Privatekey))
			if err != nil {
				resp.Message = "Gagal Encode Token : " + err.Error()
			} else {
				resp.Status = true
				resp.Message = "Selamat Datang SUPERADMIN"
				resp.Token = tokenstring
			}
		} else {
			resp.Message = "Password Salah"
		}
	}
	return ReturnStringStruct(resp)
}

func ReturnStringStruct(Data any) string {
	json, _ := json.Marshal(Data)
	return string(json)
}

func LoginOTP(TOKEN, MongoEnv, dbname, Colname string, r *http.Request) string {
	var resp Credential
	mconn := SetConnection(MongoEnv, dbname)
	var dataadmin Admins
	err := json.NewDecoder(r.Body).Decode(&dataadmin)
	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		if err != nil {
			resp.Message = "error parsing application/json: " + err.Error()
		} else {
			if PasswordValidator(mconn, Colname, Admin{
				Email:    dataadmin.Email,
				Password: dataadmin.Password,
				Role:     dataadmin.Role,
			}) {
				datarole := GetOneAdmin(mconn, "admin", Admins{Email: dataadmin.Email})
				data := OTP{
					Email:   dataadmin.Email,
					Role:    datarole.Role,
					DateOTP: time.Now(),
					OTPCode: CreateOTP(),
				}
				InsertOtp(mconn, "otp", data)

				var nohp = dataadmin.PhoneNum

				dt := &wa.TextMessage{
					To:       nohp,
					IsGroup:  false,
					Messages: "Hai hai kak \n Ini OTP kakak " + data.OTPCode,
				}
				res, _ := atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv(TOKEN), dt, "https://api.wa.my.id/api/send/message/text")
				resp.Status = true
				resp.Message = "Hai Silahkan cek WhatsApp untuk OTPnya yaa"
				resp.Token = res.Response
			} else {
				resp.Message = "Password Salah"
			}
		}
	}
	return ReturnStringStruct(resp)
}
