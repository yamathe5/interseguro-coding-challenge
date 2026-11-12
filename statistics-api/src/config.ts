export function config(env = process.env) {
  const port = Number(env.PORT ?? 3000);
  const epsilon = Number(env.DIAGONAL_EPSILON ?? 1e-9);
  if (!Number.isInteger(port) || port < 1 || port > 65535) throw new Error('Invalid PORT');
  if (!Number.isFinite(epsilon) || epsilon < 0) throw new Error('Invalid DIAGONAL_EPSILON');
  return { port, epsilon };
}
