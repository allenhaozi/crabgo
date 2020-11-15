package runtime

import (
	"github.com/mohae/deepcopy"
)

type Scheme struct {
	schemeName string
	gvToList   map[GroupVersion][]Object
	gvToMap    map[Index]Object
}

func NewScheme() *Scheme {
	return &Scheme{
		gvToList: make(map[GroupVersion][]Object, 0),
		gvToMap:  make(map[Index]Object, 0),
	}
}

func (s *Scheme) AddSupportType(gv GroupVersion, types ...Object) {
	for _, obj := range types {
		objCopy := deepcopy.Copy(obj)
		s.gvToList[gv] = append(s.gvToList[gv], objCopy.(Object))
	}
}

func (s *Scheme) AddKnownTypeWithIndex(i Index, obj Object) {
	s.gvToMap[i] = obj
}

func (s *Scheme) GetGVToList() map[GroupVersion][]Object {
	return s.gvToList
}

func (s *Scheme) GetObjectByIndex(i Index) Object {
	if obj, ok := s.gvToMap[i]; ok {
		return obj
	} else {
		return nil
	}
}

func (s *Scheme) GetGVToMap() map[Index]Object {
	return s.gvToMap
}
