# Facade

Exposing several components through a single interface

## Motivation

- Balancing complexity and presentation/usability
- Typical home
    - Many subsytems (electrical, sanitation)
    - Complex internal structure
    - End user is not exposed to the internals
- Same with software!

The facade design pattern provides a simple, easy to understand/user interface over a large and sophisticated body of code.

## Summary

- Build a facade to provide a simplified API over a set of components
- May wish to (optionally) expose internals through the facade
- May allow users to 'escalate' to use more complex APIs if they need to