package kindergarten

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

import (
	"fmt"
	"strings"
    "sort"
    "regexp"
)

type Garden struct {
	FirstRow    string
	SecondRow   string
	MapChildren map[string]int
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	sd := strings.Split(diagram, "\n")
	if len(sd) != 3 {
		return nil, fmt.Errorf("invalid diagram")
	}

	if regexp.MustCompile(`[a-z]`).MatchString(sd[1]) {
		return nil, fmt.Errorf("lower case")
	}

	if len(sd[1])%2 != 0 {
		return nil, fmt.Errorf("diagram length is odd")
	}

	if len(sd[1]) != len(sd[2]) {
		return nil, fmt.Errorf("different row length")
	}

	sortedChildren := append([]string{}, children...)
	fmt.Println("sortChildren", sortedChildren)

	mapChildren := make(map[string]int)
	sort.Slice(sortedChildren, func(i, j int) bool {
		return sortedChildren[i] < sortedChildren[j]
	})
	for i, c := range sortedChildren {
		if _, exist := mapChildren[c]; exist {
			return nil, fmt.Errorf("duplicate child: %s", c)
		}
		mapChildren[c] = i * 2
	}

	return &Garden{
		FirstRow:    strings.ReplaceAll(sd[1], " ", ""),
		SecondRow:   strings.ReplaceAll(sd[2], " ", ""),
		MapChildren: mapChildren,
	}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	num, ok := g.MapChildren[child]
	if !ok {
		return []string{}, false
	}

	ownPlanets := []string{}
	for i := num; i <= num+1; i++ {
		ownPlanet, err := g.ToPlanets(string(g.FirstRow[i]))
		if err != nil {
			return []string{}, false
		}
		ownPlanets = append(ownPlanets, ownPlanet)
	}

	for i := num; i <= num+1; i++ {
		ownPlanet, err := g.ToPlanets(string(g.SecondRow[i]))
		if err != nil {
			return []string{}, false
		}
		ownPlanets = append(ownPlanets, ownPlanet)
	}
	return ownPlanets, true
}

func (g *Garden) ToPlanets(s string) (string, error) {
	switch s {
	case "G":
		return "grass", nil
	case "C":
		return "clover", nil
	case "R":
		return "radishes", nil
	case "V":
		return "violets", nil
	default:
		return "", fmt.Errorf("invalid diagram")
	}
}



