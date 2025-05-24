package service

type Room struct {
	Room    int    `json:"room"`
	P1      string `json:"p1"`
	P2      string `json:"p2"`
	IsStart bool   `json:"is_start"`
	Map     *Map   `json:"map"`
}
