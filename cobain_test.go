package cobain

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewAdminRole(t *testing.T) {
	var admindata Admin
	admindata.Email = "ryaasishlah@gmail.com"
	admindata.Password = "mantap"
	admindata.Role = "admin"
	mconn := SetConnection("MONGOSTRING", "psbapk")
	CreateNewAdminRole(mconn, "admin", admindata)
}

func TestDeleteAdmin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "psbapk")
	var admindata Admin
	admindata.Email = "musa@gmail.com"
	DeleteAdmin(mconn, "admin", admindata)
}

func CreateNewAdminToken(t *testing.T) {
	var admindata Admin
	admindata.Email = "ryaasishlah@gmail.com"
	admindata.Password = "mantap"
	admindata.Role = "admin"

	// Create a MongoDB connection
	mconn := SetConnection("MONGOSTRING", "psbapk")

	// Call the function to create a admin and generate a token
	err := CreateAdminAndAddToken("your_private_key_env", mconn, "admin", admindata)

	if err != nil {
		t.Errorf("Error creating admin and token: %v", err)
	}
}

func TestGFCPostHandlerAdmin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "psbapk")
	var admindata Admin
	admindata.Email = "ryaasishlah@gmail.com"
	admindata.Password = "mantap"
	admindata.Role = "admin"
	CreateNewAdminRole(mconn, "admin", admindata)
}

func TestCatalog(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "psbapk")
	var catalogdata Catalog
	catalogdata.Nomorid = 1
	catalogdata.Title = "garut"
	catalogdata.Description = "keren banget"
	catalogdata.Image = "https://images3.alphacoders.com/165/thumb-1920-165265.jpg"
	CreateNewCatalog(mconn, "catalog", catalogdata)
}

func TestAllCatalog(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "psbapk")
	catalog := GetAllCatalog(mconn, "catalog")
	fmt.Println(catalog)
}

func TestGeneratePasswordHash(t *testing.T) {
	password := "ganteng"
	hash, _ := HashPass(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	match := CompareHashPass(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("alagaday", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "psbapk")
	var admindata Admin
	admindata.Email = "edi@gmail.com"
	admindata.Password = "pecin"

	filter := bson.M{"email": admindata.Email}
	res := atdb.GetOneDoc[Admin](mconn, "admin", filter)
	fmt.Println("Mongo Admin Result: ", res)
	hash, _ := HashPass(admindata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CompareHashPass(admindata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "psbapk")
	var admindata Admin
	admindata.Email = "bangsat"
	admindata.Password = "ganteng"

	anu := IsPasswordValid(mconn, "admin", admindata)
	fmt.Println(anu)
}

func TestAdminFix(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "psbapk")
	var admindata Admin
	admindata.Email = "ryaasishlah@gmail.com"
	admindata.Password = "mantap"
	admindata.Role = "admin"
	CreateAdmin(mconn, "admin", admindata)
}

func TestLoginn(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "psbapk")
	var admindata Admin
	admindata.Email = "ryaasishlah@gmail.com"
	admindata.Password = "mantap"
	IsPasswordValid(mconn, "admin", admindata)
	fmt.Println(admindata)
}
