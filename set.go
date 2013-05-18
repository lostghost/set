package set

import "sync"

type Set struct {
	sync.RWMutex
	m map[interface{}]bool
}

type MapFunction func(interface{}) interface{}

// creates a new, initially empty set structure.
func New(items ...interface{}) Set {
	var set Set
	set.m = make(map[interface{}]bool)
	if len(items) > 0 {
		for i := range items {
			set.Add(items[i])
		}
	}
	return set
}

// adds the element item to the set, if it is not present already.
func (s Set) Add(item interface{}) {
	s.Lock()
	s.m[item] = true
	s.Unlock()
}

// removes the element item from the set, if it is present.
func (s Set) Remove(item interface{}) {
	s.Lock()
	delete(s.m, item)
	s.Unlock()
}

// checks whether the value item is in the set.
func (s Set) Contains(item interface{}) bool {
	s.RLock()
	_, exists := s.m[item]
	s.RUnlock()
	if exists {
		return true
	}
	return false
}

// returns the number of elements in the set.
func (s Set) Size() int {
	return len(s.m)
}

// returns a list containing the elements of the set in some arbitrary order.
func (s Set) Enumerate() []interface{} {
	out := []interface{}{}
	s.RLock()
	for i := range s.m {
		out = append(out, i)
	}
	s.RUnlock()
	return out
}

// checks whether the set is empty.
func (s Set) IsEmpty() bool {
	if s.Size() > 0 {
		return false
	}
	return true
}

// delete all elements of the set.
func (s Set) Clear() {
	s.Lock()
	for i := range s.m {
		delete(s.m, i)
	}
	s.Unlock()
}

// returns the set of distinct values resulting from applying function f to each element of the set.
func (s Set) Map(f MapFunction) Set {
	out := New()
	s.RLock()
	for i := range s.m {
		out.Add(f(i))
	}
	s.RUnlock()
	return out
}

// TODO: union(S,T): returns the union of sets S and T.
// TODO: intersection(S,T): returns the intersection of sets S and T.
// TODO: difference(S,T): returns the difference of sets S and T.
// TODO: subset(S,T): a predicate that tests whether the set S is a subset of set T.
// TODO: iterate(S): returns a function that returns one more value of S at each call, in some arbitrary order.
// TODO: build(x1,x2,…,xn,): creates a set structure with values x1,x2,…,xn.
// TODO: create_with_capacity(n): creates a new set structure, initially empty but capable of holding up to n elements.
// TODO: capacity(S): returns the maximum number of values that S can hold.
// TODO: pop(S): returns an arbitrary element of S, deleting it from S.
// TODO: filter(P,S): returns the subset containing all elements of S that satisfy a given predicate P.
// TODO: equal(S1, S2): checks whether the two given sets are equal (i.e. contain all and only the same elements).
