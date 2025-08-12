package tree

import (
    "sort"
    "errors"
    "fmt"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	rootAmount := 0
	for i := range records {
		if records[i].ID == 0 && records[i].Parent == 0 {
			rootAmount++
		}
	}

	if rootAmount != 1 {
		return nil, errors.New("no root node")
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	err := validateRecords(records)
	if err != nil {
		return nil, err
	}
	var ans *Node
	for i := 0; i < len(records); i++ {
		ans, err = BuildNode(records[i], ans)
		if err != nil {
			return nil, err
		}
	}

	return ans, nil
}

func BuildNode(r Record, ans *Node) (*Node, error) {
	if r.Parent > r.ID {
		return nil, errors.New("invalid parent")
	}

	if r.ID == 0 && r.Parent == 0 && ans == nil {
		return &Node{ID: r.ID, Children: []*Node{}}, nil
	}

	if r.ID == r.Parent {
		return nil, errors.New("circular reference")
	}

	if r.Parent == ans.ID {
		for i := range ans.Children {
			if ans.Children[i].ID == r.ID {
				return nil, errors.New("duplicate node")
			}
		}

		ans.Children = append(ans.Children, &Node{ID: r.ID, Children: []*Node{}})
		return ans, nil
	}

	for i := 0; i < len(ans.Children); i++ {
		_, err := BuildNode(r, ans.Children[i])
		if err != nil {
			return nil, err
		}
	}
	return ans, nil
}

func validateRecords(records []Record) error {
	n := len(records)
	idMap := make(map[int]bool)
	for _, r := range records {
		idMap[r.ID] = true
	}

	for i := 0; i < n; i++ {
		if !idMap[i] {
			return fmt.Errorf("missing ID %d", i)
		}
	}
	return nil
}
