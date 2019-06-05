package config

type PromptRepo interface {
	FilePath(defaultPath string) string
	ConsumerUrl(defaultUrl string) string
	ConcurrentThreads(defaultThreads int) int
	PushPerRequest(pushPerRequest int) int
	PushBLockSize(defaultBLockSize int64) int64

	ShowGroups(config *Config)
}

type Storage interface {
	GetConfig() (*Config, error)
	SaveConfig(conf *Config) error
}

type Config struct {
	Group []Group `yaml:"group_config"`
}

type Group struct {
	Name    string `yaml:"name"`
	BaseUrl string `yaml:"base_url"`
}
