package GoFrequency

import (
	"strings"
)

func Analysis(input string) map[string]int{
	freqMap := make(map[string]int)
	freqMap["len"] = len(input)
	for key := range input{
		freqMap[strings.Split(input, "")[key]]++
	}
	return freqMap
}

func Compare(map1 map[string]int, map2 map[string]int) float64{
	pts := map1["len"] + map2["len"]
	keys := MapKeys(map1)
	for key := range keys{
		if keys[key] != "len" {
			pts = pts - Abs(map1[keys[key]] - map2[keys[key]])
		}
	}
	return float64(pts) / float64(map1["len"]+map2["len"])
}

func MapKeys(m map[string]int) []string{
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}