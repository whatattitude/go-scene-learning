package main

import (
	"github.com/whatattitude/go-kit-example/httpserver/library/fp"
)

type structArray struct {
	Value     int
	ValueArr  []int
	ArrayItem []structArrayItem
}

type structArrayItem struct {
	Value    int
	ValueArr []int
}

func deepCopy(s *structArray) *structArray {
	var resp structArray
	resp = *s
	resp.ArrayItem = make([]structArrayItem, len(s.ArrayItem))
	copy(resp.ArrayItem, s.ArrayItem)

	for i := range resp.ArrayItem {
		resp.ArrayItem[i].ValueArr = make([]int, len(s.ArrayItem[i].ValueArr))
		copy(resp.ArrayItem[i].ValueArr, s.ArrayItem[i].ValueArr)
	}

	return &resp
}

func (s *structArray) deepAdd() {
	s.Value++
	for i := range s.ValueArr {
		s.ValueArr[i]++
	}
	for i := range s.ArrayItem {
		s.ArrayItem[i].Value++
		for j := range s.ArrayItem[i].ValueArr {
			s.ArrayItem[i].ValueArr[j]++
		}

	}

}

func addByRuntineAndRange(s *structArray) *structArray {
	resp := deepCopy(s)
	s.ArrayItem[0].ValueArr[0]++
	s.ArrayItem[0].ValueArr[0]++
	fp.Fp(resp)
	fp.Fp(s)
	return s
}
