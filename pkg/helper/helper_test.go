package helper

import (
	"live-resume/pkg/models"
	"testing"
)

func TestMain(t *testing.T) {
  item := models.Basic{
    Firstname: "A",
  }

  Unbundle(&item)
}
