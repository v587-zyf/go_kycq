package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		//Usage()
		//return
	}
	fmt.Println("以下包将会被改变版本号: ")
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
	fmt.Printf("请输入版本号: ")
	var ver string
	fmt.Scanln(&ver)
	if len(ver) > 0 {
		fmt.Printf("输入的新版本号为 < %s > , 确认开始改变打包的版本号", ver)
	} else {
		fmt.Printf("没有输入新版本号，将用老的版本号打包，确认开始打包")
	}
	fmt.Scanln()
	if len(ver) == 0 {
		return
	}
	rand.Seed(time.Now().UnixNano())
	r := 10000+rand.Intn(10000)
	ver = ver + "."+strconv.Itoa(r)
	newVer := "APP_VERSION = " + "\"" + ver + "\"" + "\n"
	for i := 1; i < len(os.Args); i++ {
		ChageVersion(os.Args[i], newVer)
	}
	file,err:=os.Create("server_game_update.sql")
	if err!=nil{
		fmt.Println(err)
	}
	content:=[]byte("UPDATE `system_info` SET `content`='"+strconv.Itoa(r)+"' WHERE id=9;")
	err =ioutil.WriteFile("server_game_update.sql",content,0777)
	if err!=nil {
		fmt.Println(err)
	}
	defer file.Close()
	return
}

func Usage() {
	fmt.Printf("需要包名")
	fmt.Scanln()
}

func ChageVersion(p string, version string) {
	file := p + string(os.PathSeparator) + "main.go"
	fmt.Printf("改变版本号 %s %s\n", file, version)
	if contents, err := ioutil.ReadFile(file); err != nil {
		fmt.Printf("改变版本号读文件失败 %s %s\n", file, err)
	} else {
		//替换到行尾
		re, e := regexp.Compile("APP_VERSION = .*\n")
		if e != nil {
			fmt.Printf("改变版本号失败 %s %s\n", file, e)
			return
		}
		result := re.ReplaceAllString(string(contents), version)
		e = ioutil.WriteFile(file, []byte(result), 0644)
		if e != nil {
			fmt.Printf("改变版本号失败 %s %s\n", file, e)
			return
		}
		fmt.Printf("改变版本号成功 %s %s\n", file, version)
	}
}
