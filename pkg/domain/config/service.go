package config



type Service interface {
	ShowConfig()
	GetConfig() *Config

	ConfigAll()
}

type service struct {
	prompt  PromptRepo
	storage Storage
	config  *Config
}

func NewService(prompt PromptRepo, storage Storage) Service {
	srv := service{
		prompt:  prompt,
		storage: storage,
	}
	srv.config = srv.GetConfig()
	return &srv
}

func (s *service) GetConfig() *Config {
	config, err := s.storage.GetConfig()
	if err != nil {
		config = &Config{
			Group: []Group{},
		}
	}
	return config
}

func (s *service) ConfigToken() {
	c := s.config
	c.FuryToken = s.prompt.FuryToken(c.FuryToken)

	s.storage.SaveConfig(c)
}

func (s *service) ConfigRun() {
	c := s.config

	c.ConsumerUrl = s.prompt.ConsumerUrl(c.ConsumerUrl)
	c.FilePath = s.prompt.FilePath("./")
	c.PublishCount = s.prompt.PushBLockSize(c.PublishCount)

	s.storage.SaveConfig(c)
}

func (s *service) ConfigAll() {
	c := s.config

	c.ConsumerUrl = s.prompt.ConsumerUrl(c.ConsumerUrl)
	c.FuryToken = s.prompt.FuryToken(c.FuryToken)

	c.FilePath = s.prompt.FilePath("./")
	c.PublishCount = s.prompt.PushBLockSize(c.PublishCount)
	c.PublishPerRequest = s.prompt.PushPerRequest(c.PublishPerRequest)
	c.ConcurrentThreads = s.prompt.ConcurrentThreads(c.ConcurrentThreads)

	s.storage.SaveConfig(c)
}

func (s *service) ShowConfig() {
	s.prompt.ShowConfig(s.config)
}
