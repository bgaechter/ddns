package ddnscli

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)
func GetPublicIPAddress() string {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Fatal(err)
	}
	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("public ip address is %s", ip)
	return string(ip)
}


func LoadConfig() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.godyn")   // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config file not found; ignore error if desired
		// Call init
	} else {
		// Config file was found but another error was produced
	}
	}
}

func InitConfig() {

}
