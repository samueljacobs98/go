# Visitor

## Motivation

- Need to define a new operation on an entire type hierarchy
    - E.g., given a document model (lists, paragraphs, etc.), we want to add printing funcitonality
- Do not want to keep modifying every type in the hierarchy
- Want to have the new functionality separate (SRP)
- This approach is often used for traversal
    - Alternative to Iterator
    - Hierarchy members help you traverse themselves

The Visitor pattern is a pattern where a component (visitor) is allowed to traverse the entire hierarchy of types. Implemented by propagating a single _Accept()_ method through the entire hierarchy.

## Dispatch

- Which function to call?
- Single dispatch: depends on name of request and type of receiver
- Double dispatch: depends on name of request and type of _two_ receivers (type of visitor, type of element being visited)

## Summary

- Propagate an _Accept(v *Visitor)_ method throughout the entire hierarchy
- Create a visitor with _VisitFoo(f Foo)_, _VisitBar(b Bar)_, ...for each element in the hierearchy
- Each _Accept()_ simply calls _Visitor.VisitXxx(self)