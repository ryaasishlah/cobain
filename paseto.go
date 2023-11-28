package cobain

import (
	"encoding/json"
	json2 "encoding/json"

	"net/http"
	"os"

	"github.com/whatsauth/watoken"
)

func ReturnStringStruct(Data any) string {
	json, _ := json2.Marshal(Data)
	return string(json)
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

// Insert data
func InsertCatalog(MongoEnv, dbname, colname, publickey string, r *http.Request) string {
	resp := new(Credential)
	req := new(Catalog)
	conn := MongoCreateConnection(MongoEnv, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Found"
	} else {
		checkadmin := IsAdmin(tokenlogin, os.Getenv(publickey))
		if !checkadmin {

			resp.Status = false
			resp.Message = "Anda tidak bisa Insert data karena bukan HR atau admin"

		} else {
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				resp.Message = "error parsing application/json: " + err.Error()
			} else {
				pass, err := HashPass(req.Account.Password)
				if err != nil {
					resp.Status = false
					resp.Message = "Gagal Hash Code"
				}
				InsertDataCatalog(conn, colname, Catalog{
					CatalogId:   req.CatalogId,
					Title:       req.Title,
					Description: req.Description,
					Image:       req.Image,
					Status:      req.Status,
				})
				InsertAdmindata(conn, req.Account.Email, req.Account.Role, pass)
				resp.Status = true
				resp.Message = "Berhasil Insert data"
			}
		}
	}
	return ReturnStringStruct(resp)
}
