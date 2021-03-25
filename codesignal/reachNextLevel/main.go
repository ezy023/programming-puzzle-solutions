package main

func reachNextLevel(experience, threshold, reward int) bool {
	return experience+reward >= threshold
}
