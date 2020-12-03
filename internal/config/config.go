package config

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

type Configuration struct {
	Key         string
	AsEnvName   string
	Description string
	Value       interface{}
}

func (configuration Configuration) GetStringVal() string {
	return configuration.Value.(string)
}

func (configuration Configuration) GetIntVal() int {
	return configuration.Value.(int)
}

func (configuration Configuration) GetBoolVal() bool {
	return configuration.Value.(bool)
}

func (configuration Configuration) Register() {

	switch value := configuration.Value.(type) {
	case string:
		flag.StringVar(&value, configuration.Key, LookupEnvOrString(configuration.AsEnvName, value), configuration.Description)
	case int:
		flag.IntVar(&value, configuration.Key, LookupEnvOrInt(configuration.AsEnvName, value), configuration.Description)
	case bool:
		flag.BoolVar(&value, configuration.Key, LookupEnvOrBool(configuration.AsEnvName, value), configuration.Description)
	}

}

type ConfigurationSet map[string]*Configuration

var singleton = make(ConfigurationSet)

func Get(key string) *Configuration {
	configuration, _ := singleton[key]
	return configuration
}

func (set ConfigurationSet) Add(key string, asEnvName string, defaultValue interface{}, description string) ConfigurationSet {
	set[key] = &Configuration{
		Key:         key,
		AsEnvName:   asEnvName,
		Description: description,
		Value:       defaultValue,
	}
	return set
}

func Add(key string, asEnvName string, defaultValue interface{}, description string) ConfigurationSet {
	return singleton.Add(key, asEnvName, defaultValue, description)
}

func (set ConfigurationSet) Register() ConfigurationSet {

	for _, configuration := range set {
		configuration.Register()
	}

	return set
}

func Register() ConfigurationSet {
	return singleton.Register()
}

func (set ConfigurationSet) Load() ConfigurationSet {

	flag.Parse()

	return set
}

func Load() ConfigurationSet {
	return singleton.Load()
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("LookupEnvOrInt[%s]: %v", key, err)
		}
		return v
	}
	return defaultVal
}

func LookupEnvOrBool(key string, defaultVal bool) bool {
	defaultStringVal := "false"
	if defaultVal {
		defaultStringVal = "true"
	}

	val := strings.Trim(strings.ToLower(LookupEnvOrString(key, defaultStringVal)), " ")

	if val == "true" || val == "t" || val == "y" || val == "yes" || val == "on" {
		return true
	}

	return false
}
