package main

import "fmt"
import "strconv"
import "strings"

// Question One
func numberOfHits(displayCounts []string) map[string]int {
	hitsByDomain := map[string]int{}
	for _, countStr := range displayCounts {
		countPieces := strings.Split(countStr, ",")
		hits, err := strconv.Atoi(countPieces[0])
		if err != nil {
			panic(fmt.Sprintf("Cannot convert %s to int", countPieces[0]))
		}

		fullDomain := countPieces[1]
		charIndex := len(fullDomain)-1
		for charIndex >= -1 {
			if charIndex == -1 || rune(fullDomain[charIndex]) == rune('.') {
				domain := fullDomain[charIndex+1:]
				hitsByDomain[domain] += hits
			}
			charIndex--
		}
	}
	return hitsByDomain
}

type trail struct {
	pieces 		[]string
	lastIndex 	int
}

// Question Two
func longestCommonTrail(userOne []string, userTwo []string) []string {
	// preprocess the first user into a map of url -> place in the list
	userOneVisits := map[string]int{}
	for idx, url := range userOne {
		userOneVisits[url] = idx
	}

	commonTrails := []*trail{}
	for _, url := range userTwo {
		if userOneIdx, ok := userOneVisits[url]; ok {
			var found bool
			for _, commonTrail := range commonTrails {
				if commonTrail.lastIndex < userOneIdx {
					found = true
					commonTrail.pieces = append(commonTrail.pieces, url)
					commonTrail.lastIndex = userOneIdx
				}
			}
			if !found {
				commonTrails = append(commonTrails, &trail{lastIndex: userOneIdx, pieces: []string{url}})
			}
		}
	}

	var longestTrail trail
	for _, commonTrail := range commonTrails {
		if len(commonTrail.pieces) > len(longestTrail.pieces) {
			longestTrail = *commonTrail
		}
	}

	return longestTrail.pieces
}

func main() {
	// Question One
	//displayCounts := []string{
	//	"900,google.com",
	//	"60,mail.yahoo.com",
	//	"10,mobile.sports.yahoo.com",
	//	"40,sports.yahoo.com",
	//	"300,yahoo.com",
	//	"10,stackoverflow.com",
	//	"2,en.wikipedia.org",
	//	"1,es.wikipedia.org",
	//}
	//fmt.Println(numberOfHits(displayCounts))

	// Question Two
	//userOne := []string{"/eight.html", "/four.html", "/six.html", "/seven.html", "/one.html"}
	//userTwo := []string{"/nine.html", "/two.html", "/three.html", "/four.html", "/six.html", "/seven.html"}
	//fmt.Println(longestCommonTrail(userOne, userTwo))
	
}
