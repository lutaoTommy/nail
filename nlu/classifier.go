package nlu

import "strings"

type Result struct {
	Intent Intent
	Score  int
}

func Classify(text string) Result {
	if text == "" {
		return Result{Intent: IntentUnknown}
	}

	var best rule
	bestScore := 0

	for _, r := range rules {
		score := 0
		for _, t := range r.triggers {
			if strings.Contains(text, t) {
				score += r.priority
			}
		}
		if score > bestScore {
			bestScore = score
			best = r
		}
	}

	if bestScore == 0 {
		return Result{Intent: IntentUnknown}
	}

	return Result{
		Intent: best.intent,
		Score:  bestScore,
	}
}
