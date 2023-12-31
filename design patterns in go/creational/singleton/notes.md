# Singleton

## Motivation

- For some components it only makes sense to have one in the system
    - Database repository
    - Object factory
- E.g., the construction call is expensive
    - We only do it once
    - We give everyone the same instance
- Want to prevent anyone creating additional copies
- Need to take care of lazy insatiation

A singleton is a component which is instantiated only once.

## Summary

- Lazy one-time initialisation using sync.Once
- Adhere to DIP: depend on interfaces, not concrete types
