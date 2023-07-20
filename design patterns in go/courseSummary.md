# Curse Summary

## Creational

- Builder
    - Separate component for when object construction gets too complicated
    - Can create mutually cooperating sub-builders
    - Often has a fluent interface
- Factories
    - Factory functions (constructors) are common
    - Factory can be simple function or dedicated struct
- Prototype
    - Creation of an object from an existing object
    - Requires either explicit deep copy or copy through serialisation
- Singleton
    - When you need to ensure just a single instance exists
    - Can be made thread-safe and lazy
    - Consider extracting interface or using dependency injection

## Structural

- Adapter
    - Converts the interface you get to the interface you need
- Bridge
    - Decouple abstraction from implementation
- Composite
    - Allows clients to treat individual objects and compositions of objects uniformly
- Decorator
    - Attach additional responsibilites to objects
    - Can be done through embedding or pointers
- Facade
    - Provide a single unified interface over a set of interfaces
- Flyweight
    - Efficiently support very large numbers of similar objects
- Proxy
    - Provide a surrogate object that forwards calls to the real object while performing additional functions
    - E.g., access control, communication, logging, etc.

## Behavioural

- Chain of Responsibility
    - Allow components to process information/events in a chain
    - Each element in the chain refers to next element; or
    - Make a list and go through it
- Command
    - Encapsulate a request into a separate object
    - Good for audit, replay, undo/redo
    - Part of CQS/CQRS
- Interpreter
    - Transform textual input into structures (e.g.ASTs)
    - Used by interpreters, compilers, static analysis tools, etc.
    - _Compiler Theory_ is a separate branch of Computer Science
- Iterator
    - Provides an interface for accessing elements of an aggregate object
- Mediator
    - Provides mediation services between several objects
    - E.g., message passing, chat room
- Memento
    - Yields tokens representing system states
    - Tokens do not allow direct manipulation, but can be used in appropriate APIs
- Observer
    - Allows notifications of changes/happenings in a components
- State
    - We model systems by having one of a possible states and transitions between these states
    - Such a system is called a _state machine_
    - Special frameworks exist to orchestrate state machines
- Strategy and Template Method
    - Both define a skeleton algorithm with details filled in by implementer
    - Strategy uses composition; Template Method doesn't
- Visitor
    - Allows non-intrusive addition of functionality to hierarchies