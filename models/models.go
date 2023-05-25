package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       *string            `json:"last_name" validate:"required,min=2,max=30"`
	Password        *string            `json:"password" validate:"required,min=8"`
	Email           *string            `json:"email" validate:"email,required"`
	Phone           *string            `json:"phone" validate:"required"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         *string            `json:"user_id"`
	User_Cart       []ProductUser      `json:"usercart" bson:"usercart"`
	Address_Details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"order_status" bson:"order_status"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Image        *string            `json:"image"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Image        *string            `json:"image" bson:"image"`
	Price        *uint64            `json:"price" bson:"price"`
	Rating       *uint8             `json:"rating" json:"rating"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"_id"`
	Flat       *uint64            `json:"flat_name" bson:"house_name"`
	House      *uint64            `json:"house_name" bson:"house_name"`
	Hood       *string            `json:"hood_name" bson:"hood_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Country    *string            `json:"country_name" bson:"country_name"`
	Postcode   *uint64            `json:"post_code" bson:"post_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_list" json:"order_list"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          *uint64            `json:"total_price" bson:"total_price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
