package magicCards

type MagicCard struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Collection string `json:"collection"`
	Cost [3]int `json:"cost"`
}
