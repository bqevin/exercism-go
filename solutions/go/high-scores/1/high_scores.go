package highscores

import "slices"

type HighScores struct{
    all, topThree []int
    latest, best int
    
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
    l := len(scores)
    latest := scores[l - 1];
    sorted := make([]int, l);
    numTop := 3
    
    copy(sorted, scores)
    
    slices.SortFunc(sorted, func(a, b int) int {
        return b - a
    });
    
    if l < 3 {
        numTop = l
    }
    
	return &HighScores{
        all: scores,
        topThree: sorted[:numTop],
        latest: latest,
        best: sorted[0],
    }
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.all;
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.latest
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	return s.best;
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	return s.topThree;
}
