import express from 'express';
import { randomUUID } from 'node:crypto';
import { routes } from './http/routes';
import { errorHandler, fail } from './http/error-handler';

export function createApp(epsilon = 1e-9) {
  const app = express();
  app.disable('x-powered-by');
  app.use((req, res, next) => {
    const incoming = req.get('X-Request-ID');
    const requestId = incoming && /^[A-Za-z0-9_-]{1,64}$/.test(incoming) ? incoming : randomUUID();
    const start = Date.now();
    res.set('X-Request-ID', requestId);
    res.on('finish', () => console.info(JSON.stringify({requestId, method: req.method, status: res.statusCode, durationMs: Date.now()-start})));
    next();
  });
  app.use(express.json({limit: '512kb', strict: false}));
  app.use(routes(epsilon));
  app.use((_req, res) => { fail(res, 404, 'NOT_FOUND', 'Route not found'); });
  app.use(errorHandler);
  return app;
}
