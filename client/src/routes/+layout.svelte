<script lang="ts">
	import '../app.css';
	import LeftSideBar from "$lib/components/LeftSideBar.svelte";
	let { children } = $props();
</script>

<main class="flex gap-2 min-h-screen">
	<LeftSideBar/>
	{@render children()}
</main>
