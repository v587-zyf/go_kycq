package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var logConfig = `
{
    "types": [
        {
            "type": "console"
        },
        {
            "type": "file",
            "filename": "logs/${programName}/login.log",
            "maxlines": 10000,
            "maxsize": 10240000,
            "daily": true,
            "maxdays": 2,
            "rotate": true,
            "category": "login,default"
        },
        {
            "type": "file",
            "filename": "logs/diamond.log",
            "maxlines": 10000,
            "maxsize": 10240000,
            "daily": true,
            "maxdays": 2,
            "rotate": true,
            "category": "diamond,default"
        }
    ],
    "levels": {
        "login": "DEBUG",
        "diamond": "DEBUG"
    },
    "logPath": "./"
}
`

func TestLogger(t *testing.T) {
	ioutil.WriteFile("./logger.json", []byte(logConfig), 0660)
	defer os.Remove("./logger.json")
	delete(loggerCategories, "default")
	initLogger("./logger.json")
	for k, v := range loggerCategories {
		fmt.Println("category:", k)
		for _, lc := range v {
			fmt.Println("type:", lc.typ, lc.config)
		}
	}
	logger := Get("default", true)
	logger.Info("hello login and yuanbao")
	logger.Close()
}
