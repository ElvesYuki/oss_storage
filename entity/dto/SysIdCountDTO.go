package dto

type SysIdCount struct {
	Id      int64  `json:"id"`
	Module  string `json:"module"`
	Step    int64  `json:"step"`
	Counter int64  `json:"counter"`
}
