<script lang="ts">
  import '../app.css';
  import { lang } from '$lib/i18n/index.svelte';
  let { children } = $props();
</script>

<header class="border-b border-slate-800 px-6 py-4 flex items-center justify-between">
  <a href="/" class="font-bold text-lg">Skilluv · Go</a>
  <nav class="flex gap-4 items-center">
    <a href="/" class="hover:underline">home</a>
    <a href="/notes" class="hover:underline">notes</a>
    <button class="rounded border border-slate-700 px-2 py-1 text-xs hover:bg-slate-800" onclick={() => lang.toggle()}>
      {lang.t.lang_switch}
    </button>
  </nav>
</header>
<main class="p-6 max-w-3xl mx-auto">{@render children()}</main>
<footer class="mt-16 px-6 py-4 text-xs text-slate-500 border-t border-slate-800">
  MIT · <a href="https://skilluv.io" class="underline">skilluv.io</a>
</footer>
