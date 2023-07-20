# Mediator

## Motivation

- Components may go in and out of a system at any time
	- Chat room participants
	- Players in an MMORPG
- It makes no sense for them to have direct references to one another
	- Those references may go dead
- Solution: have then all refer to some central component that faciliates communication

The mediator is a design pattern where the component faciliates communication between other components without them necessarily being aware of each other or having direct (reference) access to each other.

## Summary

- Create the mediator and have each object in the system point to it
    - E.g., assign a field in factory function
- Mediator engages in bidirectional communication with its connected components
- Mediator has methods the components an call
- Components have methods that the mediator can call
- Event processing (e.g., Rx) libraries make communication easier to implement