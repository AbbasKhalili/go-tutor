package dtos

type WalletDto struct {
	Address       string  `json:"address"`
	Balance       float64 `json:"balance"`
	NativeBalance float32 `json:"nativeBalance"`
}
