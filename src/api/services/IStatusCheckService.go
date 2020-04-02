package services

//IStatusCheckService ...
type IStatusCheckService interface {
	GetStatusHealth() (string, error)
}
