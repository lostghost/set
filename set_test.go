package set

import (
	"fmt"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	testSet := New()
	if testSet.Size() != 0 {
		t.Errorf("A new set should have a size of 0")
	}

	testSet = New("one", "two", "three", 1, 2, 3)
	if testSet.Size() != 6 {
		t.Errorf("A new set initialized with 6 items should have a size of 6")
	}

	items := []interface{}{"english", "spanish", "italian", "french"}
	testSet = New(items...)
	if testSet.Size() != len(items) {
		t.Errorf("A new set initialized with an array of items should have a size equal to the len of the array")
	}
}

func TestAdd(t *testing.T) {
	testSet := New()
	testSet.Add("one")
	if testSet.Size() != 1 {
		t.Errorf("Adding one string item to a new set should have a size of 1")
	}

	testSet.Add(4)
	if testSet.Size() != 2 {
		t.Errorf("Adding an int item should increase size to 2")
	}

	testSet.Add("three")
	if testSet.Size() != 3 {
		t.Errorf("Adding another unique string item should increase size to 3")
	}

	testSet.Add("one")
	if testSet.Size() != 3 {
		t.Errorf("Adding a duplicate item should not increase the size")
	}

	testSet = New()
	for i := 0; i < 100; i++ {
		testSet.Add(i)
	}
	if testSet.Size() != 100 {
		t.Errorf("Adding 100 int items to a new set should have a size of 100")
	}

	testSet = New()
	for i := 0; i < 100; i++ {
		testSet.Add(fmt.Sprintf("item_num_%d", i))
	}
	if testSet.Size() != 100 {
		t.Errorf("Adding 100 string items to a new set should have a size of 100")
	}
}

func TestRemove(t *testing.T) {
	testSet := New()
	testSet.Add("one")
	testSet.Add("two")

	sampleSize := testSet.Size()
	testSet.Remove("one")
	if testSet.Size() != sampleSize-1 {
		t.Errorf("Removing an item from a set should decrease the size by 1")
	}

	sampleSize = testSet.Size()
	testSet.Remove("unknown")
	if testSet.Size() != sampleSize {
		t.Errorf("Removing an item that is not in a set should not change the set size")
	}

	testSet = New()
	for i := 0; i < 100; i++ {
		testSet.Add(i)
	}
	for i := 0; i < 100; i++ {
		testSet.Remove(i)
	}
	if testSet.Size() != 0 {
		t.Errorf("Adding 100 items and them removing them should have a set size 0")
	}
}

func TestContains(t *testing.T) {
	testSet := New()
	if testSet.Contains("nothing") {
		t.Errorf("A new set should not contain an item")
	}

	testSet.Add("English")
	if testSet.Contains("English") == false {
		t.Errorf("After adding an item, the set should contain that item")
	}
	if testSet.Contains("english") {
		t.Errorf("String items in the set should be case sensitive")
	}

	testSet.Add(6453)
	if testSet.Contains(6453) == false {
		t.Errorf("After adding an int item, the set should contain that item")
	}

	if testSet.Contains("arbitrary item") {
		t.Errorf("After adding items, a set should not contain an arbitrary string item")
	}

	if testSet.Contains(29543) {
		t.Errorf("After adding items, a set should not contain an arbitrary int item")
	}
}

func TestEnumerate(t *testing.T) {
	items := []interface{}{"english", "spanish", 5, 8, "french"}
	testSet := New()
	for i := range items {
		testSet.Add(items[i])
	}
	enum := testSet.Enumerate()
	for i := range items {
		found := false
		for j := range enum {
			if items[i] == enum[j] {
				found = true
				break
			}
		}
		if found == false {
			t.Errorf("Enumerate should contains the elements that were added")
		}
	}
	for i := range enum {
		found := false
		for j := range items {
			if enum[i] == items[j] {
				found = true
				break
			}
		}
		if found == false {
			t.Errorf("Enumerate shouldn't contain elements that were not added")
		}
	}
}

func TestIsEmpty(t *testing.T) {
	testSet := New()
	if testSet.IsEmpty() == false {
		t.Errorf("A new set should be empty")
	}

	testSet.Add("one")
	if testSet.IsEmpty() {
		t.Errorf("A set after adding an item should not be empty")
	}

	testSet.Remove("one")
	if testSet.IsEmpty() == false {
		t.Errorf("A set after removing the only item should be empty")
	}

	testSet = New()
	items := []interface{}{"english", "spanish", 5, 8, "french"}
	for i := range items {
		testSet.Add(items[i])
	}
	if testSet.IsEmpty() {
		t.Errorf("A set after adding several items should not be empty")
	}

	for i := range items {
		testSet.Remove(items[i])
	}
	if testSet.IsEmpty() == false {
		t.Errorf("A set after removing all of the items that were added should be empty")
	}
}

