package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	gcontext "github.com/gorilla/context"
	"github.com/vaibhavchalse99/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://vaibhavchalse:Pass%40123%40123@cluster0.rsjlcxc.mongodb.net/?retryWrites=true&w=majority"
const dbName = "library-management"
const userCol = "user"
const bookCol = "book"
const bookCopiesCol = "book-copies"
const recordHistoryCol = "record-history"

var userCollection *mongo.Collection
var bookCollection *mongo.Collection
var bookCopiesCollection *mongo.Collection
var recordHistoryCollection *mongo.Collection

type ErrMsg struct {
	Message string `json:"message"`
}

type ResData struct {
	Data interface{} `json:"data"`
}

type SuccessMsg struct {
	Message string `json:"message"`
}

type BookReqBody struct {
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Price  string `json:"price,omitempty"`
	Count  int    `json:"count,omitempty"`
}

type BookWithCopies struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Author    string             `json:"author,omitempty"`
	Price     string             `json:"price,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty"`
	// AvailableBooks int                `json:"availableBooks"`
	Copies []models.BookCopies `json:"copies,omitempty"`
}

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	userCollection = client.Database(dbName).Collection(userCol)
	bookCollection = client.Database(dbName).Collection(bookCol)
	bookCopiesCollection = client.Database(dbName).Collection(bookCopiesCol)
	recordHistoryCollection = client.Database(dbName).Collection(recordHistoryCol)
}

func createToken(userId string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	var sampleSecretKey = []byte(os.Getenv("SECRET_KEY"))

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicaion/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user models.User
	var dbUser models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := userCollection.FindOne(context.Background(), bson.D{{Key: "email", Value: user.Email}}).Decode(&dbUser)

	if err != nil && err != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(ErrMsg{Message: "something went wrong"})
	}
	if dbUser.Email != "" {
		json.NewEncoder(w).Encode(ErrMsg{Message: "user already exists"})
		return
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = "END_USER"
	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(SuccessMsg{Message: "User added successfully"})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicaion/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	cursor, err := userCollection.Find(context.Background(), bson.D{{}}, options.Find().SetProjection(bson.M{"password": 0}))
	if err != nil {
		log.Fatal(err)
	}
	var userList []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		userList = append(userList, user)
	}
	defer cursor.Close(context.Background())
	json.NewEncoder(w).Encode(userList)
}

func LoginAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicaion/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	filter := bson.D{{
		Key: "$and", Value: bson.A{
			bson.D{{Key: "email", Value: user.Email}},
			bson.D{{Key: "password", Value: user.Password}},
		},
	}}
	var result models.User
	err := userCollection.FindOne(context.Background(), filter, options.FindOne().SetProjection(bson.M{"password": 0})).Decode(&result)
	if err != nil {
		// log.Fatal(err)
		json.NewEncoder(w).Encode(ErrMsg{Message: "Please send proper data"})
		return
	}
	token, err := createToken(result.Id.Hex())
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	result.Token = token
	json.NewEncoder(w).Encode(result)
}

func getAuthTokenFromHeader(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	arr := strings.Split(bearerToken, " ")
	if len(arr) == 2 {
		return arr[1]
	}
	return ""
}

func verifyJwt(bearerToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid Token")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func IsLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenValue := getAuthTokenFromHeader(r)
		token, err := verifyJwt(tokenValue)
		if err != nil {
			json.NewEncoder(w).Encode(ErrMsg{Message: err.Error()})
			return
		}
		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				var user models.User
				claimUserId := claims["userId"]
				id, _ := primitive.ObjectIDFromHex(fmt.Sprint(claimUserId))
				err := userCollection.FindOne(context.Background(), bson.M{"_id": id}, options.FindOne().SetProjection(bson.M{"password": 0, "token": 0})).Decode(&user)
				if err != nil {
					json.NewEncoder(w).Encode(ErrMsg{Message: "something went wrong"})
					return
				}
				gcontext.Set(r, "user", user)
				next(w, r)
			}
		} else {
			json.NewEncoder(w).Encode(ErrMsg{Message: "Invalid token"})
			return
		}
	})
}

func GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	user := gcontext.Get(r, "user")
	json.NewEncoder(w).Encode(user)
}

func UpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
	user, ok := gcontext.Get(r, "user").(models.User)
	if ok {
		var newUser models.User
		_ = json.NewDecoder(r.Body).Decode(&newUser)

		if newUser.FullName == "" || newUser.Password == "" {
			json.NewEncoder(w).Encode(ErrMsg{Message: "Invalid Data"})
			return
		}

		// userEmail, _ := primitive.ObjectIDFromHex(user.Email)

		filter := bson.M{"email": user.Email}
		update := bson.M{"$set": bson.M{"fullName": newUser.FullName, "password": newUser.Password}}

		_, err := userCollection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			json.NewEncoder(w).Encode(ErrMsg{Message: "something went wrong"})
			return
		}
		json.NewEncoder(w).Encode(SuccessMsg{Message: "Data updadated successfully"})
		return
	}
	json.NewEncoder(w).Encode(ErrMsg{Message: "something went wrong"})
}

