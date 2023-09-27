package service

import (
	"nunu-project/internal/middleware"
	"nunu-project/pkg/helper/sid"
	"nunu-project/pkg/log"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *middleware.JWT
}

func NewService(logger *log.Logger, sid *sid.Sid, jwt *middleware.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
	}
}
