package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type myServeHTTP struct {
}


// 在handler 这里做 中间件的处理。。。。
func (this myServeHTTP)ServeHTTP(w http.ResponseWriter,r *http. Request){

	fmt.Printf("%+v\n",r.URL.Path)

	if r.URL.Path == "/favicon.ico" || r.URL.Path == "/hello"{

		var message = struct {
			Detail string `json:"for"`
		}{
			Detail:"maybe our server is turn some problems ... you can ",
		}

		data,_ := json.Marshal(message)
		w.Write(data)
		w.WriteHeader(200)
	}


}

var (
	myhandler = myServeHTTP{}
)
func main(){

	server := http.Server{
		Addr:    ":8888",
		Handler: myhandler,

		ReadHeaderTimeout: time.Second * 10 ,
		WriteTimeout: time.Second * 10 ,
		MaxHeaderBytes: 1 << 20 ,
	}


	log.Fatal(server.ListenAndServe())
}
