package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var jsonData map[string]interface{}

func setField(obj interface{}, name string, value interface{}) error {
	structData := reflect.ValueOf(obj).Elem()
	fieldValue := structData.FieldByName(name)

	if !fieldValue.IsValid() {
		return fmt.Errorf("utils.setField() No such field: %s in obj ", name)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value ", name)
	}

	fieldType := fieldValue.Type()
	val       := reflect.ValueOf(value)

	valTypeStr   := val.Type().String()
	fieldTypeStr := fieldType.String()
	if valTypeStr == "float64" && fieldTypeStr == "int" {
		val = val.Convert(fieldType)
	} else if fieldType != val.Type() {
		return fmt.Errorf("Provided value type " + valTypeStr + " didn't match obj field type " + fieldTypeStr)
	}
	fieldValue.Set(val)
	return nil
}

func setStructByJSON(obj interface{}, mapData map[string]interface{}) error {
	for key, value := range mapData {
		if err := setField(obj, key, value); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func initJson()  {
	bytes, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println("Read file ", err)
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("Invaild config: ", err.Error())
		os.Exit(-1)
	}

}

type dBConfig struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	URL          string
	Charset      string
}

var DBConfig dBConfig

func initDB()  {
	 if err := setStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}));err != nil {
	 	fmt.Println(err)
		 return
	 }
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
	DBConfig.URL = url
}


type serverConfig struct {
	APIPoweredBy       string
	SiteName           string
	Host               string
	ImgHost            string
	Env                string
	LogDir             string
	LogFile            string
	ClientAPIPrefix    string
	ServerAPIPrefix    string
	UploadImgDir       string
	ImgPath            string
	MaxMultipartMemory int
	Port               int
	StatsEnabled       bool
	TokenSecret        string
	TokenMaxAge        int
	PassSalt           string
	MailUser           string //域名邮箱账号
	MailPass 		   string
	MailHost           string //smtp邮箱域名
	MailPort           int    //smtp邮箱端口
	MailFrom           string //邮件来源
}


var ServerConfig serverConfig

func initServer()  {
	if err := setStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}));err != nil {
		fmt.Println(err)
		return
	}
	sep := string(os.PathListSeparator)
	execPath, _ := os.Getwd()
	length := utf8.RuneCountInString(execPath)
	lastChar := execPath[length-1:]
	if lastChar != sep {
		execPath = execPath + sep
	}

	if ServerConfig.UploadImgDir == "" {
		pathArr := []string{"website", "static", "upload", "img"}
		uploadImgDir := execPath + strings.Join(pathArr, sep)
		ServerConfig.UploadImgDir = uploadImgDir
	}

	ymdStr := time.Now().Format("2006-01-02")

	if ServerConfig.LogDir == "" {
		ServerConfig.LogDir = execPath
	} else {
		length := utf8.RuneCountInString(ServerConfig.LogDir)
		lastChar := ServerConfig.LogDir[length-1:]
		if lastChar != sep {
			ServerConfig.LogDir = ServerConfig.LogDir + sep
		}
	}

	ServerConfig.LogFile = ServerConfig.LogDir + ymdStr + ".log"

}

func init()  {
	initJson()
	initServer()
	initDB()
}