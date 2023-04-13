package models

type User struct {
	ID             int
	Name           string
	CharacterToken string
	Boulder        BoulderDifficulty
	TopRope        TopRopeDifficulty
}
