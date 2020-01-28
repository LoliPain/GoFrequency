package GoFrequency

import (
	"math"
	"strings"
)

func Map(input string) map[string]int{
	countMap := make(map[string]int)
	countMap["len"] = len(input)
	for key := range input{
		countMap[strings.Split(input, "")[key]]++
	}
	return countMap
}

func Score(map1 map[string]int, map2 map[string]int) float64{
	pts := float64(map1["len"] + map2["len"])
	var score float64
	keys := Keys(map1)
	for key := range keys{
		if keys[key] != "len" {
			score = pts - math.Abs(float64(map1[keys[key]]) - float64(map2[keys[key]]))
		}
	}
	return score / pts
}

func Keys(m map[string]int) []string{
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Analysis(input map[string]int) map[string]float64{
	keys := Keys(input)
	freqMap := make(map[string]float64)
	for key := range keys{
		if keys[key] != "len" {
			freqMap[keys[key]] = float64(input[keys[key]]) / float64(input["len"]) * 100
		}
	}
	return freqMap
}