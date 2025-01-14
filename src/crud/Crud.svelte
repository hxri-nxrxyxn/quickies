<script>
  import { get, post, put, del } from "./fetch"
  let name = $state("")

  const doGet = () => {
  	console.log(name)
  }
</script>

<main>
	<input type="text" bind:value={name}>	
	<button onclick={doGet}>hey</button>
</main>

<style>
	
</style>