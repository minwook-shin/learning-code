var arr = [1, 2];
var newArr = [...arr, 3];
console.log(newArr);


function sum(x, y) {
    return x + y;
}
const num = [1, 2];
console.log(sum(...num));

a = {
    hello: "hello",
    world: "world"
}
b = {
    javascript: "javascript",
    ...a
}
console.log(b);