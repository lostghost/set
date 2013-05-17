package set

type Set map[interface{}]bool

// creates a new, initially empty set structure.
func New() Set {
	var set Set
	set = make(map[interface{}]bool)
	return set
}

// adds the element item to the set, if it is not present already.
func (s Set) Add(item interface{}) {
	s[item] = true
}

// removes the element item from the set, if it is present.
func (s Set) Remove(item interface{}) {
	delete(s, item)
}

// checks whether the value item is in the set.
func (s Set) Contains(item interface{}) bool {
	if s[item] == true {
		return true
	}
	return false
}

// returns the number of elements in the set.
func (s Set) Size() int {
	return len(s)
}

// TODO: union(S,T): returns the union of sets S and T.
// TODO: intersection(S,T): returns the intersection of sets S and T.
// TODO: difference(S,T): returns the difference of sets S and T.
// TODO: subset(S,T): a predicate that tests whether the set S is a subset of set T.
// TODO: is_empty(S): checks whether the set S is empty.
// TODO: iterate(S): returns a function that returns one more value of S at each call, in some arbitrary order.
// TODO: build(x1,x2,…,xn,): creates a set structure with values x1,x2,…,xn.
// TODO: enumerate(S): returns a list containing the elements of S in some arbitrary order.
// TODO: create_with_capacity(n): creates a new set structure, initially empty but capable of holding up to n elements.
// TODO: capacity(S): returns the maximum number of values that S can hold.
// TODO: pop(S): returns an arbitrary element of S, deleting it from S.
// TODO: map(F,S): returns the set of distinct values resulting from applying function F to each element of S.
// TODO: filter(P,S): returns the subset containing all elements of S that satisfy a given predicate P.
// TODO: clear(S): delete all elements of S.
// TODO: equal(S1, S2): checks whether the two given sets are equal (i.e. contain all and only the same elements).
