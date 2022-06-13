package entities

type Inventory struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Quantity uint `json:"quantity"`
	Deleted bool `json:"deleted"`
	Comment string `json:"comment"`
}
