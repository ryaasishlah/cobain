package cobain

import (
	"fmt"
	"testing"

	"github.com/whatsauth/watoken"
)

var privatekey = "privatekey"
var publickeyb = "publickey"
var encode = "encode"

func TestGenerateKeyPASETO(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("alagaday", privateKey)
	fmt.Println(hasil, err)
}

func TestHashPass(t *testing.T) {
	password := "bebas"

	Hashedpass, err := HashPass(password)
	fmt.Println("error : ", err)
	fmt.Println("Hash : ", Hashedpass)
}

func TestHashFunc(t *testing.T) {
	conn := MongoCreateConnection("MONGOSTRING", "Coba")
	admindata := new(Admin)
	admindata.Email = "coba@gmail.com"
	admindata.Password = "bebas"

	data := GetOneAdmin(conn, "admin", Admin{
		Email:    admindata.Email,
		Password: admindata.Password,
	})
	fmt.Printf("%+v", data)
	fmt.Println(" ")
	hashpass, _ := HashPass(admindata.Password)
	fmt.Println("Hasil hash : ", hashpass)
	compared := CompareHashPass(admindata.Password, data.Password)
	fmt.Println("result : ", compared)
}

func TestTokenEncoder(t *testing.T) {
	conn := MongoCreateConnection("MONGOSTRING", "Coba")
	privateKey, publicKey := watoken.GenerateKey()
	admindata := new(Admin)
	admindata.Email = "coba@gmail.com"
	admindata.Password = "bebas"

	data := GetOneAdmin(conn, "admin", Admin{
		Email:    admindata.Email,
		Password: admindata.Password,
	})
	fmt.Println("Private Key : ", privateKey)
	fmt.Println("Public Key : ", publicKey)
	fmt.Printf("%+v", data)
	fmt.Println(" ")

	encode := TokenEncoder(data.Email, privateKey)
	fmt.Printf("%+v", encode)
}

func TestInsertAdmindata(t *testing.T) {
	conn := MongoCreateConnection("MONGOSTRING", "Coba")
	password, err := HashPass("iyasganteng")
	fmt.Println("err", err)
	data := InsertAdmindata(conn, "iyas", "role", password)
	fmt.Println(data)
}

func TestDecodeToken(t *testing.T) {
	deco := watoken.DecodeGetId("public",
		"token")
	fmt.Println(deco)
}

func TestCompareEmail(t *testing.T) {
	conn := MongoCreateConnection("MONGOSTRING", "Coba")
	deco := watoken.DecodeGetId("public",
		"token")
	compare := CompareEmail(conn, "admin", deco)
	fmt.Println(compare)
}

func TestEncodeWithRole(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	role := "admin"
	Email := "coba@gmail.com"
	encoder, err := EncodeWithRole(role, Email, privateKey)

	fmt.Println(" error :", err)
	fmt.Println("Private :", privateKey)
	fmt.Println("Public :", publicKey)
	fmt.Println("encode: ", encoder)

}

func TestDecoder2(t *testing.T) {
	pay, err := Decoder(publickeyb, encode)
	admin, _ := DecodeGetAdmin(publickeyb, encode)
	role, _ := DecodeGetRole(publickeyb, encode)
	use, ro := DecodeGetRoleandAdmin(publickeyb, encode)
	fmt.Println("admin :", admin)
	fmt.Println("role :", role)
	fmt.Println("admin and role :", use, ro)
	fmt.Println("err : ", err)
	fmt.Println("payload : ", pay)
}
