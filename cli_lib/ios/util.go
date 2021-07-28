package iosLib

import "encoding/json"

type Slice2Str struct {
	Items []string `json:"items"`
}

func NewSlice2Str() *Slice2Str {
	return &Slice2Str{}
}

func (s2s *Slice2Str)Add(item string)  {
	s2s.Items = append(s2s.Items,item)
}

func (s2s *Slice2Str)String() string  {
	j,_:=json.Marshal(s2s.Items)

	return string(j)
}


