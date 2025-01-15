package dtos

type TransferDto struct {
	ChainId int64   `json:"chainId"`
	Amount  float64 `json:"amount"`
}
