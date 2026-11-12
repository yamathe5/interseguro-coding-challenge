import { spawn, spawnSync } from 'node:child_process';
import { once } from 'node:events';
import { resolve } from 'node:path';
import { setTimeout } from 'node:timers/promises';

const binary = process.argv[2];
if (!binary) throw new Error('Pass the path to the compiled matrix-api binary');
const matrixPort = process.env.SMOKE_PORT ?? '19808';
const statisticsPort = process.env.SMOKE_STATISTICS_PORT ?? '19300';
const url = 'http://127.0.0.1:' + matrixPort;
const services = [];
function start(command, args, options) {
  const child = spawn(command, args, { stdio: 'inherit', ...options });
  child.on('error', err => console.error(err.message));
  services.push(child);
  return child;
}
function smoke(...args) {
  const result = spawnSync(process.execPath, ['scripts/smoke.mjs', ...args], {
    stdio: 'inherit', env: {...process.env, MATRIX_URL: url},
  });
  if (result.status !== 0) throw new Error('Smoke test failed');
}

try {
  const statistics = start(process.execPath, ['dist/server.js'], {
    cwd: 'statistics-api', env: {...process.env, PORT: statisticsPort},
  });
  const go = start(resolve(binary), [], {
    env: {...process.env, PORT: matrixPort, STATISTICS_URL: 'http://127.0.0.1:' + statisticsPort},
  });
  let ready = false;
  for (let attempt = 0; attempt < 30 && !ready; attempt++) {
    if (statistics.exitCode !== null || go.exitCode !== null || !statistics.pid || !go.pid) {
      throw new Error('A service failed to start');
    }
    try {
      ready = (await fetch(url + '/health/ready', {signal: AbortSignal.timeout(1000)})).ok;
    } catch {
      // Services may still be binding their ports.
    }
    if (!ready) await setTimeout(250);
  }
  if (!ready) throw new Error('Services did not become ready');
  smoke();
  const closed = once(statistics, 'exit');
  statistics.kill();
  await closed;
  smoke('--dependency-down');
} finally {
  for (const child of services) {
    if (child.exitCode === null) child.kill();
  }
}
