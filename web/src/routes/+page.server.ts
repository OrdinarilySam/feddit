// import type { Post } from "$lib/types";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
  const res = await fetch("http://localhost:8080/posts")
  const data = await res.json()
  return {
    data
  };
  // return []
}