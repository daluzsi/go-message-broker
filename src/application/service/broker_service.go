package service

type brokerService struct{}

func NewListenService() *brokerService {
	return &brokerService{}
}

func (ls *brokerService) Listen() {

}
