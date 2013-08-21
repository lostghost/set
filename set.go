// The set package implements a set collection -- a store a unique values without order.
package set

type Set map[interface{}]struct{}

type MapFunction func(interface{}) interface{}

type FilterFunction func(interface{}) bool

// Creates a new set structure. The New constructor can take a variadic paramater to initialize the set with an initial set of values.
func New(items ...interface{}) Set {
	o := make(Set)
	if len(items) > 0 {
		for i := range items {
			o.Add(items[i])
		}
	}
	return o
}

// Adds the element item to the set, if it is not present already.
func (s Set) Add(item interface{}) {
	s[item] = struct{}{}
}

// Removes the element item from the set, if it is present.
func (s Set) Remove(item interface{}) {
	delete(s, item)
}

// Checks whether the value item is in the set.
func (s Set) Contains(item interface{}) bool {
	_, exists := s[item]
	return exists
}

// Returns the number of elements in the set.
func (s Set) Size() int {
	return len(s)
}

// Returns a list containing the elements of the set in some arbitrary order.
func (s Set) Enumerate() []interface{} {
	o := []interface{}{}
	for i := range s {
		o = append(o, i)
	}
	return o
}

// Checks whether the set is empty.
func (s Set) IsEmpty() bool {
	if s.Size() > 0 {
		return false
	}
	return true
}

// Delete all elements of the set.
func (s Set) Clear() {
	for i := range s {
		delete(s, i)
	}
}

// Returns the set of distinct values resulting from applying function f to each element of the set.
func (s Set) Map(f MapFunction) Set {
	o := New()
	for i := range s {
		o.Add(f(i))
	}
	return o
}

// Returns the subset containing all elements of the set that satisfy a given predicate f.
func (s Set) Filter(f FilterFunction) Set {
	o := New()
	for i := range s {
		if f(i) {
			o.Add(i)
		}
	}
	return o
}

// Returns an arbitrary element of the set, deleting it from the set
func (s Set) Pop() interface{} {
	var o interface{}
	for o = range s {
		break
	}
	s.Remove(o)
	return o
}

// A predicate that tests whether the set is a subset of set t.
func (s Set) Subset(t Set) bool {
	o := true
	for i := range s {
		if t.Contains(i) == false {
			o = false
			break
		}
	}
	return o
}

// Returns the union of sets s and t.
func Union(s, t Set) Set {
	o := New(s.Enumerate()...)
	for i := range t {
		o.Add(i)
	}
	return o
}

// Returns the intersection of sets s and t.
func Intersection(s, t Set) Set {
	o := New()
	for i := range t {
		if s.Contains(i) {
			o.Add(i)
		}
	}
	return o
}

// Returns the difference of sets s and t.
func Difference(s, t Set) Set {
	o := New()
	for i := range t {
		if s.Contains(i) == false {
			o.Add(i)
		}
	}
	for i := range s {
		if t.Contains(i) == false {
			o.Add(i)
		}
	}
	return o
}

// TODO: iterate(S): returns a function that returns one more value of S at each call, in some arbitrary order.
// TODO: create_with_capacity(n): creates a new set structure, initially empty but capable of holding up to n elements.
// TODO: capacity(S): returns the maximum number of values that S can hold.
// TODO: equal(S1, S2): checks whether the two given sets are equal (i.e. contain all and only the same elements).
