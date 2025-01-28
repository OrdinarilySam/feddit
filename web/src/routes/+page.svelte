<script lang="ts">
	import Post from "./Post.svelte";

  export let data;
</script>

<div>
  <h1>Posts</h1>
  <a href="/new">+</a>
</div>

{#each data.data as post}
  <Post data={post} />
{/each}

<style>
  a {
    text-decoration: none;
    color: inherit;
    font-size: 2rem;
    background-color: var(--color-emerald);
    height: 2.5rem;
    width: 2.5rem;
    text-align: center;
    border-radius: 1rem;
  }

  div {
    display: flex;
    flex-direction: row;
    align-items: center;
  }
</style>