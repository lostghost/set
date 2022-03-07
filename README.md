# set - golang set collection

This is a very basic implementation of a set collection type in golang

http://en.wikipedia.org/wiki/Set_(computer_science)

"In computer science, a set is an abstract data structure that can store certain values, without any particular order, and no repeated values. It is a computer implementation of the mathematical concept of a finite set. Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests a value for membership in a set."

This implementation supports initializing empty sets and sets with an initial collection. It supports __adding__ and __removing__ items. You can test whether a set __contains__ a specific value, whether the set __is empty__, or whether the set is a __subset__ of another set. You can check for the set __size__, __pop__ items from the set, __clear__ the content of the set, and return an __enumerator__ of the set contents. You can __map__ a function against the set contents and __filter__ a set against a filter function. Additionally, you can get the __union__, __intersection__ and __difference__ of two sets.

## USAGE

```go
package main

import(
    "github.com/shopsmart/set"
    "log"
)

func main() {
    // Initialize an empty set
    mySet := set.New()

    // Initialize a set with a collection of values
    colors := set.New("red", "blue", "green", "yellow")

    // Add items
    colors.Add("purple")

    // Check to see if an item is in the set
    log.Println(colors.Contains("red")) // true

    // Check the size of the set
    log.Println(colors.Size()) // 5

    // Distinct values will only be added one time
    colors.Add("purple")
    log.Println(colors.Size()) // Still 5

    // Check to see if a set is empty
    log.Println(mySet.IsEmpty()) // true
    log.Println(mySet.Size()) // 0
    log.Println(colors.IsEmpty()) // false
    log.Println(colors.Size()) // 5

    // Remove items from the set
    colors.Remove("red")
    log.Println(colors.Contains("red")) // false

    // Clear a map
    colors.Clear()
    log.Println(colors.IsEmpty()) // true
    log.Println(colors.Size()) // 0
}
```
