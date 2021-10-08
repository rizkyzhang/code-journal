"use strict";
const num = 8;
const str = "Ts";
const bool = true;
const x = "Anything";
const nums = [1, 2, 3, 4, 5];
const arr = [1, "two", true];
// Tuple (the types of data have to be the same position as specified)
const tuple = [2, "two", false];
const tuple2 = [
    [1, true],
    [2, false],
];
// Union
const y = 2;
// Enum
var Direction;
(function (Direction) {
    Direction[Direction["Up"] = 1] = "Up";
    Direction[Direction["Down"] = 2] = "Down";
    Direction[Direction["Left"] = 3] = "Left";
    Direction[Direction["Right"] = 4] = "Right";
})(Direction || (Direction = {}));
var Direction2;
(function (Direction2) {
    Direction2["Up"] = "Up";
    Direction2["Down"] = "Down";
    Direction2["Left"] = "Left";
    Direction2["Right"] = "Right";
})(Direction2 || (Direction2 = {}));
const user = {
    name: "John",
    age: 20,
};
// Type assertion
const id = "2";
// const userId = <number>id;
const userId = id;
// Functions
function mul(x, y) {
    return x * y;
}
function output(message) {
    console.log(message);
}
const user2 = {
    name: "John",
    age: 20,
};
// Generic
function isEmptyArray(arr) {
    return arr.length === 0;
}
const result = isEmptyArray([1, 2, 3]);
//# sourceMappingURL=index.js.map