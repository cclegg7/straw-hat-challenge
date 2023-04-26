package scores

import (
	"github.com/cclegg7/straw-hat-challenge/models"
)

const (
	baseScorePerClimb         = 100
	difficultyScoreAdjustment = 50
	challengeBonus            = 25
)

func (c *Calculator) climbingScore(user *models.User) (int, error) {
	var topBoulders []*models.Climb
	var topTopRopes []*models.Climb

	for _, rotation := range rotations {
		startWeek := rotation.startWeek
		endWeek := rotation.endWeek
		boulders, err := c.database.GetTopClimbsInCategoryForUserID(user.ID, startWeek, endWeek, models.ClimbCategory_Boulder)
		if err != nil {
			return 0, err
		}
		topBoulders = append(topBoulders, boulders...)
		topRopes, err := c.database.GetTopClimbsInCategoryForUserID(user.ID, startWeek, endWeek, models.ClimbCategory_TopRope)
		if err != nil {
			return 0, err
		}
		topTopRopes = append(topTopRopes, topRopes...)
	}

	score := 0
	for _, boulder := range topBoulders {
		score += scoreForBoulder(user, models.BoulderDifficulty(boulder.Rating), boulder.IsChallenge)
	}
	for _, topRope := range topTopRopes {
		score += scoreForTopRope(user, models.TopRopeDifficulty(topRope.Rating), topRope.IsChallenge)
	}

	return score, nil
}

func scoreForBoulder(user *models.User, difficulty models.BoulderDifficulty, challenge bool) int {
	ratingDiff := int(difficulty) - int(user.Boulder)
	score := baseScorePerClimb + difficultyScoreAdjustment*ratingDiff
	if challenge {
		score += challengeBonus
	}

	if score < 0 {
		return 0
	} else {
		return score
	}
}

func scoreForTopRope(user *models.User, difficulty models.TopRopeDifficulty, challenge bool) int {
	ratingDiff := int(difficulty) - int(user.TopRope)
	score := baseScorePerClimb + difficultyScoreAdjustment*ratingDiff
	if challenge {
		score += challengeBonus
	}
	if score < 0 {
		return 0
	} else {
		return score
	}
}
