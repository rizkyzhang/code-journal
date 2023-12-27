- Technically, JavaScript `instanceof` operator checks the prototype chain to see if any constructor in the prototype chain is equal to the given class. That means if you use inheritance, an instance of a subclass is also an instance of the base class. Example below:
```js
class BaseClass {}
class SubClass extends BaseClass {}

const obj = new SubClass();

obj instanceof SubClass; // true
obj instanceof BaseClass; // true
```
