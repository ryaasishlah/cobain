package cobain

import (
	"context"
	"fmt"
	"os"

	"github.com/aiteung/atdb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		// DBString: "mongodb+srv://ryaasishlah123:ryaas123@ryaas.x1atjad.mongodb.net/", //os.Getenv(MONGOCONNSTRINGENV),
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func InsertAdmindata(MongoConn *mongo.Database, email, role, password string) (InsertedID interface{}) {
	req := new(Admin)
	req.Email = email
	req.Password = password
	req.Role = role
	return InsertOneDoc(MongoConn, "admin", req)
}

func InsertAdminsdata(MongoConn *mongo.Database, admin Admins) (InsertedID interface{}) {
	return InsertOneDoc(MongoConn, "admins", admin)
}

func DeleteAdmin(mongoconn *mongo.Database, collection string, admindata Admin) interface{} {
	filter := bson.M{"email": admindata.Email}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
}

func FindAdmin(mongoconn *mongo.Database, collection string, admindata Admin) Admin {
	filter := bson.M{"email": admindata.Email}
	return atdb.GetOneDoc[Admin](mongoconn, collection, filter)
}

func IsPasswordValid(mongoconn *mongo.Database, collection string, admindata Admin) bool {
	filter := bson.M{"email": admindata.Email}
	res := atdb.GetOneDoc[Admin](mongoconn, collection, filter)
	return CompareHashPass(admindata.Password, res.Password)
}

func MongoCreateConnection(MongoString, dbname string) *mongo.Database {
	MongoInfo := atdb.DBInfo{
		DBString: os.Getenv(MongoString),
		DBName:   dbname,
	}
	conn := atdb.MongoConnect(MongoInfo)
	return conn
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetOneAdmin(MongoConn *mongo.Database, colname string, admindata Admins) Admins {
	filter := bson.M{"email": admindata.Email}
	data := atdb.GetOneDoc[Admins](MongoConn, colname, filter)
	return data
}

func PasswordValidator(MongoConn *mongo.Database, colname string, admindata Admin) bool {
	filter := bson.M{"email": admindata.Email}
	data := atdb.GetOneDoc[Admin](MongoConn, colname, filter)
	hashChecker := CompareHashPass(admindata.Password, data.Password)
	return hashChecker
}
