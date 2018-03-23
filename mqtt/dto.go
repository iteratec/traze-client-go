package mqtt

type (
	Game struct {
		Name          string `json:"name"`
		ActivePlayers int    `json:"activePlayers"`
	}
	Bike struct {
		PlayerId        int      `json:"playerId"`
		CurrentLocation [2]int   `json:"currentLocation"`
		Direction       string   `json:"direction"`
		Trail           [][2]int `json:"trail"`
	}
	Grid struct {
		Height int     `json:"height"`
		Width  int     `json:"width"`
		Tiles  [][]int `json:"tiles"`
		Bikes  []Bike  `json:"bikes"`
	}
	Player struct {
		id              int    `json:"id"`
		name            string `json:"name"`
		owned           int    `json:"owned"`
		color           string `json:"color"`
		frags           int    `json:"frags"`
	}
)
