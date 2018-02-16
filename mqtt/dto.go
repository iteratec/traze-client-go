package mqtt

type (
	Grid struct {
		Height int     `json:"height"`
		Width  int     `json:"width"`
		Tiles  [][]int `json:"tiles"`
	}
)
