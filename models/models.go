package models

import (
	"time"

	"github.com/vaibhavchalse99/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName  string             `json:"fullName,omitempty" bson:"fullName,omitempty"`
	Email     string             `json:"email,omitempty"`
	Password  string             `json:"password,omitempty"`
	Role      db.RoleValue       `json:"role,omitempty"`
	Token     string             `json:"token,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty"`
}

type Book struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Author    string             `json:"author,omitempty"`
	Price     string             `json:"price,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty"`
}

type Status string

const (
	Available    Status = "Available"
	Issued       Status = "Issued"
	NotAvailable Status = "Not available"
)

type BookCopies struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BookId primitive.ObjectID `json:"bookId,omitempty"`
	Status Status             `json:"status,omitempty"`
}

type History struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId     primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	BookCopyId primitive.ObjectID `json:"bookCopyId,omitempty" bson:"bookCopyId,omitempt"`
	BookId     primitive.ObjectID `json:"bookId,omitempty" bson:"bookId,omitempt"`
	IssuedAt   time.Time          `json:"issuedAt,omitempty" bson:"issuedAt,omitempt"`
	ReturnedAt time.Time          `json:"returnedAt,omitempty" bson:"returnedAt,omitempt"`
}
