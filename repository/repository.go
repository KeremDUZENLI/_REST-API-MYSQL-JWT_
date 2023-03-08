package repository

import (
	"context"
	"net/http"
	"strconv"

	"jwt-project/common/constants"
	"jwt-project/database"
	"jwt-project/database/model"
	"jwt-project/middleware/token"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	encryptionSize := 14
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), encryptionSize)
	return string(bytes)
}

func Update(ctx context.Context, foundPerson model.Person) {
	firstToken, refreshToken := token.GenerateToken(*foundPerson.Email, *foundPerson.FirstName, *foundPerson.LastName, *foundPerson.UserType, foundPerson.UserId)
	token.UpdateAllTokens(firstToken, refreshToken, foundPerson.UserId)
	database.Collection(database.Connect(), constants.TABLE).FindOne(ctx, bson.M{"userid": foundPerson.UserId}).Decode(&foundPerson)
}

func InsertNumberInDatabase(c *gin.Context, ctx context.Context, person model.Person) *mongo.InsertOneResult {
	resultInsertionNumber, _ := database.Collection(database.Connect(), constants.TABLE).InsertOne(ctx, person)
	return resultInsertionNumber
}

func Stages(c *gin.Context) (primitive.D, primitive.D, primitive.D) {
	recordPerPage, errorConvertionRecord := strconv.Atoi(c.Query("recordPerPage"))
	if errorConvertionRecord != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, errorConvertionPage := strconv.Atoi(c.Query("page"))
	if errorConvertionPage != nil || page < 1 {
		page = 1
	}

	startIndex, errorConvertionStartIndex := strconv.Atoi(c.Query("startIndex"))
	if errorConvertionStartIndex != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Provide a valid integer start number"})
	}

	matchStage := bson.D{{"$match", bson.D{{}}}}

	groupStage := bson.D{{"$group", bson.D{
		{"_id", bson.D{{"_id", "null"}}},
		{"total_count", bson.D{{"$sum", 1}}},
		{"data", bson.D{{"$push", "$$ROOT"}}}}}}

	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"total_count", 1},
			{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}}}}}

	return matchStage, groupStage, projectStage
}

func Results(c *gin.Context, ctx context.Context) *mongo.Cursor {
	matchStage, groupStage, projectStage := Stages(c)
	result, _ := database.Collection(database.Connect(), constants.TABLE).Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage,
	})

	return result
}