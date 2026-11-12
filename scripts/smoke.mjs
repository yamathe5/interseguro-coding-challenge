import assert from 'node:assert/strict';

const base = process.env.MATRIX_URL ?? 'http://127.0.0.1:8080';
const down = process.argv.includes('--dependency-down');

async function call(path, status, body) {
  const response = await fetch(new URL(path, base), {
    method: body === undefined ? 'GET' : 'POST',
    headers: {'Content-Type':'application/json', 'X-Request-ID':'integration-smoke'},
    body: body === undefined ? undefined : typeof body === 'string' ? body : JSON.stringify(body),
    signal: AbortSignal.timeout(10_000),
  });
  assert.equal(response.status, status, path);
  assert.equal(response.headers.get('x-request-id'), 'integration-smoke');
  const result = await response.json();
  if (status >= 400) assert.equal(result.error.requestId, 'integration-smoke');
  return result;
}

const matrices = '/api/v1/matrices';
const matrix = [[1,2,3],[4,5,6]];
assert.equal((await call('/health/live', 200)).status, 'ok');

if (down) {
  assert.equal((await call('/health/ready', 503)).error.code, 'NOT_READY');
  assert.equal((await call(matrices+'/analyze', 502, {matrix, operation:'qr'})).error.code, 'DEPENDENCY_FAILURE');
  await call(matrices+'/qr', 200, {matrix});
} else {
  assert.equal((await call('/health/ready', 200)).status, 'ok');
  const rotated = [[4,1],[5,2],[6,3]];
  assert.deepEqual(await call(matrices+'/rotate', 200, {matrix}), {matrix:rotated, rotationDegrees:90});
  const rotation = await call(matrices+'/analyze', 200, {matrix, operation:'rotate'});
  assert.deepEqual(rotation.result.matrix, rotated);
  assert.deepEqual(rotation.statistics, {max:6,min:1,average:3.5,sum:21,anyDiagonal:false,diagonalMatrices:[],matrixCount:1});
  const qr = await call(matrices+'/qr', 200, {matrix});
  const analysis = await call(matrices+'/analyze', 200, {matrix, operation:'qr'});
  assert.deepEqual(analysis.result, qr);
  const elements = [...qr.q.flat(), ...qr.r.flat()];
  const expected = elements.reduce((a,b) => a+b, 0);
  assert.equal(analysis.statistics.matrixCount, 2);
  assert.equal(analysis.statistics.max, Math.max(...elements));
  assert.equal(analysis.statistics.min, Math.min(...elements));
  assert.ok(Math.abs(analysis.statistics.sum-expected) < 1e-9);
  assert.ok(Math.abs(analysis.statistics.average-expected/elements.length) < 1e-9);
  assert.equal((await call(matrices+'/qr', 400, '{')).error.code, 'INVALID_JSON');
  assert.equal((await call(matrices+'/qr', 422, {matrix:[[1],[2,3]]})).error.code, 'INVALID_MATRIX');
  assert.equal((await call(matrices+'/rotate', 422, {matrix,direction:'left'})).error.code, 'INVALID_DIRECTION');
}
console.log('HTTP smoke passed: ' + (down ? 'dependency unavailable' : 'both services available'));
