# 📘 Core OOP Topics

### 1. **Classes and Objects**

- Class: blueprint for objects    
- Object: instance of a class    
- Fields/Attributes (state)    
- Methods (behavior)
### 2. **Encapsulation**

- Hiding internal state    
- Providing controlled access through getters/setters    
- Access modifiers: public, private, protected (in Go: exported/unexported)
### 3. **Inheritance**

- Creating new classes from existing ones    
- Reusing code and behavior    
- Types: single, multiple, multilevel inheritance
### 4. **Polymorphism**

- Same interface, different implementations    
- Compile-time (overloading) vs runtime (overriding)    
- In Go: achieved using **interfaces**
### 5. **Abstraction**

- Hiding implementation details    
- Exposing only essential functionality    
- Abstract classes or interfaces
### 6. **Composition vs Aggregation**

- Composition: “has-a” relationship, strong ownership    
- Aggregation: “has-a” relationship, weak ownership    
- Go prefers **composition** over inheritance    
### 7. **Constructors / Initialization**

- Object creation and initialization    
- In Go: factory functions (`NewType()`)
### 8. **Destructors / Resource Management**

- Cleanup when objects are no longer needed    
- In Go: **garbage collection** handles memory    
### 9. **Method Overloading / Overriding**

- Overloading: same method name, different parameters (Go doesn’t support)    
- Overriding: subclass provides new implementation (Go: via embedding + interfaces)
### 10. **Interfaces / Protocols**

- Define a contract of methods    
- Any type satisfying the interface can be used 
### 11. **Operator Overloading**

- Redefining operators for custom types    
- Go **does not support operator overloading**
### 12. **Static vs Instance Members**

- Static: belongs to class (Go: package-level variables/functions)    
- Instance: belongs to object (struct fields/methods)
### 13. **Access Modifiers**

- Public, private, protected    
- Go uses capitalization: **exported (public)** vs **unexported (private)**    
### 14. **Design Principles / Patterns**

- SOLID principles (Single responsibility, Open/Closed, etc.)    
- Common OOP patterns: Factory, Singleton, Observer, Strategy, etc.    
- Go encourages composition over inheritance and interfaces for design patterns
### 15. **Event Handling / Callbacks**

- Observer pattern / event-driven programming    
- Go: channels, function variables, interfaces