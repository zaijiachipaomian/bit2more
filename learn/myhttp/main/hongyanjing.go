package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//http://music.163.com/song/media/outer/url?id=214333.mp3
func main(){

	resp, err := http.Get("http://music.163.com/song/media/outer/url?id=214333.mp3")

	if err != nil{
		fmt.Println("err",err)
		return
	}


	fmt.Printf("%+v \n",resp)

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)


	fmt.Printf("%s\n",string(data))

	file, err := os.OpenFile("./mp3/coolsong.mp3", os.O_RDWR|os.O_CREATE, 0766)

	if err != nil {
		fmt.Printf("error ",err)
		return
	}

	defer file.Close()


	file.Write(data)
}

