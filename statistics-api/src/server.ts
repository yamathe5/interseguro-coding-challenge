import { config } from './config';
import { createApp } from './app';
const { port, epsilon } = config();
const server = createApp(epsilon).listen(port, '0.0.0.0');
server.requestTimeout = 10_000;
server.headersTimeout = 5_000;
for (const signal of ['SIGTERM', 'SIGINT'] as const) {
  process.once(signal, () => server.close());
}
