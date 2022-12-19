package repository

import (
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateDocFromStruct creates bson.D for update from tags of struct. This call may panic. Always send a struct, Im not sure what may happens
func CreateDocFromStruct(st interface{}) bson.D {
	doc := bson.D{}
	t := reflect.TypeOf(st)
	v := reflect.ValueOf(st)
	kind := t.Kind()
	if kind == reflect.Ptr {
		t = t.Elem()
		kind = t.Kind()
		v = v.Elem()
	}
	if kind != reflect.Struct {
		log.Errorf("parameter should be struct, instead is [%v] name [%s]", kind, t.Name())
		return doc
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		key := field.Tag.Get("bson")
		tags := strings.Split(key, ",")
		if len(tags) > 0 {
			if contains(tags, "-") || contains(tags, "_id") {
				continue
			}
			if contains(tags, "omitempty") && isZeroOfUnderlyingType(v.Field(i).Interface()) {
				continue
			}
			if len(tags) > 0 {
				key = tags[0]
			} else {
				key = field.Name
			}
			if len(key) == 0 {
				key = field.Name
			}

		} else {
			switch key {
			case "-", "_id":
				continue
			case "":
				key = field.Name
			}

		}
		doc = append(doc, bson.E{Key: key, Value: v.Field(i).Interface()})
	}
	return doc
}
func contains(where []string, what string) bool {
	for _, w := range where {
		if w == what {
			return true
		}
	}
	return false
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
