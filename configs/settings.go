package configs

import "github.com/spf13/viper"

const (
	CONFIG_PATH = "."
	CONFIG_TYPE = "yaml"
	CONFIG_NAME = "config"
)

type Settings struct {
	v *viper.Viper
}

func NewSettings() (*Settings, error) {
	settings := new(Settings)
	settings.v = viper.New()

	settings.v.AddConfigPath(CONFIG_PATH)
	settings.v.SetConfigName(CONFIG_NAME)
	settings.v.SetConfigType(CONFIG_TYPE)

	if err := settings.v.ReadInConfig(); err != nil {
		return nil, err
	}

	return settings, nil
}

func (s *Settings) ReadToStruct(key string, value interface{}) error { // pass the reference
	return s.v.UnmarshalKey(key, value)
}
