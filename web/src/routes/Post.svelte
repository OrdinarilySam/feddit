<script lang="ts">
  import type { Post } from "$lib/types";

  export let data: Post;
</script>


<div>
  <h2>{data.title}</h2>
  <p>{data.content}</p>
</div>

<style>
  div {
    background-color: var(--color-powder-blue);
    padding: 0.5rem;
    margin-bottom: 0.5rem;
  }

  h2 {
    margin: 0;
  }

  p {
    margin-bottom: 0;
  }

</style>
