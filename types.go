package matekasse

type ID uint32
type account struct {
	Id      ID     `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
type nameIdPair struct {
	Id   ID     `json:"id"`
	Name string `json:"name"`
}
