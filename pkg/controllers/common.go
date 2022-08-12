package controllers

import (
	"context"
	"fmt"
	"net/http"

  "github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ------\\
// ------ ||
// COMMON ||
// ------ ||
// ------//

// Validate data
var validate = validator.New()

// Open collection for operations
var userCollection *mongo.Collection = OpenCollection(Client, "users")

func obtainID(c *gin.Context) primitive.ObjectID {
  // UserID
  var userID = c.Params.ByName("id")
  var docID, _ = primitive.ObjectIDFromHex(userID)

  return docID
}

func findContent(c *gin.Context, ctx context.Context, item bson.M, opt *options.FindOneOptions) {
  // Get ID
  docID := obtainID(c)

  if err := userCollection.FindOne(ctx, bson.M{"_id": docID}, opt).Decode(&item);
  err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    fmt.Println(err)
    return
  }

  c.JSON(http.StatusOK, item)
}
