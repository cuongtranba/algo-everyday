package prefixsum

func maximumPopulation(logs [][]int) int {
	population := make(map[int]int)

	for _, log := range logs {
		birth, death := log[0], log[1]
		for year := birth; year < death; year++ {
			population[year]++
		}
	}

	maxPopulation := 0
	earliestYear := 1950 // Năm sớm nhất có thể theo đề bài

	for year := 1950; year <= 2050; year++ { // Khoảng năm theo đề bài
		if population[year] > maxPopulation {
			maxPopulation = population[year]
			earliestYear = year
		}
	}

	return earliestYear
}
