package GoFrequency

import (
	"math"
	"strings"
)

func Map(input string) map[string]int {
	countMap := make(map[string]int)
	countMap["len"] = len(strings.Split(input, ""))
	for key := 1; key < countMap["len"]; key++ {
		countMap[strings.Split(input, "")[key]]++
	}
	return countMap
}

func Score(map1 map[string]int, map2 map[string]int) float64 {
	keys := Keys(map1)
	map2Keys := Keys(map2)
	for _, key := range map2Keys {
		i := 0
		for _, item := range keys {
			if item == key {
				i++
			}
		}
		if i == 0 {
			keys = append(keys, key)
		}
	}
	pts := float64(map1["len"] + map2["len"])
	score := pts
	for _, key := range keys {
		score = score - math.Abs(float64(map1[key])-float64(map2[key]))
	}
	return score / pts
}

func Keys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		if k != "len" {
			keys = append(keys, k)
		}
	}
	return keys
}

func Analysis(input map[string]int) map[string]float64 {
	keys := Keys(input)
	freqMap := make(map[string]float64)
	for key := range keys {
		freqMap[keys[key]] = float64(input[keys[key]]) / float64(input["len"]) * 100
	}
	return freqMap
}
