package day11

type state [][]uint8

const (
	floor            uint8 = iota
	unoccupied_chair uint8 = iota
	occupied_chair   uint8 = iota
)

const (
	top_left     uint8 = iota
	top_mid      uint8 = iota
	top_right    uint8 = iota
	mid_left     uint8 = iota
	mid_right    uint8 = iota
	bottom_left  uint8 = iota
	bottom_mid   uint8 = iota
	bottom_right uint8 = iota
)

// map of direction to pair of row offset and col offset
var direction_map = map[uint8][]int{
	top_left:     {-1, -1},
	top_mid:      {-1, 0},
	top_right:    {-1, 1},
	mid_left:     {0, -1},
	mid_right:    {0, 1},
	bottom_left:  {1, -1},
	bottom_mid:   {1, 0},
	bottom_right: {1, 1},
}
