const {test} = require('node:test');
const assert = require('node:assert/strict');
const {once} = require('node:events');
const {statistics} = require('../dist/domain/statistics');
const {createApp} = require('../dist/app');

test('aggregates globally and applies diagonal tolerance', () => {
  assert.deepEqual(statistics([[[1,0],[0,2]],[[3,4],[5,6]]], 1e-9),
    {max:6,min:0,sum:21,average:2.625,anyDiagonal:true,diagonalMatrices:[0],matrixCount:2});
  assert.deepEqual(statistics([[[0,1e-10],[-1e-10,-2.5]], [[0,0]], [[1,2],[0,1]]],1e-9).diagonalMatrices,[0]);
  assert.equal(statistics([[[-3.5,-2]]],1e-9).max,-2);
  assert.equal(statistics([[[1e16,1,-1e16]]],1e-9).sum,1);
});

test('HTTP statistics and errors', async t => {
  const server = createApp().listen(0, '127.0.0.1');
  t.after(() => new Promise(resolve => server.close(resolve)));
  await once(server,'listening');
  const endpoint = 'http://127.0.0.1:' + server.address().port + '/internal/v1/statistics';
  for (const [body,status] of [['{"matrices":[[[1]]]}',200],['{',400],['{}',422],
    ['{"matrices":[[[1],[2,3]]]}',422],['{"matrices":[[[null]]]}',422],
    ['{"matrices":[[[1e309]]]}',422],['{"matrices":[[[1e308,1e308]]]}',422]]) {
    const response = await fetch(endpoint,{method:'POST',headers:{'Content-Type':'application/json','X-Request-ID':'test'},body});
    assert.equal(response.status,status,body);
    assert.equal(response.headers.get('x-request-id'),'test');
    const json = await response.json();
    if (status === 200) assert.equal(json.sum,1);
    else assert.equal(json.error.requestId,'test');
  }
});
