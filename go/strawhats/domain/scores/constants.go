package scores

type rotation struct {
	startWeek int
	endWeek   int
	distance  int
}

var rotations = []*rotation{
	{startWeek: 16, endWeek: 19, distance: 4},
	{startWeek: 20, endWeek: 23, distance: 8},
	{startWeek: 24, endWeek: 27, distance: 11},
	{startWeek: 28, endWeek: 31, distance: 14},
}
