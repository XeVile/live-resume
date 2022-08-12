package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"live-resume/pkg/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ########################
// ------------------------
// -------------------------
// C.R.U.D functions
// -------------------------
// ------------------------
// ########################

// GetItem ... Read Item Info
// -----------------------------
// #############################
func GetItem(c *gin.Context) {
  // Set timeout of ctx to 100 sec
  var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

  defer cancel()

  // Check which item requested
  query := c.Request.RequestURI[strings.LastIndex(c.Request.RequestURI, "/") + 1:]

  // Create Item instance
  var item bson.M
  opt := options.FindOne().SetProjection(bson.M{query: 1})

  findContent(c, ctx, item, opt)
}

//opt := options.FindOne().SetProjection(bson.M{
  //  "basic"      : 1,
  //  "education"  : 1,
  //  "skill"      : 1,
  //  "job"        : 1,
  //  "project"    : 1,
  //  "list"       : 1,
  //})



// UpdateBasic ... Update basic info
// ---------------------------------
// #################################
func UpdateBasic(c *gin.Context) {
  // Set timeout of ctx to 100 sec
  var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

  defer cancel()

  // Get ID
  docID := obtainID(c)

  body := c.Request.Body
  fmt.Println(body)

  var basic models.Basic

  // Binding the pointed variable
  if err := c.BindJSON(&basic); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    fmt.Println(err)
    return
  }

  var format []bson.E
  format = append(format, bson.E{Key: "basic", Value: basic.Firstname})
  format = append(format, bson.E{Key: "lastname" , Value: basic.Lastname })
  update := bson.E{"basic", format}
  fmt.Println(update)
  //update := bson.D{{"basic", [6]bson.E{
  //  {Key: "firstname", Value: basic.Firstname},
  //  {Key: "lastname" , Value: basic.Lastname },
  //  {Key: "email"    , Value: basic.Email    },
  //  {Key: "phone"    , Value: basic.Phone    },
  //  {Key: "secphone" , Value: basic.SecPhone },
  //  {Key: "address"  , Value: basic.Address  },
  //}}}

  result, err := userCollection.UpdateOne(ctx, bson.M{"_id": docID}, bson.D{{"$set", update}})
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, result.ModifiedCount)
}

// UpdateJob ... Update job info
// ---------------------------------
// #################################
func UpdateJob(c *gin.Context) {
  // Set timeout of ctx to 100 sec
  var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

  defer cancel()

  // Get ID
  docID := obtainID(c)

  var job models.Job

  // Binding the pointed variable
  if err := c.BindJSON(&job); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    fmt.Println(err)
    return
  }

  update := bson.M{"title"      :  job.Title      ,
                   "description":  job.Description,
                   "companyName":  job.CompanyName,
                   "location"   :  job.Location   ,
                   "startDate"  :  job.StartDate  ,
                   "endDate"    :  job.EndDate    ,}
 
  result, err := userCollection.UpdateOne(ctx, bson.M{"_id": docID}, bson.D{{"$set", update}})
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, result.ModifiedCount)
}

// UpdateSkill ... Update Skill info
// ---------------------------------
// #################################
func UpdateSkill(c *gin.Context) {
  // Set timeout of ctx to 100 sec
  var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

  defer cancel()

  // Get ID
  docID := obtainID(c)

  var skill models.Skill

  // Binding the pointed variable
  if err := c.BindJSON(&skill); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    fmt.Println(err)
    return
  }

  update := bson.M {"name": skill.Name,
                    "description": skill.Description,
                    "proficiency": skill.Proficiency,}
 
  result, err := userCollection.ReplaceOne(ctx, bson.M{"_id": docID}, update)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, result.ModifiedCount)
}
