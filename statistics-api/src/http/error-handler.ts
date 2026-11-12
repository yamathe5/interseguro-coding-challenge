import type { ErrorRequestHandler, Response } from 'express';
export function fail(res: Response, status: number, code: string, message: string) {
  return res.status(status).json({error: { code, message, requestId: res.getHeader('X-Request-ID') }});
}
export const errorHandler: ErrorRequestHandler = (err, _req, res, next) => {
  if (res.headersSent) { next(err); return; }
  if (err.status === 400) { fail(res, 400, 'INVALID_JSON', 'Invalid JSON'); return; }
  if (err.status === 413) { fail(res, 413, 'BODY_TOO_LARGE', 'Body exceeds 512 KiB'); return; }
  if (err.status === 415) { fail(res, 415, 'UNSUPPORTED_MEDIA_TYPE', 'Unsupported encoding or charset'); return; }
  fail(res, 500, 'INTERNAL_ERROR', 'Unexpected error');
};
