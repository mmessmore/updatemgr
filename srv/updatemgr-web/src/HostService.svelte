<script>
	import Reboot from './Reboot.svelte';
	import Update from './Update.svelte';


	async function getHosts() {
		let stringy = "";
		let url = `./hosts`;
		const hostsRes = await fetch(url,{
			cache: "no-cache",
			mode: "cors",
			referrerPolicy: "unsafe-url"
		});

		let data = await hostsRes.json();
		for (const key in data['hosts']) {
			data.hosts[key]['date'] = new Date(data.hosts[key]['online']['timestamp'] * 1000);
		}
		return data;
	}

	let hosts = getHosts();

	// refresh every 30 seconds
	setInterval(() => {
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
	<p>Last check-in {host['date'].toLocaleString()}</p>
{#if host['reboot_required']['reboot_required'] }
<Reboot reboot_data={host['reboot_required']} />
{/if}
<h3>Updates Available</h3>
{#if host['updates_available']['packages'].length > 0 }
	<Update hostname={host['name']}/>
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