func TestClear(t *testing.T) {
	testSet := New()
	items := []interface{}{"english", "spanish", 5, 8, "french"}
	for i := range items {
		testSet.Add(items[i])
	}
	testSet.Clear()
	if testSet.Contains("spanish") {
		t.Errorf("After clearing a set, it should not contain an item that was added")
	}
	if testSet.IsEmpty() == false {
		t.Errorf("After clearing a set, the set should be empty")
	}
}

func TestMap(t *testing.T) {
	testSet := New()
	items := []int{1, 2, 3, 4}
	for i := range items {
		testSet.Add(items[i])
	}
	mappedSet := testSet.Map(func(item interface{}) interface{} {
		return item.(int) * 2
	})
	for i := range items {
		if mappedSet.Contains(items[i]*2) == false {
			t.Errorf("The result of the int * 2 mapping function should contain the value %d", items[i]*2)
		}
	}
}

func TestFilter(t *testing.T) {
	testSet := New("ken", "brad", "jeff", "ryan", "tim", "greg", "scott")

	filterFunc := func(i interface{}) bool {
		if len(i.(string)) == 4 {
			return true
		}
		return false
	}

	includedItems := []interface{}{"brad", "jeff", "ryan", "greg"}
	excludedItems := []interface{}{"ken", "tim", "scott"}

	filterSet := testSet.Filter(filterFunc)
	if filterSet.Size() != len(includedItems) {
		t.Errorf("There should be %d elements in the filtered set (%d found)", len(includedItems), filterSet.Size())
	}

	for i := range includedItems {
		if filterSet.Contains(includedItems[i]) == false {
			t.Errorf("All four letter items should be inluded in the filtered set (%v missing)", includedItems[i])
		}
	}
	for i := range excludedItems {
		if filterSet.Contains(excludedItems[i]) {
			t.Errorf("Only four letter items should be included in the filtered set (%v unexpected)", excludedItems[i])
		}
	}
}

func TestPop(t *testing.T) {
	items := []interface{}{"english", "spanish", "french", "italian"}
	testSet := New(items...)
	setSize := testSet.Size()
	item := testSet.Pop()

	if testSet.Size() != setSize-1 {
		t.Errorf("Pop on a non-zero size set should reduce the set size by one")
	}

	if testSet.Contains(item) {
		t.Errorf("The item returned by Pop should no longer be contained in the set")
	}

	found := false
	for i := range items {
		if item == items[i] {
			found = true
			break
		}
	}
	if found == false {
		t.Errorf("The result of Pop should have been one of the items that was added to the set")
	}

	testSet = New()
	item = testSet.Pop()
	if item != nil {
		t.Errorf("The expected result of Pop on an empty set is nil, the actual result was %v", item)
	}

	for i := 0; i < 100; i++ {
		testSet.Add(i)
	}
	for i := 0; i < 100; i++ {
		_ = testSet.Pop()
	}
	if testSet.IsEmpty() == false {
		t.Errorf("Popping the same number of unique items that were added should result in an empty set")
	}
}

func TestSubset(t *testing.T) {
	testSet := New("one", "two", "three", 4, 5, 6)
	superSet := New("one", "two", "three", "four", "five", "six", 1, 2, 3, 4, 5, 6)
	altSet := New(1, 2, 3, 4, 5, 6)

	if testSet.Subset(superSet) == false {
		t.Errorf("testSet should be a subset of superSet")
	}

	if testSet.Subset(altSet) {
		t.Errorf("testSet is not a subset of altSet")
	}
}

func TestUnion(t *testing.T) {
	groupOne := []interface{}{"red", "blue", "green"}
	groupTwo := []interface{}{"red", "yellow", "purple", "black"}

	setOne := New(groupOne...)
	setTwo := New(groupTwo...)
	setOneSize := setOne.Size()
	setTwoSize := setTwo.Size()

	unionSet := Union(setOne, setOne)
	if unionSet.Size() != setOne.Size() {
		t.Errorf("The size of the union of two identical sets should be the size of the original set")
	}

	unionSet = Union(setOne, setTwo)
	if setOne.Size() != setOneSize {
		t.Errorf("Union should not affect the size of the first set")
	}
	if setTwo.Size() != setTwoSize {
		t.Errorf("Union should not affect the size of the second set")
	}

	if unionSet.Size() != (setOne.Size() + setTwo.Size() - 1) {
		t.Errorf("The size of the union of two sets with one common item should be the sum of the sizes of the two sets minus 1")
	}

	for i := range groupOne {
		if unionSet.Contains(groupOne[i]) == false {
			t.Errorf("The union of two sets should contain all of the elements of the first set, missing %v", groupOne[i])
		}
	}
	for i := range groupTwo {
		if unionSet.Contains(groupTwo[i]) == false {
			t.Errorf("The union of two sets should contain all of the elements of the second set, missing %v", groupTwo[i])
		}
	}
}

