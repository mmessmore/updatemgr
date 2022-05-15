<script>
	export let reboot_data = {};

	let hostname = reboot_data['name'];
	let required = reboot_data['reboot_required'];

	async function callUpdate() {
		let res = fetch(`./reboot/${hostname}`,{
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
