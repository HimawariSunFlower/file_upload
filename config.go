package main

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

var (
	ossKey     string
	ossSecret  string
	projectMap map[string]*Project
)

type Project struct {
	OssPath         string   `mapstructure:"osspath"`     //上传的oss路径
	OssPathEndpoint string   `mapstructure:"ossendpoint"` //oss的endpoint
	BucketName      string   `mapstructure:"bucketname"`  //oss的bucketname
	Path            string   `mapstructure:"path"`        //project的绝对路径
	Versionfile     string   `mapstructure:"versionfile"` //基于path的取版本号的文件路径
	Files           []string `mapstructure:"files"`       //基于path的路径
}

func loadConfig() {
	viper := viper.New()
	viper.SetConfigType("toml")
	viper.SetConfigFile("./file_upload_config.toml")
	viper.ReadInConfig()
	ossKey = viper.GetString("osskeyid")
	ossSecret = viper.GetString("osskeysecret")

	err := viper.UnmarshalKey("projectList", &projectMap)
	if err != nil {
		panic(err)
	}
}

func getVersion(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	filestr := string(file)
	//filestr := `         VERSION string = "1.0"`
	regexp, _ := regexp.Compile(`VERSION+.*\B`)
	version := regexp.FindString(filestr)
	if version == "" {
		panic("找不到version数据,请检查配置")
	}

	return version[strings.Index(version, `"`)+1 : strings.LastIndex(version, `"`)]
}

func spliceDirStr(left, right string) string {
	if strings.HasSuffix(left, "/") && strings.HasPrefix(right, "/") {
		return left + right[1:]
	} else if strings.HasSuffix(left, "/") || strings.HasPrefix(right, "/") {
		return left + right
	} else {
		return left + "/" + right
	}
}
