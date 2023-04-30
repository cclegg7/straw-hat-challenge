package scores

import (
	"github.com/cclegg7/straw-hat-challenge/models"
)

func (c *Calculator) runningScore(user *models.User) (int, error) {
	runsByWeek, err := c.database.GetRunsByWeekForUserID(user.ID)
	if err != nil {
		return 0, err
	}

	startWeek := rotations[0].startWeek
	lastWeek := rotations[len(rotations)-1].endWeek
	runCounts := []int{}
	for currentWeek := startWeek; currentWeek <= lastWeek; currentWeek++ {
		runs, ok := runsByWeek[currentWeek]
		if !ok {
			runCounts = append(runCounts, 0)
			continue
		}

		totalRuns := runs.Weekday + runs.Weekend
		if totalRuns >= 3 {
			runCounts = append(runCounts, 3)
			continue
		}
		if currentWeek < lastWeek {
			neededRuns := 3 - totalRuns
			if nextWeekRuns, ok := runsByWeek[currentWeek+1]; ok {
				if neededRuns <= nextWeekRuns.Weekend {
					totalRuns = 3
					nextWeekRuns.Weekend = nextWeekRuns.Weekend - neededRuns
				}
			}
		}
		runCounts = append(runCounts, totalRuns)
	}

	score := 0
	for _, count := range runCounts {
		score += 50 * count
		if count == 3 {
			score += 100
		}
	}

	var completedDistanceGoals int
	for _, rotation := range rotations {
		for currentWeek := rotation.startWeek; currentWeek <= rotation.endWeek; currentWeek++ {
			if runs, ok := runsByWeek[currentWeek]; ok && runs.MaxDistance >= rotation.distance {
				completedDistanceGoals++
				break
			}
		}
	}

	score += 100 * completedDistanceGoals

	return score, err
}
