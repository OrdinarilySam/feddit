<script lang="ts">
	import Post from './Post.svelte';
	import type { Post as PostT } from '$lib/types';

	export let data: PostT[];
</script>

<div>
	<h1>Posts</h1>
	<a href="/new">New Post</a>
</div>

{#each data as post}
	<Post data={post} />
{/each}

<style>
	a {
		text-decoration: none;
		color: inherit;
		font-size: 1.5rem;
		background-color: var(--color-rose-quartz);
		text-align: center;
		border-radius: 1rem;
		padding: 0.5rem 1rem;
	}

	div {
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-between;
		margin: 0.5rem 0;
	}

	h1 {
		margin: 0;
	}
</style>
