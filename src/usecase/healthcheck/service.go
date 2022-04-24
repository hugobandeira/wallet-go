package healthcheck

type Service struct {
}

func (s *Service) Health() string {
	return "ok"
}
