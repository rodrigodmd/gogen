package template

import (
	"github.com/rodrigodmd/gogen/pkg/domain/config"
)

const (
	fileSuccessFormat = "%d_success.csv"
	fileErrorFormat   = "%d_error.csv"

	clientUrl = "http://develop-worker.merchant-orders-api.melifrontends.com/merchant_orders/consumer/publish_topic"
)

type ProcessFunc func(line []string) error

type Service interface {
	//CheckStatus() status.Status
	Initialize() error
	//Download(url string) error
}

type service struct {
	currentConfig *config.Config
	conf          config.Service

}

func New(conf config.Service) Service {
	proc := service{
		conf:   conf,
	}

	return &proc
}

func (s *service) Initialize() error {
	s.currentConfig = s.conf.GetConfig()



	return nil
}

//func (s *service) CheckStatus() status.Status {
//	return s.status.
//}


