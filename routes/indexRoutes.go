package routes

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
    // ``` this function diplays all the events 
    // paging it 10/1 page ``` 

    if r.Method == "GET" { 
        w.Write([]byte("this is the event page "))
    }
    if r.Method == "POST" {
        w.WriteHeader(405)
    }
}

