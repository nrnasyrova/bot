package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allproducts
}

func (s *Service) Get(id int) (*Product, bool) {
	if id <0 || id >= len(allproducts) {
		return nil, false
	}
	return &allproducts[id], true
}
