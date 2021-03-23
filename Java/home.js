console.log("hello everyone");

//alert("you win");

// how to write a comment inline

// var b ="rezaul karim";
// console.log(b);

// var number = 45;
// console.log(45);


// var b = prompt("What is Your name");
// document.getElementById("sometext").innerHTML= b;

//number in javascript
var num1 = 10;
// Increment num1 by1
num1++;
// decrement nub by 1
num1--;
console.log(num1); 

// devide /, multiply * , remainder %
console.log(num1 % 6);

// increment by decrement by any nmber i want

num1 +=10;
console.log(num1);

// function baby
    // 1. create a function
    // 2. call the function

    //create
function fun (){
    console.log('this is a function');
}
    //call
    fun();
    // LETS CREATE a function that take in a name and 
    //says hello follow by name
    //name: "reza"
    //return = hello reza
    
    function greeting(){

       var name= prompt("what is your name");
        var result= 'Hello '+ name; //string concatenation
        console.log(result);
    }
    //greeting();

    // how do arguments work in functions?
    // how do we add 2 numbers together in a function
    
    function sumNumbers(num1, num2){
        var result= num1 + num2
        console.log(result);
    }
    sumNumbers(20, 20);

    // while loops
    var num =0;
    
    while (num <100){
        num+= 1;
        console.log(num);
    }