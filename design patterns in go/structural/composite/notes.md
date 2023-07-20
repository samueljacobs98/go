# Composite

Treating individual and aggregate objects uniformly

## Motivation

- Objects use other objects' fields/methods through embedding
- Composition lets us make compound objects
    - E.g., a mathematical expression composed of simple expressions; or
    - A shape group made of several different shapes
- Composite design pattern is used to treat both single (scalar) and composite objects uniformly
    - I.e., Foo and []Foo have common APIs

Composite design pattern is a mechanism for treating individual (Scalar) objects and compositions of objects in a uniform manner.

## Summary

- Objects can use other objects via composition
- Some composed and singular objects need similar/identical behaviours
- Composite design pattern lets us treat both types of objects uniformly
- Iteration supported with the Iterator design pattern