func GetBookList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicaion/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	queryStatus := r.URL.Query().Get("status")

	var cursor *mongo.Cursor
	var err error
	if queryStatus != "" {
		lookupStage := bson.D{{Key: "$lookup", Value: bson.M{"from": "book-copies", "localField": "_id", "foreignField": "bookId", "as": "copies"}}}
		filter := bson.D{{Key: "$match", Value: bson.M{"copies.status": queryStatus}}}
		// projectStage := bson.D{{Key: "$project", Value: bson.M{"name": 1, "author": 1, "price": 1, "availableBooks": bson.D{{Key: "$size", Value: "$copies"}}}}}
		cursor, err = bookCollection.Aggregate(context.Background(), mongo.Pipeline{lookupStage, filter})
	} else {
		cursor, err = bookCollection.Find(context.Background(), bson.D{{}})
	}
	if err != nil {
		json.NewEncoder(w).Encode(ErrMsg{Message: "something went wrong"})
		return
	}
	var bookList []BookWithCopies
	for cursor.Next(context.Background()) {
		var book BookWithCopies
		if err := cursor.Decode(&book); err != nil {
			json.NewEncoder(w).Encode(ErrMsg{Message: "something went wrong"})
			return
		}
		bookList = append(bookList, book)
	}
	defer cursor.Close(context.Background())
	json.NewEncoder(w).Encode(ResData{Data: bookList})
}

func createCopiesData(count int, bookId interface{}) []interface{} {
	primitiveBookId := bookId.(primitive.ObjectID)
	var booklist []interface{}

	for i := 0; i < count; i++ {
		var copy = bson.D{
			{Key: "status", Value: "Available"},
			{Key: "bookId", Value: primitiveBookId},
		}
		booklist = append(booklist, copy)
	}
	return booklist
}

func Addbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicaion/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var newBook, dbBook models.Book
	var book BookReqBody

	_ = json.NewDecoder(r.Body).Decode(&book)
	if book.Count == 0 {
		json.NewEncoder(w).Encode(SuccessMsg{Message: "Please enter a valid count"})
		return
	}

	err := bookCollection.FindOne(context.Background(), bson.D{{Key: "name", Value: book.Name}}).Decode(&dbBook)
	if err != nil && err != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(ErrMsg{Message: "something went wrong"})
		return
	}
	if dbBook.Name != "" {
		json.NewEncoder(w).Encode(ErrMsg{Message: "book already present"})
		return
	}

	newBook.Name = book.Name
	newBook.Author = book.Author
	newBook.Price = book.Price
	newBook.CreatedAt = time.Now()
	newBook.UpdatedAt = time.Now()

	insertedBook, err := bookCollection.InsertOne(context.Background(), newBook)
	if err != nil {
		json.NewEncoder(w).Encode(ErrMsg{Message: "something went wrong"})
		return
	}

	bookCopies := createCopiesData(book.Count, insertedBook.InsertedID)

	_, err = bookCopiesCollection.InsertMany(context.Background(), bookCopies)
	if err != nil {
		json.NewEncoder(w).Encode(ErrMsg{Message: "something went wrong"})
		return
	}
	json.NewEncoder(w).Encode(SuccessMsg{Message: "Data inserted successfully"})
}

type RecordsReq struct {
	UserId     primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	BookId     primitive.ObjectID `json:"bookId,omitempty" bson:"bookId,omitempty"`
	ReturnedAt string             `json:"returnedAt,omitempty" bson:"returnedAt,omitempty"`
}

func isBookCopyAvailable(bookId primitive.ObjectID) (bookRecord models.BookCopies, err error) {
	err = bookCopiesCollection.FindOne(context.Background(), bson.D{{Key: "bookId", Value: bookId}, {Key: "status", Value: "Available"}}).Decode(&bookRecord)
	return
}

func isUserPresent(userId primitive.ObjectID) (userRecord models.User, err error) {
	err = userCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: userId}}).Decode(&userRecord)
	return
}

func isBookAlreadyIssued(userId, bookId primitive.ObjectID) (record models.History, err error) {
	err = recordHistoryCollection.FindOne(context.Background(), bson.D{{Key: "bookId", Value: bookId}, {Key: "userId", Value: userId}, {Key: "issuedAt", Value: bson.M{"$lt": time.Now()}}, {Key: "returnedAt", Value: bson.M{"$gte": time.Now()}}}).Decode(&record)
	return
}

func AssignBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicaion/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var record RecordsReq
	_ = json.NewDecoder(r.Body).Decode(&record)

	returnedDate, err := time.Parse(time.RFC3339, record.ReturnedAt)
	if err != nil {
		json.NewEncoder(w).Encode(ErrMsg{Message: "please enter returned date"})
		return
	}

	bookCopyData, bookErr := isBookCopyAvailable(record.BookId)
	userData, userErr := isUserPresent(record.UserId)
	if bookErr == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(ErrMsg{Message: "Book not found"})
		return
	}
	if userErr == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(ErrMsg{Message: "User not found"})
		return
	}
	if bookErr != nil || userErr != nil {
		json.NewEncoder(w).Encode(ErrMsg{Message: "Something went wrong"})
		return
	}
	bookIssuedData, bookIssuedErr := isBookAlreadyIssued(record.UserId, record.BookId)
	if bookIssuedErr == mongo.ErrNoDocuments {
		fmt.Println("zxv", bookIssuedData)
		var bookRecord models.History
		if bookCopyData.Status != "" {
			bookRecord.UserId = userData.Id
			bookRecord.BookCopyId = bookCopyData.ID
			bookRecord.BookId = bookCopyData.BookId
			bookRecord.ReturnedAt = returnedDate
			bookRecord.IssuedAt = time.Now()
		}

		_, err := recordHistoryCollection.InsertOne(context.Background(), bookRecord)
		if err != nil {
			json.NewEncoder(w).Encode(ErrMsg{Message: "something went wrong"})
			return
		}
		json.NewEncoder(w).Encode(SuccessMsg{Message: "book assigned successfully"})
		return
	}
	json.NewEncoder(w).Encode(ErrMsg{Message: "Book isAlready issued"})
}

func SampleFunc(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hi this is vaibhav")
}
