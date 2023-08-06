## Static

- We can access private property and method with `ReflectionProperty`
- Static method/function can't access `$this`
- Static methods is usually used to create utility library

## Inheritance

- Child class can't access parent's private property/method, only public and protected is accessible
- Child class can't reduce the visibility of a parent's property/method, for
  example `public -> private` `public -> protected`, while the opposite is possible `protected -> public` with exception
  to private
- Child class can only override parent's method only with method of same signature (same parameters and return type)
  with exception to constructor
- `parent::constructor`
- `final` keyword can be used to prevent class inheritance and method overriding
- Inheritance creates tight coupling between parent and child class
- Inheritance is not always a good solution, if you find yourself often overriding parent's methods and throwing error
  because you don't want it to be accessed, you might want to consider composition
- Inheritance define **is a** relationship between class, while composition define **has a** or **part of** relationship

## Abstract class & Interface

- An abstract class cannot be instantiated. It provides an interface for other classes to extend.
- An abstract method does not have an implementation. If a class contains one or more abstract methods, it must be an
  abstract class.
- A class that extends an abstract class needs to implement the abstract methods of the abstract class.
- Concrete class is a class that implements an interface
- Abstract class
  - Can contain method implementations
  - Can contain properties
  - Can have private and protected methods
  - Can only extend a single class
- Interface
  - Can only contain method declaration
  - Can only contain methods and constants
  - Can only have public methods
  - Can implement multiple interfaces

## Trait

- Methods with same name in parent class will be replaced by trait methods used in a child class.
- Methods conflict between two trait used in a class can be resolved with `insteadof` or `as` keywords.
- `as` keyword can be used to change method's visibility in trait.
