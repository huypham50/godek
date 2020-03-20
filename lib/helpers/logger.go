package helpers

import (
	"log"
	"reflect"
	"time"
)

// LoggerLine -> log all elements in order no matter types
func LoggerLine(items ...interface{}) {
	log.Println("_____", time.Now(), "_____")
	for _, item := range items {
		log.Println(item, "--", reflect.ValueOf(item).Kind())
	}
	log.Println("____________________")
}
