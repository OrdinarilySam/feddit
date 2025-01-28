import type { Actions } from "./$types";
import { redirect } from "@sveltejs/kit";

export const actions = {
  default: async ({ request }) => {
    const data = await request.formData();
    const title = data.get("title")
    const content = data.get("content")
    const res = await fetch("http://localhost:8080/posts", {
      method: "POST", 
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        title,
        content
      })
    });

    if (!res.ok) {
      return { success: false, error: await res.json() };
    }

    throw redirect(303, '/');
  }
} satisfies Actions;