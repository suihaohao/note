package use_neo4j

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/spf13/viper"
)

var neo4jDriver *neo4j.Driver

func GetNeo4jConn() *neo4j.Driver {
	return neo4jDriver
}

var neo4jDbConfig = map[string]interface{}{
	"neo4j": map[string]interface{}{
		"root":     "neo4j",
		"ip":       "127.0.0.1",
		"port":     7687,
		"password": "my@neo4j",
	},
}

func Init() {
	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }
	driver, err := neo4j.NewDriver(fmt.Sprintf("bolt://%s:%d", viper.GetString("neo4j.ip"), viper.GetInt("neo4j.port")), neo4j.BasicAuth(viper.GetString("neo4j.root"), viper.GetString("neo4j.password"), ""), configForNeo4j40)
	fmt.Println(fmt.Sprintf("bolt://%s:%d", viper.GetString("neo4j.ip"), viper.GetInt("neo4j.port")))
	if err != nil {
		panic(err)
	}
	neo4jDriver = &driver
}
