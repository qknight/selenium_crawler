package main

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/influxdata/influxdb/client/v2"
	"io/ioutil"
	"log"
	"strconv"
	"time"
  "fmt"
//   "os"
)

type rec struct {
	Id     int
	Votes  int
	Verein string
	Title  string
}

type json_message struct {
	Date string
	Data []rec
}

const (
	MyDB     = "square_holes"
	username = "bubba"
	password = "bumblebeetuna"
)

func main() {

	f, err2 := ioutil.ReadFile("data.json")
	if err2 != nil {
		log.Fatalln("Error: ", err2)
		return
	}

	var l json_message
	err2 = json.Unmarshal(f, &l)
	if err2 != nil {
		log.Fatalln("Error: ", err2)
		return
	}

	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}
	
	layout := "2006-01-02T15:04:05.000Z"
  t, err3 := time.Parse(layout, l.Date)

  if err3 != nil {
      fmt.Println(err)
      return
  }
// 	fmt.Println("#### ", time.Now(), " time.Now()")
//   fmt.Println("|||| ", t.Local().String())
//   fmt.Println("&&&& ", t.String())
//      
//   os.Exit(1)
	
	for _, r := range l.Data {
		pt, err := client.NewPoint("frei", map[string]string{"Id": strconv.Itoa(r.Id)}, structs.Map(r), t.Local())
		bp.AddPoint(pt)
		if err != nil {
			log.Fatalln("Error: ", err)
		}
	}

	//     // Create a point and add to batch
	//     tags := map[string]string{"cpu": "cpu-total"}
	//     fields := map[string]interface{}{
	//         "idle":   10.1,
	//         "system": 53.3,
	//         "user":   46.6,
	//     }
	//     pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	/*
	   if err != nil {
	       log.Fatalln("Error: ", err)
	   }

	   bp.AddPoint(pt)*/

	// Write the batch
	c.Write(bp)
}
