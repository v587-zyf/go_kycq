package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	_ "net/http/pprof"
	"regexp"
	"runtime"
	"time"

	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	_ "github.com/go-sql-driver/mysql"
)

// 玩家队伍的名字

const (
	APP_NAME    = "yulong_genconfig"
	APP_VERSION = "0.5.0"
)

var (
	cntjs          = flag.String("cntjs", "./cn-t.js", "set cn-t.js filepath ")    //繁体简体的库文件
	lang           = flag.String("lang", "zh-cn", "set language: zh-cn or zh-tw ") //需要生成的语言版本
	showVersion    = flag.Bool("version", false, "print version string")
	showHelp       = flag.Bool("help", false, "show help")
	client         = flag.Bool("client", true, "make client data file")
	gameDbBasePath = flag.String("gamedb", "../../../ylserver/data/configs", "specify gamedb file")
	tplPath        = flag.String("tplPath", "./", "specify templates file")
	baseGroupName  = flag.String("baseGroupName", "data", "specify base group name when export multiple groups")
	reducedObj     = flag.Bool("reducedObj", false, "generate with reduced object, with fields marshaled to json string when level > 2")
)

func version() string {
	return fmt.Sprintf("%s v%s (built w/%s)", APP_NAME, APP_VERSION, runtime.Version())
}

func main() {
	logger.Init()
	fmt.Println("begin")
	now := time.Now()

	defer func() {
		fmt.Println("\n\n服务器游戏配置打包耗时：", time.Since(now).Seconds())
	}()

	flag.Parse()
	if *showHelp {
		flag.Usage()
		return
	}
	if *showVersion {
		fmt.Println(version())
		return
	}

	//语言版本设置
	if *lang != "zh-cn" {
		setLanguageSet(*lang, *cntjs) //设置语言版本
	}

	err := gamedb.Load(*gameDbBasePath)
	if err != nil {
		fmt.Println("加载配置表错误：\n", err)
		return
	}

	//if *client{
	//	err = DoExport(*tplPath, gamedb.GetDb(), *baseGroupName, *reducedObj)
	//}
	//
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
}

//设置语言版本
func setLanguageSet(lang, cntjs string) bool {
	//从js文件中读取简体繁体子集
	data, err := ioutil.ReadFile(cntjs)
	if err != nil {
		fmt.Println("cntjs read file error:", err)
		return false
	}
	str := string(data)

	reg1, err := regexp.Compile(`cn_s *= *"(.*)" *;`)
	if err != nil {
		fmt.Println("cntjs regexp1 error:", err)
		return false
	}
	reg2, err := regexp.Compile(`cn_t *= *"(.*)" *;`)
	if err != nil {
		fmt.Println("cntjs regexp2 error:", err)
		return false
	}

	result1 := reg1.FindStringSubmatch(str)
	result2 := reg2.FindStringSubmatch(str)

	if len(result1) != 2 || len(result2) != 2 {
		fmt.Println("cntjs regexp find string error")
		return false
	}

	if len(result1[1]) != len(result2[1]) {
		fmt.Println("cntjs cn.length != tw.length")
		return false
	}

	gamedb.SetLang(lang)

	gamedb.SetLanguageSet(result1[1], result2[1])

	return true

}
