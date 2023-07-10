package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	//resp, err := http.PostForm("http://localhost:8500/pay/hgamePay",
	//	url.Values{"game_orderno": {"2sssa"}, "id": {"123"}})
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i
	}

	arrNew := arr[0:3]
	fmt.Println("len:", len(arrNew))
	for i := 0; i < len(arrNew); i++ {
		fmt.Printf("%v = %v\n", i, arrNew[i])
	}

	var a, b int64
	a = 3
	b = 4
	fmt.Println("a*b =", a*b)

	// resp, err := http.PostForm("http://dev-dot-xkx2-wx.hgame.com/DotLog/dot",
	// 	url.Values{"uuid": {"uuid"}, "uName": {"aaa"}, "roleId": {"23232322"}, "serverID": {"5"}, "platformID": {"wx"}, "game": {"xkx2"}, "key": {"555"}, "info": {"info22"}})

	_, err := http.PostForm("http://192.168.5.222:8500/api/removeItem",
		url.Values{"openId": {"g1"}, "serverId":{"1"},"itemId":{"5032"}, "itemCount":{"1"}})

	if err != nil {
		// handle error
	}

	if err != nil {
		// handle error
	}

}
