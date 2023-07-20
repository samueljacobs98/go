# Template Method

## Motivation

- Algorithms can be decomposed into common parts + specifics
- Strategy pattern does this through composition
    - High-level algorithm uses an interface
    - Concrete implementations implement the interface
    - We keep a pointer to the interface; provide concrete implementations
- Template Method performs a similar operation but
    - It is typically just a function, not a struct with a reference ot the implementation
    - Can still use interfaces (just like Strategy); or
    - Can be functional (take several functions as parameters)

A Template Method is a skeleton algorithm defined in a function. Function can either use an interface (like Strategy) or can take several functions as arguments.

## Summary

- Very similar to Strategy
- Typical implementation:
    - Define an interface with common operations
    - Make use of those operations inside a function
- Alternative functional approach:
    - Make a function that takes several functions
    - Can pass in functions that capture local state
    - No need for either structs or interfaces!