import { Router } from 'express';
import { analyze } from '../application/analyze';
import { validMatrices } from './validation';
import { fail } from './error-handler';

export function routes(epsilon: number) {
  const router = Router();
  router.get(['/health/live', '/health/ready'], (_req, res) => res.json({status: 'ok'}));
  router.post('/internal/v1/statistics', (req, res) => {
    const matrices: unknown = req.body?.matrices;
    if (!validMatrices(matrices)) return fail(res, 422, 'INVALID_MATRIX', 'Expected 1–10 rectangular matrices of finite numbers up to 100×100');
    const result = analyze(matrices, epsilon);
    if (!Number.isFinite(result.sum)) return fail(res, 422, 'NUMERIC_OVERFLOW', 'Sum exceeds numeric range');
    return res.json(result);
  });
  return router;
}
