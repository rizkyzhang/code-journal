/**
 * Pseuodocode
 * 
 * ARGUMENT n
 * IF n is equal to 1
 *    RETURN false
 * ENDIF
 * 
 * IF n is equal to 0
 *    RETURN true
 * ENDIF
 * 
 * IF n > 0
 *    RETURN func(n - 2)
 * ELSE
 *    RETURN func(n + 2)
 * ENDIF
 */

function isEven(n) {
  if (n === 0) {
    return true;
  }

  if (n === 1) {
    return false;
  }

  return n > 0 ? isEven(n - 2) : isEven(n + 2);
}

console.log(isEven(-9));