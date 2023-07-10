package logger

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/astaxie/beego/logs"
)

var defaultLogConf = "./logger.json"
var defaultLogFileName = ""

type loggerConfig struct {
	Types []string
}

type logConf struct {
	typ    string
	config string
}

var loggers = make(map[string]*logs.BeeLogger)
var loggerTypes = make(map[string]string)
var loggerCategories = make(map[string][]*logConf)
var logLevel = make(map[string]int)
var levelMap = map[string]int{
	"EMERGENCY":     logs.LevelEmergency,
	"ALERT":         logs.LevelAlert,
	"CRITICAL":      logs.LevelCritical,
	"CRIT":          logs.LevelCritical,
	"ERROR":         logs.LevelError,
	"ERR":           logs.LevelError,
	"WARNING":       logs.LevelWarning,
	"WARN":          logs.LevelWarning,
	"NOTICE":        logs.LevelNotice,
	"INFORMATIONAL": logs.LevelInformational,
	"INFO":          logs.LevelInformational,
	"DEBUG":         logs.LevelDebug,
}

func init() {
	// because flag.Parse must be called after all flags defined, but we want to get logConf in init func,
	// so here we parse logConf manually
	logConfFile,defaultFileName := getLogConf()
	defaultLogFileName = defaultFileName
	if fileInfo, err := os.Stat(logConfFile); err == nil && !fileInfo.IsDir() {
		err = initLogger(logConfFile)
		if err != nil {
			panic(err)
		}
	}
	if _, ok := loggerCategories["default"]; !ok {
		loggerCategories["default"] = []*logConf{
			&logConf{typ: "console", config: fmt.Sprintf(`{"level": %d}`, logs.LevelInformational)},
		}
	}
}

func getLogConf() (string,string) {
	file := ""
	defaultFileName := ""
	for _, arg := range os.Args[1:] {
		fields := strings.Split(arg, "=")
		if len(fields) != 2  || len(fields[1]) == 0{
			continue
		}
		if fields[0] == "-logconf" || fields[0] == "--logconf" {
			file = fields[1]
		} else if fields[0] == "-deflogfilename" || fields[0] == "--deflogfilename" {
			defaultFileName = fields[1]
		}
	}
	if file != "" {
		return file,defaultFileName
	}
	return defaultLogConf,""
}

func getAsString(v interface{}) (string, bool) {
	if v == nil {
		return "", false
	}
	str, ok := v.(string)
	return str, ok
}

func checkLevels(levels map[string]interface{}) error {
	for k, v := range levels {
		level, _ := getAsString(v)
		if _, ok := levelMap[level]; !ok {
			return errors.New("unknown log level for " + k)
		}
	}
	return nil
}

// 防止float64在Marshal时产生科学计数，比如file类型的maxsize字段
func fixFloatingPoint(values map[string]interface{}) {
	for k, v := range values {
		switch value := v.(type) {
		case float64:
			values[k] = int(value)
		}
	}
}

func getLogPath(logPath interface{}) (string, error) {
	logDir, _ := getAsString(logPath)
	if len(logDir) == 0 {
		return "", nil
	}
	logDir, err := filepath.Abs(logDir)
	if err != nil {
		return "", err
	}
	return logDir, nil
}

func ensureFilePath(fileName string) {
	dir := filepath.Dir(fileName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
}

func replaceVariable(content []byte) []byte {
	_, programName := filepath.Split(os.Args[0])
	content = regexp.MustCompile("\\$\\{programName\\}").ReplaceAll(content, []byte(programName))
	return content
}

func initLogger(configFileName string) error {
	content, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}
	var config map[string]interface{}
	err = json.Unmarshal(replaceVariable(content), &config)
	if err != nil {
		return err
	}
	var levels = config["levels"].(map[string]interface{})
	var types = config["types"].([]interface{})
	logPath, err := getLogPath(config["logPath"])
	if err != nil {
		return err
	}
	err = checkLevels(levels)
	if err != nil {
		return err
	}
	for k,v := range levels {
		logLevel[k] = levelMap[v.(string)]
	}
	for _, v := range types {
		var typeConfig = v.(map[string]interface{})
		fixFloatingPoint(typeConfig)
		typ, ok := getAsString(typeConfig["type"])
		if !ok {
			continue
		}
		if typ != "console" && typ != "file" {
			continue
		}
		delete(typeConfig, "type")
		category, ok := getAsString(typeConfig["category"])
		if !ok {
			category = "default"
		}
		delete(typeConfig, "category")
		if typ == "file" && len(logPath) > 0 {
			fileName, _ := getAsString(typeConfig["filename"])
			if category == "default" && defaultLogFileName != "" {
				fileName = defaultLogFileName
			}
			fileName = path.Join(logPath, fileName)
			ensureFilePath(fileName)
			typeConfig["filename"] = fileName
		}

		var categories = strings.Split(category, ",")
		for _, category := range categories {
			category = strings.TrimSpace(category)

			levelStr, _ := getAsString(levels[category])
			if level, ok := levelMap[levelStr]; ok {
				typeConfig["level"] = level
			} else {
				typeConfig["level"] = logs.LevelInformational
			}
			rb, _ := json.Marshal(typeConfig)
			loggerCategories[category] = append(loggerCategories[category], &logConf{typ: typ, config: string(rb)})
		}
	}
	return nil
}

// Init register command line flag, but only for a placeholder in command usage, as this package parses command line itself
func Init() {
	flag.String("logconf", defaultLogConf, "specify log config file")
	flag.String("deflogfilename", "", "default log fileName")
	DeafultLoggerInit()
}

// Get get a logger by category
func Get(name string, isAsync bool) *logs.BeeLogger {
	if beeLog, ok := loggers[name]; ok {
		return beeLog
	}
	configs, ok := loggerCategories[name]
	if !ok {
		configs = loggerCategories["default"]
	}
	beeLog := logs.NewLogger(10000)
	if isAsync {
		beeLog = beeLog.Async()
	}
	for _, lc := range configs {
		beeLog.SetLogger(lc.typ, lc.config)
		beeLog.EnableFuncCallDepth(true)
	}
	loggers[name] = beeLog
	return beeLog
}

func GetLogLevel( name string ) int {
	if logLevel[name] == 0 {
		return logs.LevelInformational
	}
	return logLevel[name]
}

// Close close all opened logger, call this when program exits
func Close() {
	for _, logger := range loggers {
		logger.Close()
	}
}
