package value_objects

type Address struct {
	Street   string `json:"street"`
	Province string `json:"province"`
	City     string `json:"city"`
}
