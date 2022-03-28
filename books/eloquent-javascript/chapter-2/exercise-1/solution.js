/**
 * Pseudocode
 *
 * ARGUMENT n
 * SET row to "#"
 *
 * WHILE length of row <= n
 *      OUTPUT row
 *      APPEND "#" to row
 * ENDWHILE
 */

function triangle(n) {
  for (let row = "#"; row.length <= n; row += "#") {
    console.log(row);
  }
}

function triangle2(n) {
  let row = "#";

  while (row.length <= n) {
    console.log(row);
    row += "#";
  }
}

function triangle3(n) {
  for (let i = 1; i <= n; i++) {
    console.log("#".repeat(i));
  }
}

function triangle4(n, row = "") {
  row += "#";
  console.log(row);

  if (row.length < n) {
    triangle4(n, row);
  }
}

triangle4(7);
