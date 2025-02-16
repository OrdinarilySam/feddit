<form method="POST" action="?/login">
	<label for="username">Username</label>
	<input type="text" id="username" name="username" required />

	<label for="password">Password</label>
	<input type="password" id="password" name="password" required />

	<button type="submit">Login</button>
</form>

<style>
	form {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	label {
		font-size: 1.5rem;
	}

	input {
		font-size: 1.5rem;
		padding: 0.5rem;
		margin: 0.5rem;
	}

	button {
		font-size: 1.5rem;
		padding: 0.5rem 1rem;
		margin: 0.5rem;
		background-color: var(--color-rose-quartz);
		border: none;
		border-radius: 1rem;
	}
</style>
