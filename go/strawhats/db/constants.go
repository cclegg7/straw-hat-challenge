package db

type BoulderDifficulty uint8

const (
	BoulderDifficulty_VB   BoulderDifficulty = 0
	BoulderDifficulty_V0_1 BoulderDifficulty = 1
	BoulderDifficulty_V1_2 BoulderDifficulty = 2
	BoulderDifficulty_V2_4 BoulderDifficulty = 3
	BoulderDifficulty_V4_6 BoulderDifficulty = 4
	BoulderDifficulty_V6_8 BoulderDifficulty = 5
)

type TopRopeDifficulty uint8

const (
	TopRopeDifficulty_5_6  TopRopeDifficulty = 0
	TopRopeDifficulty_5_7  TopRopeDifficulty = 1
	TopRopeDifficulty_5_8  TopRopeDifficulty = 2
	TopRopeDifficulty_5_9  TopRopeDifficulty = 3
	TopRopeDifficulty_5_10 TopRopeDifficulty = 4
	TopRopeDifficulty_5_11 TopRopeDifficulty = 5
	TopRopeDifficulty_5_12 TopRopeDifficulty = 6
)
