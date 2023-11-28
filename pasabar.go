package cobain

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// catalog
func InsertDataCatalog(MongoConn *mongo.Database, colname string, cat Catalog) (InsertedID interface{}) {
	req := new(Catalog)
	req.CatalogId = cat.CatalogId
	req.Title = cat.Title
	req.Description = cat.Description
	req.Image = cat.Image
	req.Status = cat.Status
	return InsertOneDoc(MongoConn, colname, req)
}

func InsertAdmindata(MongoConn *mongo.Database, email, role, password string) (InsertedID interface{}) {
	req := new(Admin)
	req.Email = email
	req.Password = password
	req.Role = role
	return InsertSatuDoc(MongoConn, "admin", req)
}

func InsertSatuDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
