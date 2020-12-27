package utils

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

var preferCount = make(map[string]int)

func InDislike(target string, subStrList []string) (string, bool) {
	for _, s := range subStrList {
		if s != "" && strings.Contains(target, s) {
			return s, true
		}
	}
	return "", false
}

type wordCount struct {
	word  string
	count int
}

func InPrefer(target string, subStrList []string) (string, bool) {
	wordCountList := make([]*wordCount, 0)
	for _, s := range subStrList {
		wordCountList = append(wordCountList, &wordCount{
			word:  s,
			count: preferCount[s],
		})
	}
	sort.Slice(wordCountList, func(i, j int) bool {
		return wordCountList[i].count < wordCountList[j].count
	})

	for _, w := range wordCountList {
		s := w.word
		if s != "" && preferCount[s] < 2 && strings.Contains(target, s) {
			preferCount[s] += 1
			return s, true
		}
	}
	return "", false
}

func GetStartDateAndEndDate() (string, string) {
	t := time.Now()
	sub := 6 - t.Weekday()
	return t.Format("2006-01-02"),
		t.Add(time.Hour * 24 * time.Duration(sub)).Format("2006-01-02")
}

func RandN(n int32)int32{
	rand.Seed(time.Now().Unix())
	return rand.Int31n(n)
}
