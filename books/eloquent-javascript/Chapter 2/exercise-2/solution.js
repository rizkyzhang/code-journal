/**
 * Pseudocode
 *
 * FOR each number (i) starting from 1 to 100
 *      IF i is divisible only by 3
 *          OUTPUT "Fizz"
 *      ELSE IF i is divible only by 5
 *          OUTPUT "Buzz"
 *      ELSE IF i is divisible both by 3 and 5
 *          OUTPUT "FizzBuzz"
 *      ELSE
 *          OUTPUT i
 *      ENDIF
 * ENDFOR
 */

function fizzbuzz() {
  for (let i = 1; i <= 100; i++) {
    if (i % 15 === 0) {
      console.log("FizzBuzz");
    } else if (i % 3 === 0) {
      console.log("Fizz");
    } else if (i % 5 === 0) {
      console.log("Buzz");
    } else {
      console.log(i);
    }
  }
}

function fizzbuzz2() {
  for (let i = 1; i <= 100; i++) {
    let output = "";

    if (i % 3 === 0) {
      output += "Fizz";
    }

    if (i % 5 === 0) {
      output += "Buzz";
    }

    output.length > 0 ? console.log(output) : console.log(i);
    // output !== "" ? console.log(output) : console.log(i);
    // console.log(output || i);
  }
}

fizzbuzz2();
