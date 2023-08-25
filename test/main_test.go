package test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test_marshal(t *testing.T) {
	jsonStr := "[{\"container_name\": \"test\"}]"
	var containers []Container
	err := UnmarshalWithCustomTag(jsonStr, &containers, "myTag")
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return
	}
	fmt.Printf("containers: %+v\n", containers)
}

type Container struct {
	ContainerName string `myTag:"container_name"`
	ImageVersion  string `myTag:"image_version"`
	Publisher     string `myTag:"publisher"`
	PublishTime   string `myTag:"publish_time"`
}

func UnmarshalWithCustomTag(jsonStr string, v interface{}, tagName string) error {
	var data []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return err
	}

	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.Len(); i++ {
		t := val.Index(i).Type()
		for j := 0; j < t.NumField(); j++ {
			field := t.Field(j)
			fieldVal := val.Index(i).Field(j)
			key := field.Tag.Get(tagName)
			if key != "" {
				value, ok := data[i][key]
				if !ok {
					continue
				}
				switch field.Type.Kind() {
				case reflect.String:
					fieldVal.SetString(value.(string))
				case reflect.Int:
					num := int(value.(float64))
					fieldVal.SetInt(int64(num))
				}
			}
		}
	}
	return nil
}
