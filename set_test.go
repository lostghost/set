package set

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	testSet := New()
	if testSet.Size() != 0 {
		t.Errorf("A new set should have a size of 0")
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
