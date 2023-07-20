# Iterator

## Motivation

- Iteration (traversal) is a core functionality of various data structures
- An _iterator_ is a type that faciliates the traversal
    - Keeps a pointer to the current element
    - knows how to move to a different element
- Go allows iteration with _range_
    - Built-in support in many objects (arrays, slices, etc.)
    - Can be supported in our own structs

The Iterator is a design pattern for an object that facilitates the traversal of a data structure.

## Summary

- An iterator specifies how you can traverse an object
- Moves along the iterated collection, indicating when last element has been reached
- Not idiomatic in Go (no standard Iterable interface)