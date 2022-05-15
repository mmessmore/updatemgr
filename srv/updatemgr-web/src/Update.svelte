<script>
	export let hostname = {};

	async function callUpdate() {
		let res = await fetch(`./upgrade/${hostname}`,{
			method: "POST"
		});
		let text = await res.text();
		if (res.ok) {
			return "Updated"
		} else {
			throw new Error(text);
		}
	}

	let promise;

	function handleClick() {
		promise = callUpdate();
	}
</script>

<div class="update">
{#if typeof promise == 'undefined'}
<button on:click={handleClick}>
  Update {hostname}
</button>
{/if}
{#if typeof promise !== 'undefined'}
{#await promise}
waiting...
{:then text}
<p>{text}</p>
{:catch error}
<p style="color:red;">{error.message}</p>
{/await}
{/if}
</div>
