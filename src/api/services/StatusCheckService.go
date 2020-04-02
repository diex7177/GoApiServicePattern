package services

//StatusCheckService ...
type StatusCheckService struct {
}

//GetStatusHealth ...
func (s StatusCheckService) GetStatusHealth() (string, error) {
	return "Pong", nil
}
