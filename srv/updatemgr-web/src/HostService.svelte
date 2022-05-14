<script>
	import Reboot from './Reboot.svelte';
	import Update from './Update.svelte';

	let baseUrl = "";

	async function getHosts() {
		let stringy = "";
		let url = `${baseUrl}/hosts`;
		const hostsRes = await fetch(url,{
			cache: "no-cache",
			mode: "cors",
			referrerPolicy: "unsafe-url"
		});

		let data = await hostsRes.json();
		return data;
	}

	let hosts = getHosts();

	// refresh every 30 seconds
	setInterval(() => {
		console.log("refreshing...");
		hosts = getHosts();
	}, 30000);
</script>

<div>
{#await hosts }
waiting...
{:then data}
<div class="box-container">
{#each data['hosts'] as host}
<div class="box">
<h2>{host['name']}</h2>
{#if host['reboot_required']['reboot_required'] }
<Reboot reboot_data={host['reboot_required']} baseUrl={baseUrl} />
{/if}
<h3>Updates Available</h3>
{#if host['updates_available']['packages'].length > 0 }
	<Update baseUrl={baseUrl} hostname={host['name']}/>
	<ul class="pkgs">
		{#each host['updates_available']['packages'] as pkg}
		<li class="pkg">{pkg}</li>
		{/each}
	</ul>
{:else}
	<p>No updates available.</p>
{/if}
</div>
{/each}
</div>
{:catch error}
<span style="color:red;">{error.message}</span>
{/await}
</div>

<style>
h2 {
margin: .2em;
padding: 0px;
}
h3 {
margin: .2em;
padding: 0px;
}
.box-container {
display: grid;
row-gap: 10px;
column-gap: 10px;
grid-template-columns: repeat(6, 400px);
}
.box {
border: 1px solid #000;
max-height: 600px;
overflow: auto;
}
.pkgs {
display: grid;
grid-template-columns: repeat(3, 6em);
list-style-type: none;
}
.pkg {
border: 1px solid #0000;
font-size: 0.75em;
font-family: monospace;
}

</style>
