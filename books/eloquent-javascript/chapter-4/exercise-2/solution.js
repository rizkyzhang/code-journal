/**
 * Pseudocode
 *
 * ARGUMENT arr
 * SET reversedArr to []
 *
 * FOR each array index (i) from last index to first index
 *    PUSH element index i into reversedArr
 * ENDFOR
 *
 * RETURN reversedArr
 */

const reverseArray = (arr) => {
  const reversedArr = [];

  for (let i = arr.length - 1; i >= 0; i--) {
    reversedArr.push(arr[i]);
  }

  return reversedArr;
};
