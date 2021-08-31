package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Factory struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	SrNo        string             `bson:"srno"`
	FactoryName string             `bson:"factoryname"`
	ConPerson   string             `bson:"con_person"`
	PerDesn     string             `bson:"per_desn"`
	Mobile      int                `bson:"mobile"`
	Email       string             `bson:"email" validate:"email"`
	City        string             `bson:"city"`
	Pin         int                `bson:"pin"`
	UName       string             `bson:"uname" validate:"alphanum,required,gte=5,lte=20"`
	Password    string             `bson:"password" validate:"alphanum,required,gte=8,lte=20"`
	ConfirmPass string             `bson:"confirm_pass" validate:"alphanum,required,eqfield=Password"`
	RegiTime    primitive.DateTime `bson:"regitime"`
}
