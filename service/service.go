package service

import "../part"

type Status uint32

var (
	STATUS_INIT       = Status(1 << 0)
	STATUS_STARTUPED  = Status(1 << 1)
	STATUS_SHUTDOWNED = Status(1 << 2)
)

type Service struct {
	_m  uint32      /*magic*/
	_ps []part.Part /*parts*/
	_s  Status      /*status*/
}

func (s *Service) Magic() uint32 {
	return s._m
}

func (s *Service) Init() (err error) {
	for p := range s._ps {
		err = p.Init()
		if err != nil {
			return
		}
	}
	s._s = STATUS_INIT
	return
}

func (s *Service) Startup() (err error) {
	for p := range s._ps {
		err = p.Startup()
		if err != nil {
			return
		}
	}
	s._s = STATUS_STARTUPED
	return
}

func (s *Service) Shutdown() (err error) {
	for p := range s._ps {
		err = p.Shutdown()
		if err != nil {
			return
		}
	}
	s._s = STATUS_SHUTDOWNED
	return
}

func (s *Service) Status() Status {
	return s._s
}

func (s *Service) AddPart(p part.Part) {
	s._ps = append(s._ps, p)
}

func CreateService(m uint32) *Service {
	return &Service{m, []part.Part{}, 0}
}
