package config

import (
	"ConcurrentImageFileServer/src/utils/logUtil"
	"github.com/spf13/viper"
	_ "go.uber.org/zap"
	"os"
)

var LogPath string

var ServerConfig = struct {
	// Maximum size of single upload file
	MaxUploadSize 		int `mapstructure:"max_upload_size"`

	// Maximum overall request size
	MaxRequestBodySize 	int `mapstructure:"max_request_body_size"`

	// Upload file root directory
	UploadPath 			string `mapstructure:"upload_path"`

	// File Operating Port
	OperationPort 		string `mapstructure:"operation_port"`

	// File Visit port
	VisitPort 			string `mapstructure:"visit_port"`

	// Operation certificate
	OperationToken 		string `mapstructure:"operation_token"`

	// Generate a file directory index
	GenerateIndexPages 	bool `mapstructure:"generate_index_pages"`
} {}

func initServerConf() {
	confViper := viper.New()
	confViper.SetConfigFile(PWD + "config.yml")
	if err := confViper.ReadInConfig(); err != nil {
		panic(err)
	}
	// Unmarshall map["string"] interface{} structure data from viper
	if err := confViper.Unmarshal(&ServerConfig); err != nil {
		panic(err)
	}

	printUtilMessage(
		"Current file server configuration is:\n" +
		logUtil.GetStructureDataInfo(ServerConfig, "ServerConfig"),
	)

	ServerConfig.MaxUploadSize 		*= MB
	ServerConfig.MaxRequestBodySize *= MB

	ServerConfig.VisitPort 		= ":" + ServerConfig.VisitPort
	ServerConfig.OperationPort 	= ":" + ServerConfig.OperationPort

	if _, err := os.Stat(PWD + ServerConfig.UploadPath); err != nil {
		if os.IsNotExist(err) {
			printUtilMessage("Dir: /" + ServerConfig.UploadPath + " does not exist")
			err = os.Mkdir( PWD + ServerConfig.UploadPath, os.ModePerm)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		printUtilMessage("Dir: /" + ServerConfig.UploadPath + " already exists")
	}
}