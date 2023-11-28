package cobain

import (
	"time"
)

type Admin struct {
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"`
	Private  string `json:"private,omitempty" bson:"private,omitempty"`
	Public   string `json:"public,omitempty" bson:"public,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type ResponseEncode struct {
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
}

type ResponseBack struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type ResponseCatalog struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Catalog `json:"data"`
}

type ResponseCatalogBanyak struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Catalog `json:"data"`
}

type Catalog struct {
	CatalogId   string `json:"employeeid" bson:"employeeid,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Image       string `json:"image" bson:"image"`
	Status      bool   `json:"status" bson:"status"`
	Account     Admin  `json:"account" bson:"account,omitempty"`
}

type Cred struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ReqAdmins struct {
	Email string `json:"email"`
}

type Payload struct {
	Admin string    `json:"admin"`
	Role  string    `json:"role"`
	Exp   time.Time `json:"exp"`
	Iat   time.Time `json:"iat"`
	Nbf   time.Time `json:"nbf"`
}
