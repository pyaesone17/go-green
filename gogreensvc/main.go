package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pyaesone17/gogreen"
	"github.com/spf13/viper"
)

var version string

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmicroseconds)

	cfgFile := flag.String("c", "../", "Configuration path")
	flag.Parse()

	if len(*cfgFile) == 0 {
		flag.Usage()
		return
	}

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AddConfigPath("../")    // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	log.Printf("STARTUP: %s version %s", os.Args[0], version)
	log.Printf("Listening on: %s", viper.GetString("address"))

	svc := gogreen.NewService(viper.GetViper())
	svc.ListenAndServe()
}
