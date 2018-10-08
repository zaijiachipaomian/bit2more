package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	r = resjson{
		Now: time.Now(),
		Res: "welcome %s ",
	}
)

func main() {

	// 监听 路径 是 /index
	//
	http.HandleFunc(`/index`, func(writer http.ResponseWriter, request *http.Request) {

		//	fmt.Printf("request = %+v \n",request)
		r.Res = fmt.Sprintf(r.Res, request.RemoteAddr)
		data, _ := json.Marshal(r)

		fmt.Print(string(data))
		fmt.Fprint(writer, string(data))

	})


	// 启动服务。。。 如果 有过过滤的服务， handle
	http.ListenAndServe(":8888", nil)

}

type resjson struct {
	Now time.Time `json:"now"`
	Res string    `json:"res"`
}
