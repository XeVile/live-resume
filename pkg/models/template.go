package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Basic struct {
  Firstname      string              `json:"firstname,omitempty"`
  Lastname       string              `json:"lastname,omitempty"`
  Email          string              `json:"email,omitempty"`
  Phone          string              `json:"phone,omitempty"`
  SecPhone       string              `json:"secphone,omitempty"`
  Address        string              `json:"address,omitempty"`
}

type Education struct {
	Title           string             `json:"title,omitempty"`
	Description     string             `json:"description,omitempty"`
	InstitutionName string             `json:"institutionName,omitempty"`
	StartDate       string             `json:"startDate,omitempty"`
	EndDate         string             `json:"endDate,omitempty"`
}

type List struct {
	Name    string             `json:"name,omitempty"`
	Content string             `json:"content,omitempty"`
}

type Skill struct {
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Proficiency string             `json:"proficiency,omitempty"`
}

type Job struct {
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	CompanyName string             `json:"companyName,omitempty"`
	Location    string             `json:"location,omitempty"`
	StartDate   string             `json:"startDate,omitempty"`
	EndDate     string             `json:"endDate,omitempty"`
}

type Project struct {
  Title       string             `json:"title,omitempty"`
  Description string             `json:"description,omitempty"`
	AddDetails  string             `json:"addDetails,omitempty"`
	URL         string             `json:"url,omitempty"`
}

type UserData struct {
  ID             primitive.ObjectID `bson:"_id"`
  Basic          *Basic              `json:"basic,omitempty"`
	Education      *[]Education        `json:"education,omitempty"`
	Skill          *[]Skill            `json:"skill,omitempty"`
	Job            *[]Job              `json:"job,omitempty"`
  Project        *[]Project          `json:"project,omitempty"`
	List           *[]List             `json:"list,omitempty"`
}

func fill(sourcePath string, templatePath string, outputFilename string) (string, error) {
  // Read from path
  src, readErr := ioutil.ReadFile(sourcePath)
  ud := UserData{}

  if readErr != nil {
    panic(readErr)
  } else {
    err := json.Unmarshal([]byte(src), &ud)
    
    if err != nil {
      panic(err)
    }
  }

  t := template.Must(template.ParseFiles(templatePath))

  outputDir := filepath.Dir(templatePath)
  outputPath := outputDir + "/" + outputFilename

  resume, err := os.Create(outputPath)

  if err != nil {
    panic(err)
  } else {
    t.Execute(resume, ud)
    fmt.Printf("Template Filled -> %s.html\n", outputFilename)
    return outputPath, nil
  }
}
