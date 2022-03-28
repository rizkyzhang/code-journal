/**
 * Pseudocode
 * ARGUMENT min, max
 * SET nums to []
 *
 * FOR each number (i) from min to max
 *    PUSH i into nums
 * ENDFOR
 *
 * RETURN nums
 */

const range = (min, max) => {
  const nums = [];

  for (let i = min; i <= max; i++) {
    nums.push(i);
  }

  return nums;
};

const sum = (nums) => {
  let sum = 0;

  for (const num of nums) {
    sum += num;
  }

  return sum;
};

const sum2 = (nums) => nums.reduce((accu, num) => accu + num, 0);

console.log(sum2(range(1, 10)));
