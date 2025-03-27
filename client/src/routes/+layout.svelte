<script lang="ts">
	import '../app.css';
	import LeftSideBar from "$lib/components/LeftSideBar.svelte";
	let { children } = $props();
</script>

<main class="flex min-h-screen pl-[100px]">
	<LeftSideBar/>
	{@render children()}
</main>
