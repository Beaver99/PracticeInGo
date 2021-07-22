# Design Pattern

[toc]

## Dependency Injection

* [ref](https://stackify.com/dependency-injection/)

* Context

  * You can introduce interfaces to **break the dependencies between higher and lower level classes**. If you do that, both classes depend on the interface and no longer on each other. I explained this approach in great details in my article about the [dependency inversion principle](https://stackify.com/dependency-inversion-principle/).
  * But even if you implement it perfectly, you still keep a dependency on the lower level class. **The interface only decouples the usage of the lower level class but not its instantiation. At some place in your code, you need to *instantiate the implementation of the interface.*** That prevents you from replacing the implementation of the interface with a different one.
  * The goal of the dependency injection technique is to remove this dependency by separating the usage from the creation of the object. This reduces the amount of required boilerplate code and improves flexibility.

* Some Concepts

  * The 4 **roles** in dependency injection -- If you want to use this technique, you need classes that fulfill four basic roles. These are:

    1. The **service** you want to use.
    2. The **client** that uses the service.
    3. An **interface** that’s used by the client and implemented by the service.
    4. The ***injector*** which *creates a service instance and injects it into the client*.
* Implementation in Practice

## Observer



## MonoInstance ##



## Factory ##



## Builder ##



## Decorator ##



