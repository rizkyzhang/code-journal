/**
 * Pseudocode
 * 
 * ARGUMENT string, charToCount
 * 
 * SET count to 0
 * 
 * FOR each char in string
 *    IF char is equal to charToCount
 *        INCREMENT count by 1
 *    ENDIF
 * ENDFOR
 * 
 * RETURN count
 */

function countChar(string, charToCount) {
  let count = 0;

  for (const char of string) {
    if (char === charToCount) {
      count++;
    }
  }

  return count;
}

const countBs = string => countChar(string, "B");



console.log(countBs("BBCB"));
console.log(countChar("kakkerlak", "k"));