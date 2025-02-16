<nav>
	<h1><a href="/">Website Title</a></h1>
	<ul>
		<li><a href="/login">Login</a></li>
	</ul>
</nav>

<div>
	<slot />
</div>

<style>
	:global(*) {
		box-sizing: border-box;

		--color-aquamarine: #b5ffe1;
		--color-celadon: #93e5ab;
		--color-mint: #65b891;
		--color-dark-cyan: #4e878c;
		--color-dark-green: #00241b;

		--color-gunmetal: #272d2d;
		--color-rose-quartz: #a39ba8;
		--color-powder-blue: #b8c5d6;
		--color-alice-blue: #edf5fc;
		--color-emerald: #23ce6b;
	}

	:global(body) {
		margin: 0;
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
		background-color: var(--color-alice-blue);
		color: var(--color-gunmetal);
	}

	nav {
		padding: 0.25rem 0.5rem;
		width: 100%;
		background-color: var(--color-powder-blue);
		margin: 0;
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
	}

	ul {
		display: flex;
		flex-direction: row;
		margin: 0;
		list-style-type: none;
		padding-left: 0;
	}

	li {
		margin-left: 0;
	}

	h1 {
		margin: 0;
		font-size: 2.5rem;
	}

	a {
		text-decoration: none;
		color: var(--color-gunmetal);
	}

	div {
		padding: 0 0.5rem;
	}
</style>
