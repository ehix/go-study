# OOP in Go:

[Video](https://www.youtube.com/watch?v=jexEpE7Yv2A&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=17)

Essential elements of OOP have been:
- abstraction
- encapsulation
- polymorphism
- inheritance

Sometimes those last two items are combined or confused.
Go's appoach is similar, but different.

### Abstraction
#### Decoupling behaviour from implementation details

The Unix file sustem API is a great example.
Roughly five basic functions hide all the details.
- open, close, read, write, ioctl

Many different OS things can be treated like files.

### Encapsulation
#### Hiding implementation details from misuse

Hiding infomation, the user doesn't know or depend on the details.

It's hard to maintain an abstraction if the details are exposed:
- the internals may be manipulated in ways contrary to the concept behind the abstraction.
- user of the abstraction may come to depend on the internal details, but those might change.

Encapsulation usually means controlling the visablility of names, `private` variables.

### Polymorphism
#### Multiple types behind a single interface
<i>Many shapes (in greek...)</i>

Go does through interfaces not subtypes

shape class, subclasses circle, square, triangle.

Three main types:
1. ad-hoc: function/operator overloading.
2. parametric: commonly known as "generic programming", generics.
3. subtype: subclasses substituting for superclasses, specialisation.

"Protocol-oriented" programming users explicit interface types, now supported in many popular languages (an ad-hoc method).
In this case, behaviour is completely separate from implementation, which is good for abstraction.

### Inheritance

Go doesn't have inheritance, because there's no classes.

Has conflicting meanings:
- substitution (subtype) polymorphism
- structual sharing of implementation details. Most subclasses are allowed to see the internals of the superclass

In theory, inheritance should always imply subtyping, the subclass should be a "kind of" superclass.

### Why would inheritance be bad?
It injects a dependence on the superclass into the subclass:
- what if the super changes
- what if the abstract concept is leaky?
e.g shape -> shapes.area() -> line (makes shapes but doesn't have an area).

Not having inheritance means better encap and isolation. Prefer composition to inheritance.

<i>"Interfaces will force you to think in terms of communication between objects".</i>
