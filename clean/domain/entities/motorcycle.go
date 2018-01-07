package entities

type Motorcycle struct {
	Id    uint64 `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  string `json:"year"`
	//createdUtc  time.Time `json:"createdUtc"`
	//deletedUtc  time.Time `json:"deletedUtc"`
	//modifiedUtc time.Time `json:"modifiedUtc"`
	//rowVersion  []byte    `json:"rowVersion"`
}
