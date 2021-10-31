package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Node struct {
	Children []*Node
	ID       int
}

type Record struct {
	Parent, ID int
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil,nil
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i,v := range records {
		if i != v.ID {
			return nil, errors.New("error")
		}
	}

	err2 := check(records)
	if err2 != nil{
		return nil, err2
	}

	root, err := build(records, nil)
	if err != nil {
		return nil, err
	}
	if len(records) != checkFromRoot(root) +1 {
		return nil, errors.New("error")
	}
	return root, nil
}

func checkFromRoot(root *Node) int  {
	sum := 0
	for _,val := range root.Children {
		sum += checkFromRoot(val) +1
	}
	return sum
}

func check(records []Record) error {
	for _, y := range records{
		fmt.Println(y)
		if count(records, y) > 1 {
			return errors.New("duplicate node")
		}
	}
	return nil
}

func count(records []Record, record Record) int {
	count := 0
	for _,r := range records {
		if r == record {
			count++
		}
	}
	return count
}

func build(records []Record, root *Node) (*Node, error) {
	for _, record := range records {
		if record.ID == 0 && root == nil {
			if record.ID != 0 || record.Parent != 0 {
				return nil, errors.New("No roots parent")
			}
			var child []*Node
			return build(records, &Node{child, record.ID})
		} else if root != nil && record.Parent == root.ID && record.ID != 0 {
			if record.ID < root.ID {
				return nil, errors.New("higher ")
			}
			var child []*Node
			nRoot, err := build(records, &Node{child, record.ID})
			if err != nil {
				return nil, err
			}
			root.Children = append(root.Children, nRoot)
		}
	}
	if root == nil {
		return nil, errors.New("no root node")
	}
	return root, nil
}
