export function validMatrices(value: unknown): value is number[][][] {
  return Array.isArray(value) && value.length > 0 && value.length <= 10 &&
    value.every(matrix => Array.isArray(matrix) && matrix.length > 0 && matrix.length <= 100 &&
      Array.isArray(matrix[0]) && matrix[0].length > 0 && matrix[0].length <= 100 &&
      matrix.every(row => Array.isArray(row) && row.length === matrix[0].length &&
        row.every(x => typeof x === 'number' && Number.isFinite(x))));
}
