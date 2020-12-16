package routes

import (
	"backEndTest/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Event struct {
    EventID int64 `json:"eventid"`
    EventName string `json:"eventname"`
    EventDesc string `json:"eventdesc"`
    EventDate string `json:"eventdate"`
    EventTime  string `json:"eventtime"`
    EventRegion string `json:"eventregion"`
}

//Index this function handles the diaplay of the event 
func Index(w http.ResponseWriter, r *http.Request) {
   if r.Method == "GET" { 
		w.Header().Set("content-type", "application/json")
        param := r.Header.Get("Authorization")   
        fmt.Println(param)
        result, err := model.GetEventName()
        if err != nil {
            log.Panic(err.Error())
        }
        event := Event{}
        for result.Next(){
            err := result.Scan(&event.EventID, &event.EventName,  &event.EventDate, &event.EventTime, &event.EventRegion)    
            if err != nil {
                log.Panic(err.Error())
            }
            respData,_ := json.Marshal(event)
            w.Write(respData)
        }
     }
    if r.Method == "POST" {
        http.Error(w, "this method is not supported in this route", 405 )
    }
}
//Events handles the posting of event
//this function parses the request body to a struct and send it to the database 
func Events(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST"{
		w.Header().Set("content-type", "application/json")
        var event Event
        reqBody, _ := ioutil.ReadAll(r.Body)
        json.Unmarshal(reqBody, &event)
        stmt, err := model.InsertEvent()
        if err != nil {
            log.Panic(err.Error())
        }
        _,err = stmt.Exec(event.EventName, event.EventDesc, event.EventDate, event.EventTime, event.EventRegion)
        if err != nil {
            log.Panic(err.Error())
        }
       
    }
}

func GetEventByRegion(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET"{
		w.Header().Set("content-type", "application/json")
        vars := r.URL.Query().Get("region")
        result, err := model.SortEventByRegion(vars)
        if err != nil {
            log.Panic(err.Error())
        }
        event := Event{} 
        for result.Next(){
            err :=  result.Scan(&event.EventID, &event.EventName, &event.EventDate, &event.EventTime, &event.EventRegion) 
            if err != nil{
                log.Panic(err.Error())
            }
            respData,_ := json.Marshal(event)
            w.Write(respData)
        }
    }
}

func EventDetails(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET"{
		w.Header().Set("content-type", "application/json")
        param := r.URL.Query().Get("eventid")
        fmt.Println(param)
        result, err := model.GetEventByID(param)
        if err != nil{
            log.Panic(err.Error())
        }
        var event Event
        err = result.Scan(&event.EventID, &event.EventName, &event.EventDesc, &event.EventDate, &event.EventTime, &event.EventRegion)
        if err != nil {
            log.Panic(err.Error())
        }
        respData,_ := json.Marshal(event)
        w.Write(respData)
    }
}
