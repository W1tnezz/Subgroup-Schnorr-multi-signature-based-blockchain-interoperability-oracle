package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"ioporaclenode/pkg/test"
	"os"
	"os/signal"
)

func main() {
	configFile := flag.String("c", "./configs/config.json", "filename of the config file")
	flag.Parse()

	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read config: %v", err)
	}

	var config test.Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unmarshal config into struct, %v", err)
	}

	log.Infof("Loaded config file %s", *configFile)

	testAccount, err := test.NewTestAccount(config)
	if err != nil {
		log.Fatalf("New oracle node: %v", err)
	}

	go func() {
		if err := testAccount.Run(); err != nil {
			log.Fatalf("Run node: %v", err)
		} // 运行node
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

}
