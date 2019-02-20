/*jshint esversion: 6 */

console.log("hello,world!");

// var
var a = 10;
console.log("a : " + a);
var a = 100;
console.log("a : " + a);

// let
let b = 20;
console.log("b : " + b);
b = 200; // 'b' has already been declared.
console.log("b : " + b);

//const
const c = 30;
console.log("c : " + c);
// c = 300; Attempting to override 'c' which is a constant.
const arr = [];
arr.push(1);
console.log("arr :" + arr);

//var scope
var total = 0;
for (var i = 0; i < 10; i++) {
    total++;
}
console.log("total : " + total);
console.log("i : " + i);

//let scope
let realTotal = 0;
for (let j = 0; j < 10; j++) {
    realTotal++;
    console.log("j : " + j);
}
// console.log("j : " + j);

//for of
let iterable = [3, 5, 7];
iterable.foo = "hello";
for (let i of iterable) {
    console.log(i);
}

//hoisting
function same() {
    return 10;
}
var func = same();
console.log("func : " + func);

function same() {
    return 20;
}

// Arrows
let f = () => {
    return 404;
};
let list = [1, 2, 3, 4, 5];
list.forEach((value) => {
    console.log("value: " + value);
});

// Enhanced Object Literal
let dict = {
    words: 100,
    print() {
        console.log("hello, world!");
    }
};
dict.print();

// Template Strings
let name = "Bob";
let hello = `Hello ${name}, 
how are you today? `;

//Generators
function* gen() {
    yield*["a", "b", "c"];
}
var a = gen();
console.log(a.next()); // { value: "a", done: false }
console.log(a.next()); // { value: "b", done: false }
console.log(a.next()); // { value: "c", done: false }
console.log(a.next()); // { value: undefined, done: true }

//Modules
//Node.js doesn't support ES6's import yet.

//Classes
var Foo = class {
    constructor() {}
    bar() {
        return "Hello World!";
    }
};
var instance = new Foo();
instance.bar(); // "Hello World!"

//Destructuring
let [first, second] = [1, 2];
console.log(first);

// proxy
var p = new Proxy({}, {
    get: function (target, name) {
        return name in target ?
            target[name] :
            404;
    },
    set: function (obj, prop, value) {
        if (prop === 'score') {
            if (value > 100) {
                throw new RangeError();
            }
        }
        obj[prop] = value;
    }
});
p.score = '100';
console.log(p.score);

// Promise
Promise.resolve(console.log("start")).then(() => {
    console.log("end");
});