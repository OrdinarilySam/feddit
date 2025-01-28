

<form method="POST">
  <label for="title">Title</label>
  <input name="title" id="new-form-title" />
  <label for="content">Content</label>
  <textarea name="content" id="new-form-content"></textarea>
  <button type="submit">Submit</button>
</form>

<style>
  form {
    display: flex;
    flex-direction: column;
  }

  form > * {
    margin-bottom: 0.5rem;
  }

  input {
    max-width: 200px;
  }

  textarea {
    max-width: 300px;
  }
</style>