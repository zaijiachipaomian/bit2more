package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// 采集 尤果网站的 陈一的图片。。。 练习使用
// 其他的数据不做过多的分心，

var (
	//https://www.ugirls.com/Shop/Detail/Product-467.html
	//https://www.ugirls.com/Shop/Detail/Product-505.html
	url = `https://www.ugirls.com/Shop/Detail/Product-467.html`

	srcReg = regexp.MustCompile(`<img src="([^"]+)"`)
)
func main(){

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("err",err.Error())
		return
	}

	defer resp.Body.Close()

	//fmt.Printf("%+v\n",resp)
	/*
		{Status:200 OK
	StatusCode:200 Proto:HTTP/2.0 ProtoMajor:2 ProtoMinor:0
	Header:map[Server:[Tengine] Content-Type:[text/html; charset=utf-8] Vary:[Accept-Encoding]
	Set-Cookie:[PHPSESSID=bp1edkl3o7gtk39bfddgfdd9n4; path=/] X-Traceid:[14186761cc96899ef0b9eadcc14786a2]
	X-S-Tags:[49169] Access-Control-Allow-Origin:[*] Date:[Mon, 08 Oct 2018 11:42:54 GMT] X-Resp-Time:[0.113]
	Access-Control-Allow-Methods:[GET,POST]] Body:0xc0000710e0 ContentLength:-1 TransferEncoding:[] Close:false Uncompressed:true
	Trailer:map[] Request:0xc0000e8000 TLS:0xc00043afd0}
	 */

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {

	}

	//fmt.Printf("%s \n",string(data))

	submatch := srcReg.FindAllStringSubmatch(string(data), -1)


	for _, v:=range  submatch {
		//fmt.Println(v[1])
		if ! strings.ContainsAny(v[1],".png") || ! strings.ContainsAny(v[1],".jpg"){
			fmt.Println("back")
			continue
		}
		go 	func (url string){
			resp, err := http.Get(url)

			if err != nil {
				fmt.Printf("fetch %s error %s\n",url , err)
				return
			}

			defer resp.Body.Close()

			data,err:= ioutil.ReadAll(resp.Body)

			if err != nil {
				return
			}
			split := strings.Split(url, "/")
			file, err := os.OpenFile("./image/hezhouzhou/"+split[len(split)-1],  os.O_CREATE|os.O_RDWR,0766)

			if err != nil {
				fmt.Println("error ",err )
				return
			}

			defer file.Close()

			file.Write(data)

		}(v[1])
	}

//	fmt.Printf("%+v\n",submatch)

	time.Sleep(time.Minute * 3 )
}
