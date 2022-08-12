package helper

import (
	"fmt"
	"reflect"
)

func Unbundle(item interface{}) {
  r := reflect.ValueOf(item).Elem()

  itemMap := make(map[string]interface{})

  for i := 0; i < r.NumField(); i++ {
    field := r.Type().Field(i)

    itemMap[field.Name] = r.Field(i).Interface()
  }
  fmt.Println(itemMap)
}
