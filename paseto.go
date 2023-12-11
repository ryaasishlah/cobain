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

func Registrasi(token, mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var dataadmin Admin
	err := json.NewDecoder(r.Body).Decode(&dataadmin)
	if emailExists(mongoenv, dbname, dataadmin) {
		response.Message = "Username telah dipakai"
	} else {
		if err != nil {
			response.Message = "error parsing application/json: " + err.Error()
		} else {
			hash, hashErr := HashPass(dataadmin.Password)
			if hashErr != nil {
				response.Message = "Gagal Hash Password" + err.Error()
			}
			InsertAdmindata(mconn, collname, dataadmin.Email, hash, dataadmin.No_whatsapp)
			response.Message = "Berhasil Input data"

			var email = dataadmin.Email
			var password = dataadmin.Password
			var nohp = dataadmin.No_whatsapp

			dt := &wa.TextMessage{
				To:       nohp,
				IsGroup:  false,
				Messages: "Selamat anda berhasil registrasi, berikut adalah username anda: " + email + " dan ini adalah password anda: " + password + "\nDisimpan baik baik ya",
			}

			atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv(token), dt, "https://api.wa.my.id/api/send/message/text")
		}
	}
	return GCFReturnStruct(response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func Login(token, Privatekey, MongoEnv, dbname, Colname string, r *http.Request) string {
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

			var email = dataadmin.Email
			var nohp = dataadmin.No_whatsapp

			dt := &wa.TextMessage{
				To:       nohp,
				IsGroup:  false,
				Messages: "Selamat datang Admin PASABAR anda berhasil Login, anda masuk menggunakan: " + email + "\n Selamat menggunakanya ya",
			}

			atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv(token), dt, "https://api.wa.my.id/api/send/message/text")
		} else {
			resp.Message = "Password Salah"
		}
	}
	return ReturnStringStruct(resp)
}

func LoginBaru(Privatekey, MongoEnv, dbname, Colname string, r *http.Request) string {
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
			if IsPasswordValid(mconn, Colname, Admin{
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
				dt := &wa.TextMessage{
					To:       datarole.PhoneNum,
					IsGroup:  false,
					Messages: "Hai hai kak \n Ini OTP kakak " + data.OTPCode,
				}
				res, _ := atapi.PostStructWithToken[Responses]("Token", os.Getenv(TOKEN), dt, "https://api.wa.my.id/api/send/message/text")
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

func LoginOTP2(MongoEnv, dbname, Colname string, r *http.Request) string {
	var resp Credential
	mconn := MongoCreateConnection(MongoEnv, dbname)
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
				dt := &wa.TextMessage{
					To:       datarole.PhoneNum,
					IsGroup:  false,
					Messages: "Hai hai kak \n Ini OTP kakak " + data.OTPCode,
				}
				res, _ := atapi.PostStructWithToken[Responses]("Token", os.Getenv("TOKEN"), dt, "https://asia-southeast2-pasabar.cloudfunctions.net/webhook")
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
