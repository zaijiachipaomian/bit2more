package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	url =flag.String("url","","load url")
	fileName = flag.String("file","","save name")

	path = flag.String("path","","save where")

)

const (
	C_URL= "http://music.163.com/song/media/outer/url?id=167876.mp3"
	C_NAME =`有何不可`
	C_PATH = `./mp3`
)
func main(){

	flag.Parse()
	if *url == ""{
		*url = C_URL
	}

	if *fileName ==""{
		*fileName=C_NAME
	}

	if *path ==""{
		*path=C_PATH
	}

	resp, err := http.Get(*url)

	if err != nil {
		fmt.Printf("400  %+v \n",err)
		return
	}

	defer resp.Body.Close()

	file, err := os.OpenFile(*path+"/"+*fileName+".mp3", os.O_RDWR|os.O_CREATE, 0766)

	if err != nil {
		fmt.Printf("401 file  %+v \n",err)
		return
	}

	defer file.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("%d ",resp.StatusCode)
		return
	}

	_,err = io.Copy(file,resp.Body)
	if err != nil {
		fmt.Printf("405 %v\n",err)
		return
	}
	fmt.Println("done")
}

