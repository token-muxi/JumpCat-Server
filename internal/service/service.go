package service

type Room struct {
	Room     int    `json:"room"`
	P1       string `json:"p1"`
	P2       string `json:"p2"`
	P1_ready bool   `json:"p1_ready"`
	P2_ready bool   `json:"p2_ready"`
	Map      Map   `json:"map"`
}
	