package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"live-resume/pkg/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ########################
// ------------------------
// C.R.U.D functions      |
// ------------------------
// ########################

// AddUser ... Create an user
func AddUser(c *gin.Context) {
  // Set timeout of ctx to 100 sec
  var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

  defer cancel()

  // Create User instance
  var user models.UserData
  //var basic models.Basic
  //var education models.Education
  //var skill models.Skill
  //var job models.Job
  //var project models.Project
  //var list models.List

  if err := c.BindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    fmt.Println(err)
    return
  }
   
  // Validate
  validErr := validate.Struct(user)
  if validErr != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": validErr.Error()})
    fmt.Println(validErr)
    return
  }

  // Insert new object with ID
  user.ID =  primitive.NewObjectID()
  result, insertErr := userCollection.InsertOne(ctx, user)
  if insertErr != nil {
    msg := "user item was not created"
    c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
    fmt.Println(insertErr)
    return
  }

  c.JSON(http.StatusOK, result)
}


// GetUser ... Read an user
func GetUser(c *gin.Context) {
  // Set timeout of ctx to 100 sec
  var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

  defer cancel()

  // Get ID
  docID := obtainID(c)

  // Create User instance
  var user bson.M

  if err := userCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&user);
  err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    fmt.Println(err)
    return
  }

  c.JSON(http.StatusOK, user)
}

// UpdateUser ... Update an user
func UpdateUser(c *gin.Context) {
  // Set timeout of ctx to 100 sec
  var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

  defer cancel()

  // Get ID
  docID := obtainID(c)

  // Get user struct
  var user models.UserData

  // Binding the pointed variable
  if err := c.BindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    fmt.Println(err)
    return
  }

  // Validate
  validErr := validate.Struct(user)
  if validErr != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": validErr.Error()})
    fmt.Println(validErr)
    return
  }

  update := bson.M{"basic": user.Basic,
                   "education": user.Education,
                   "skill": user.Skill,
                   "job": user.Job,
                   "project": user.Project,
                   "list": user.List,}

  result, err := userCollection.ReplaceOne(ctx, bson.M{"_id": docID}, update)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, result.ModifiedCount)
}

// DeleteUser ... Update an user
func DeleteUser(c *gin.Context) {
  // Set timeout of ctx to 100 sec
  var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

  defer cancel()

  docID := obtainID(c)

  // Use docID primitive to filter and delte item
  result, err := userCollection.DeleteOne(ctx, bson.M{"_id": docID})
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, result.DeletedCount)
}
