<script>
/*
	users
	+----+---------+
	| id | name    |
	+----+---------+
	|  1 | hari    |
	|  2 | narayan |
	|  3 | al      |
	|  4 | ammu    |
	|  5 | st      |
	+----+---------+
*/
	const baseUrl = "localhost:8080"
	import { get, post, put, del } from "./fetch"
	let id = $state(0)
	let userName = $state("")
	let updatedName = $state("")

	const doGet = async () => {
		const data = await get('/users'); 
		console.log(data);
	}

	const doGetId = async () => {
		const data = await get(`/users/${id}`); 
		console.log(data);
	}

	const doPost = async () => {
		const newData = {
			name: userName,
		};
		const createdData = await post('/users', newData);
		console.log(createdData);	
	}

	const doPut = async () => {
		const updatedData = {
			name: updatedName,
		};
		const updatedResult = await put(`/users/${id}`, updatedData);
		console.log(updatedResult);
	}

	const doDel = async () => {
		const deleted = await del(`/users/${id}`);
		console.log('Deleted:', deleted); 
	}
</script>

<main>
	<div>
		<button onclick={doGet}>GET</button>
	</div>
	<div>
		<input placeholder="id" type="text" bind:value={id}>	
		<button onclick={doGetId}>GETu</button>
	</div>
	<div>
		<input placeholder="username" type="text" bind:value={userName}>	
		<button onclick={doPost}>POST</button>
	</div>
	<div>
		<input placeholder="id" type="text" bind:value={id}>	
		<input placeholder="new name" type="text" bind:value={updatedName}>	
		<button onclick={doPut}>PUT</button>
	</div>
	<div>
		<input placeholder="id" type="text" bind:value={id}>	
		<button onclick={doDel}>DELETE</button>
	</div>
</main>

<style>
	div {
		margin: 5rem 0;
	}	
</style>
