package cobain

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
)

func GCFCreateCatalog(publickey, MONGOCONNSTRINGENV, dbname, colladmin, collcatalog string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var admindata Admin
	gettoken := r.Header.Get("token")
	if gettoken == "" {
		response.Message = "Missing token in headers"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		admindata.Email = checktoken
		if checktoken == "" {
			response.Message = "Invalid token"
		} else {
			admin2 := FindAdmin(mconn, colladmin, admindata)
			if admin2.Role == "admin" {
				var datacatalog Catalog
				err := json.NewDecoder(r.Body).Decode(&datacatalog)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()
				} else {
					CreateNewCatalog(mconn, collcatalog, Catalog{
						Nomorid:     datacatalog.Nomorid,
						Title:       datacatalog.Title,
						Description: datacatalog.Description,
						Image:       datacatalog.Image,
						Status:      datacatalog.Status,
					})
					response.Status = true
					response.Message = "Catalog creation successful"
				}
			} else {
				response.Message = "ANDA BUKAN ADMIN"
			}
		}
	}
	return GCFReturnStruct(response)
}

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
	return GCFReturnStruct(resp)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}
