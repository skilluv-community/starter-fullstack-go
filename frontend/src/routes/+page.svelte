<script lang="ts">
  import { sayHello } from '$lib/api';
  import { lang } from '$lib/i18n/index.svelte';

  let name = $state('Skilluv');
  let response = $state<string | null>(null);
  let loading = $state(false);
  let error = $state<string | null>(null);

  async function greet() {
    loading = true;
    error = null;
    try {
      response = (await sayHello(name)).message;
    } catch (e) {
      error = (e as Error).message;
    } finally {
      loading = false;
    }
  }
</script>

<h1 class="text-3xl font-bold mb-2">{lang.t.title}</h1>
<p class="text-slate-400 mb-8">{lang.t.tagline}</p>

<div class="rounded border border-slate-800 p-4 flex gap-2 items-center">
  <input
    class="bg-slate-900 border border-slate-700 rounded px-3 py-2 flex-1"
    placeholder={lang.t.greeting_placeholder}
    bind:value={name}
  />
  <button
    class="bg-emerald-600 hover:bg-emerald-500 disabled:opacity-50 rounded px-4 py-2 font-semibold"
    onclick={greet}
    disabled={loading}
  >
    {lang.t.greeting_button}
  </button>
</div>

{#if response}<p class="mt-4 text-emerald-400">{response}</p>{/if}
{#if error}<p class="mt-4 text-red-400">{error}</p>{/if}
