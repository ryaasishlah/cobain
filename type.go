package cobain

import (
	"time"
)

type Admin struct {
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"`
	Token    string `json:"token,omitempty" bson:"token,omitempty"`
	Private  string `json:"private,omitempty" bson:"private,omitempty"`
	Public   string `json:"public,omitempty" bson:"public,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type Payload struct {
	User string    `json:"user"`
	Role string    `json:"role"`
	Exp  time.Time `json:"exp"`
	Iat  time.Time `json:"iat"`
	Nbf  time.Time `json:"nbf"`
}

type Responses struct {
	Response string `bson:"response" json:"response"`
}

type Admins struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	PhoneNum string `json:"phone-num" bson:"phone-num"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"`
}

type OTP struct {
	Email   string    `json:"email" bson:"email"`
	Role    string    `bson:"role" json:"role"`
	DateOTP time.Time `json:"date-otp" bson:"date-otp"`
	OTPCode string    `bson:"otp-code" json:"otp-code"`
}
