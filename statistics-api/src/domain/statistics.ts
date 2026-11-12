export function statistics(matrices: number[][][], epsilon: number) {
  let max = -Infinity;
  let min = Infinity;
  let sum = 0;
  let correction = 0;
  let count = 0;
  const diagonalMatrices: number[] = [];
  for (const [index, matrix] of matrices.entries()) {
    let diagonal = matrix.length === matrix[0].length;
    for (const [i, row] of matrix.entries()) {
      for (const [j, value] of row.entries()) {
        max = Math.max(max, value);
        min = Math.min(min, value);
        const next = sum + value;
        correction += Math.abs(sum) >= Math.abs(value) ? (sum - next) + value : (value - next) + sum;
        sum = next;
        count++;
        if (i !== j && Math.abs(value) > epsilon) diagonal = false;
      }
    }
    if (diagonal) diagonalMatrices.push(index);
  }
  sum += correction;
  return { max, min, average: sum / count, sum, anyDiagonal: diagonalMatrices.length > 0,
    diagonalMatrices, matrixCount: matrices.length };
}
