import { describe, expect, it, vi } from 'vitest';
import { sayHello } from '$lib/api';

describe('sayHello', () => {
  it('returns parsed JSON', async () => {
    vi.stubGlobal('fetch', vi.fn(async () =>
      new Response(JSON.stringify({ message: 'Hello Ada!', server_time: '2026-07-22T00:00:00Z' }), { status: 200 })
    ));
    expect((await sayHello('Ada')).message).toBe('Hello Ada!');
  });
  it('throws on non-2xx', async () => {
    vi.stubGlobal('fetch', vi.fn(async () => new Response('nope', { status: 500 })));
    await expect(sayHello('X')).rejects.toThrow(/hello failed/);
  });
});
