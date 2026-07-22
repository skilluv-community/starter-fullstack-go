const BASE = (import.meta.env.PUBLIC_API_BASE as string | undefined) ?? 'http://localhost:3001';

export type Note = { id: string; text: string; created_at: string };

export async function sayHello(name: string) {
  const url = new URL('/api/hello', BASE);
  if (name) url.searchParams.set('name', name);
  const r = await fetch(url);
  if (!r.ok) throw new Error(`hello failed: ${r.status}`);
  return r.json() as Promise<{ message: string; server_time: string }>;
}

export async function listNotes(): Promise<Note[]> {
  const r = await fetch(new URL('/api/notes', BASE));
  if (!r.ok) throw new Error(`listNotes failed: ${r.status}`);
  return r.json();
}

export async function createNote(text: string): Promise<Note> {
  const r = await fetch(new URL('/api/notes', BASE), {
    method: 'POST',
    headers: { 'content-type': 'application/json' },
    body: JSON.stringify({ text })
  });
  if (!r.ok) throw new Error(`createNote failed: ${r.status}`);
  return r.json();
}

export async function deleteNote(id: string): Promise<void> {
  const r = await fetch(new URL(`/api/notes/${id}`, BASE), { method: 'DELETE' });
  if (!r.ok) throw new Error(`deleteNote failed: ${r.status}`);
}
