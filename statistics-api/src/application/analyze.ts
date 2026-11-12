import { statistics } from '../domain/statistics';
export function analyze(matrices: number[][][], epsilon: number) {
  return statistics(matrices, epsilon);
}
