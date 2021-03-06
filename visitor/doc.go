// Package visitor is visitor design pattern in Go.
//
// 访问者模式是一种将算法与对象结构分离的软件设计模式。
// 这个模式的基本想法如下：首先我们拥有一个由许多对象构成的对象结构，这些对象的类都拥有一个accept方法用来接受访问者对象；
// 访问者是一个接口，它拥有一个visit方法，这个方法对访问到的对象结构中不同类型的元素作出不同的反应；
// 在对象结构的一次访问过程中，我们遍历整个对象结构，对每一个元素都实施accept方法，在每一个元素的accept方法中回调访问者的visit方法，
// 从而使访问者得以处理对象结构的每一个元素。我们可以针对对象结构设计不同的实在的访问者类来完成不同的操作。
// 访问者模式使得我们可以在传统的单分派语言（如Smalltalk、Java和C++）中模拟双分派技术。
// 对于支持多分派的语言（如CLOS），访问者模式已经内置于语言特性之中了，从而不再重要。
//
// For a full chinese explanation visit https://zh.wikipedia.org/wiki/%E8%AE%BF%E9%97%AE%E8%80%85%E6%A8%A1%E5%BC%8F
//
// In object-oriented programming and software engineering,
// the visitor design pattern is a way of separating an algorithm from an object structure on which it operates.
// A practical result of this separation is the ability to add new operations to existing object structures
// without modifying those structures. It is one way to follow the open/closed principle.

// In essence, the visitor allows one to add new virtual functions to a family of classes without modifying
// the classes themselves; instead, one creates a visitor class that implements all of the appropriate specializations
// of the virtual function. The visitor takes the instance reference as input, and implements the goal through double dispatch.

// For a full explanation visit https://en.wikipedia.org/wiki/Visitor_pattern

package visitor
