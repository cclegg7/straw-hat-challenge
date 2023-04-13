package scoreboard

import (
	"fmt"

	"github.com/cclegg7/straw-hat-challenge/models"
)

const (
	baseScorePerClimb         = 100
	difficultyScoreAdjustment = 50
)

func (s *Scoreboard) climbingScores(users []*models.User) (map[*models.User]int, error) {
	climbsByUser, err := s.database.GetTopClimbsForUsers(users, s.start, s.end)
	if err != nil {
		return nil, fmt.Errorf("error fetching user climbs: %w", err)
	}

	scores := make(map[*models.User]int)
	for user, climbs := range climbsByUser {
		score := 0
		for difficulty, count := range climbs.Boulders {
			score += count * scoreForBoulder(user, difficulty)
		}
		for difficulty, count := range climbs.TopRopes {
			score += count * scoreForTopRope(user, difficulty)
		}
		scores[user] = score
	}

	return scores, nil
}

func scoreForBoulder(user *models.User, difficulty models.BoulderDifficulty) int {
	score := baseScorePerClimb + difficultyScoreAdjustment*(difficulty-user.Boulder)
	if score < 0 {
		return 0
	}

	return int(score)
}

func scoreForTopRope(user *models.User, difficulty models.TopRopeDifficulty) int {
	score := baseScorePerClimb + difficultyScoreAdjustment*(difficulty-user.TopRope)
	if score < 0 {
		return 0
	}

	return int(score)
}
