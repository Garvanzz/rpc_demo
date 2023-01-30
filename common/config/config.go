package config

var Conf *Config

type Config struct {
	Common CommonConfig
	Task   TaskConfig
}

func init() {

}

type CommonConfig struct {
}

type TaskConfig struct {
	PushChan     int
	PushChanSize int
}
