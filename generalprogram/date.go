package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var v interface{}
	v = "2008-01-22T15:04:05.000"

	fmt.Println(changeTimestamp(v.(string)))

}

func changeTimestamp(v interface{}) string {
	format := "2006-01-02T15:04:05.000"

	ct, err := time.Parse(format, v.(string))
	if err != nil {
		log.Printf("error is parsing the timestamp format")
	}
	return (ct.Format("2006-01-02 15:04:05.000"))

}
