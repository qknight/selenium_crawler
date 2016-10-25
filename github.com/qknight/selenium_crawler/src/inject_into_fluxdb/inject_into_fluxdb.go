package main

import (
	"encoding/json"
	// 	"fmt"
	"github.com/fatih/structs"
	"github.com/influxdata/influxdb/client/v2"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

type rec struct {
	Id     int
	Votes  int
	Verein string
	Title  string
}

type json_message struct {
	Date time.Time
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
	//fmt.Println(f)

	var l json_message
	err2 = json.Unmarshal(f, &l)
	if err2 != nil {
		log.Fatalln("Error: ", err2)
		return
	}
	//fmt.Println("--", l)

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

	for _, r := range l.Data {
		pt, err := client.NewPoint("frei", map[string]string{"Id": strconv.Itoa(r.Id)}, structs.Map(r), time.Now())
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
