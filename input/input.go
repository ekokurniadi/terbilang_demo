package input

type NominalInput struct {
	Angka string `json:"angka" binding:"required"`
}
