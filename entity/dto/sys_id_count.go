package dto

type SysIdCountDTO struct {
	SysIdCountId int64  `json:"sysIdCountId"`
	Module       string `json:"module"`
	Step         int64  `json:"step"`
	Counter      int64  `json:"counter"`
}
