/**
 * Pseudocode
 *
 * ARGUMENT size
 *
 * WHILE row index (i) (starting from 0) < size
 *      SET row to ""
 *      WHILE column index (j) (strating from 0) < size
 *          IF both index is an even number or both index is an odd number
 *              APPEND space to row
 *          ELSE
 *              APPEND hash to row
 *          ENDIF
 *
 *          INCREMENT j by 1
 *      ENDWHILE
 *
 *      OUTPUT row
 *
 *      INCREMENT i by 1
 * ENDWHILE
 */

function chessboard(size) {
  for (let i = 0; i < size; i++) {
    let row = "";

    for (let j = 0; j < size; j++) {
      const isBothIndexEven = i % 2 === 0 && j % 2 === 0;
      const isBothIndexOdd = i % 2 !== 0 && j % 2 !== 0;

      if (isBothIndexEven || isBothIndexOdd) {
        row += " ";
      } else {
        row += "#";
      }
    }

    console.log(row);
  }
}

chessboard(10);
