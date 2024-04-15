package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	name string
	team []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

// MatchResult updates the Wins field in League
// We're chaning Wins, shouldn't we use a pointer receiver? It will change
// with a value receiver, because map is a reference type.
func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	switch {
	case score1 > score2:
		l.Wins[team1]++
	case score1 < score2:
		l.Wins[team2]++
	}
	fmt.Println("Match Result:", l.Wins)
}

// Ranking returns an ordered slice of Team in the order of Wins!
func (l League) Ranking() []string {
	teams := l.Teams
	ranking := make([]string, len(teams))
	sort.Slice(teams, func(i, j int) bool {
		team1 := l.Teams[i]
		team2 := l.Teams[j]
		return l.Wins[team1.name] > l.Wins[team2.name]
	})
	for i, v := range teams {
		ranking[i] = v.name
	}
	return ranking
}

type Ranker interface {
	Ranking() []string
}

// RankPrinter writes Result of Ranker into io.Writer
func RankPrinter(r Ranker, w io.Writer) {
	ranking := r.Ranking()
	for idx, teamname := range ranking {
		_, err := w.Write(append([]byte(teamname), '\n'))
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d. team %s \n", idx+1, teamname)
	}
}

func main() {
	sharks := Team{"sharks", []string{"bity bob", "toothy tom"}}
	bears := Team{"bears", []string{"poo bear", "grizzly grisha"}}
	birds := Team{"birds", []string{"flaling falcon", "feist phoenix"}}
	l := League{
		Teams: []Team{sharks, bears, birds},
		Wins:  map[string]int{},
	}
	fmt.Println(sharks)
	fmt.Println(bears)
	fmt.Println(l)
	l.MatchResult(sharks.name, 10, bears.name, 0)
	l.MatchResult(sharks.name, 12, bears.name, 2)
	l.MatchResult(sharks.name, 20, bears.name, 8)
	l.MatchResult(sharks.name, 12, bears.name, 13)
	l.MatchResult(sharks.name, 12, bears.name, 12)
	l.MatchResult(sharks.name, 44, bears.name, 3)
	l.MatchResult(birds.name, 20, bears.name, 0)
	l.MatchResult(birds.name, 20, sharks.name, 0)
	l.MatchResult(birds.name, 20, sharks.name, 4)

	fmt.Println("l result", l)
	fmt.Println("Ranking", l.Ranking())

	fmt.Println("RankPrinter:")
	RankPrinter(l, os.Stdout)
}
