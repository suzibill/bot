package product

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return []Product{{"123"}, {"Aboba"}, {"Amogus"}}
}
