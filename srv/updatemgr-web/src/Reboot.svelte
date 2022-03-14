<script>
	export let reboot_data = {};
	export let baseUrl = "";

	let hostname = reboot_data['name'];
	let required = reboot_data['reboot_required'];

	async function callUpdate() {
		let res = fetch(`${baseUrl}/reboot/${hostname}`,{
			method: "POST"
		});
		return res;
	}
</script>

<div class="reboot">
<button on:click={callUpdate}>
  Reboot {hostname}
</button>
</div>