func TestIntersection(t *testing.T) {
	groupOne := []interface{}{"red", "blue", "green"}
	groupTwo := []interface{}{"red", "yellow", "purple", "black", "blue"}

	groupIntersection := []interface{}{"red", "blue"}
	groupDifference := []interface{}{"green", "yellow", "purple", "black"}

	setOne := New(groupOne...)
	setTwo := New(groupTwo...)
	setOneSize := setOne.Size()
	setTwoSize := setTwo.Size()

	intersectionSet := Intersection(setOne, setOne)
	if intersectionSet.Size() != setOne.Size() {
		t.Errorf("The size of the intersection of two identical sets should be the size of the original set")
	}

	intersectionSet = Intersection(setOne, setTwo)
	if setOne.Size() != setOneSize {
		t.Errorf("Intersection should not affect the size of the first set")
	}
	if setTwo.Size() != setTwoSize {
		t.Errorf("Intersection should not affect the size of the second set")
	}

	if intersectionSet.Size() != len(groupIntersection) {
		t.Errorf("The size of the intersection of two sets with %d common items should be %d", len(groupIntersection), len(groupIntersection))
	}

	for i := range groupIntersection {
		if intersectionSet.Contains(groupIntersection[i] == false) {
			t.Errorf("The intersection should contain all of the common items: missing %v", groupIntersection[i])
		}
	}

	for i := range groupDifference {
		if intersectionSet.Contains(groupDifference[i]) {
			t.Errorf("The intersection should not contain any of the items that are not common to both sets: not expecting %v", groupDifference[i])
		}
	}
}

func TestDifference(t *testing.T) {
	groupOne := []interface{}{"red", "blue", "green"}
	groupTwo := []interface{}{"red", "yellow", "purple", "black", "blue"}

	groupIntersection := []interface{}{"red", "blue"}
	groupDifference := []interface{}{"green", "yellow", "purple", "black"}

	setOne := New(groupOne...)
	setTwo := New(groupTwo...)
	setOneSize := setOne.Size()
	setTwoSize := setTwo.Size()

	differenceSet := Difference(setOne, setOne)
	if differenceSet.Size() != 0 {
		t.Errorf("The size of the difference of two identical sets should be 0")
	}

	differenceSet = Difference(setOne, setTwo)
	if setOne.Size() != setOneSize {
		t.Errorf("Difference should not affect the size of the first set")
	}
	if setTwo.Size() != setTwoSize {
		t.Errorf("Difference should not affect the size of the second set")
	}

	if differenceSet.Size() != len(groupDifference) {
		t.Errorf("The size of the difference of two sets with %d diffent items should be %d", len(groupDifference), len(groupDifference))
	}

	for i := range groupDifference {
		if differenceSet.Contains(groupDifference[i] == false) {
			t.Errorf("The difference should contain all of the different items: missing %v", groupDifference[i])
		}
	}

	for i := range groupIntersection {
		if differenceSet.Contains(groupIntersection[i]) {
			t.Errorf("The difference should not contain any of the items that are common to both sets: not expecting %v", groupIntersection[i])
		}
	}
}

func ExampleSet() {
	// Creating a new empty set
	set := New()

	// Add items to the set
	set.Add("blue")
	set.Add("red")
	set.Add("green")

	// How many items does the set contain?
	log.Println(set.Size()) // 3

	// Adding a duplicate item doesn't change the set size
	set.Add("blue")
	log.Println(set.Size()) // 3

	// Test to see if a set contains a specific item
	log.Println(set.Contains("blue")) // true

	// Remove an item from the set
	set.Remove("blue")
	log.Println(set.Contains("blue")) // false

	// Create a set with initial values
	set = New(1, 2, 3, 5, 7, 11)

	// Find the intersection of two sets
	multiplesOf2 := New(2, 4, 6, 8, 10, 12, 14, 16, 18, 20)
	multiplesOf3 := New(3, 6, 9, 12, 15, 18, 21)
	multiplesOf2And3 := Union(multiplesOf2, multiplesOf3)
	log.Println(multiplesOf2And3.Enumerate()) // {6, 12, 18}
}

func ExampleNew() {
	// Create an empty set
	emptySet := New()
	log.Println(emptySet.IsEmpty()) // true

	// Create a set with initial values
	colors := New("red", "blue", "green")
	log.Println(colors.Size()) // 3

	// Create a set with initial values from a slice of interfaces
	slice := []interface{}{1, 2, 3, 4, "one", "two", "three", "four"}
	numbers := New(slice...)
	log.Println(numbers.Size()) // 8
}
