package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Debug(obj any) {
	raw, _ := json.MarshalIndent(obj, "", " \t")
	fmt.Println(string(raw))
}

// LocalTime is a function to get local time
func LocalTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

// ConvertStringTimeToTime is a function to convert string time to time
func ConvertStringTimeToTime(t string) time.Time {
	// thai time layout
	// layout := "2006-01-02 15:04:05.999999999 -0700 MST"
	layout := "2006-01-02 15:04:05.999 -0700 MST"
	result, err := time.Parse(layout, t)
	if err != nil {
		log.Printf("Error: Parse time failed: %v", err)
	}

	return result
}

