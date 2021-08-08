package conf

import (
	"HMThird/internal/kit"
	"encoding/json"
	"fmt"
)

const (
	config string = "connectionStrings.json"
)

var connectionStrings []connectionString

func init() {
	connectionStrings = initConnectionStrings()
}

type connectionString struct {
	dbName   string
	driver   string
	user     string
	ip       string
	port     string
	password string
}

func GetConnectionString(dbName string) (string, string, error) {
	var result connectionString
	for _, v := range connectionStrings {
		if v.dbName == dbName {
			result = v
			break
		}
	}
	return result.driver, result.getDataSource(), nil
}

func initConnectionStrings() []connectionString {
	jsonFilePath := "../../configs/" + config
	str := kit.ReadAsString(jsonFilePath)
	connectionStrings := make([]connectionString, 0)
	err := json.Unmarshal([]byte(str), &connectionStrings)
	if err != nil {
		return nil
	}
	return connectionStrings
}

func (c *connectionString) getDataSource() string {
	var source string
	switch c.driver {
	case "mysql":
		source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", c.user, c.password, c.ip, c.password, c.dbName)
	default:
		source = ""
	}
	return source
}
