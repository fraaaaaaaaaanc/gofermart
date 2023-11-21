package handlers_models

type RespUserBalance struct {
	UserBalance      float64 `json:"current"`
	WithdrawnBalance float64 `json:"withdrawn"`
}
