package healthcheck

type Service struct {
}

func (s *Service) Execute() string {
	return "ok"
}
