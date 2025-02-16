export async function login(email: string, password: string) {
	const res = await fetch('http://localhost:3000/users/sign_in', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user: { email, password } })
	});

	if (!res.ok) throw new Error('Login failed');
	const data = await res.json();

	localStorage.setItem('jwt', data.token);
	return data.user;
}
