package ddnscli

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	hostedZone    string
	subdomainName string
}

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

func setupFlags(){
	flag.String("hosted-zone", "", "Hosted Zone in which the dynamic entry should be created")
	flag.String("subdomain", "ddns", "Subdomain name that should be created")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func LoadConfig() {

	setupFlags()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.ddns")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.WriteConfig()
			// Config file not found; ignore error if desired
			// Call init
		} else {
			log.Fatal(err)
		}
	}
}

func InitConfig() {

}
