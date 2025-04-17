<script lang="ts">
    import type {User} from "$lib/types.ts";

    export let data;
    const user = data.user as User
</script>

<div class="flex justify-center h-full pt-5">
	<div class="card p-4 w-modal shadow-xl space-y-4">
		<h1 class="text-2xl text-center">Profile</h1>
		<div class="modal-form border border-surface-500 p-4 space-y-4 rounded-container-token">
			{#if user.avatar_url !== ''}
				<img src={user.avatar_url} alt="avatar">
			{/if}
			<div class="grid grid-cols-3">
				<span class="font-bold p-1.5">First name:</span>
				<input type="text" class="input col-span-2" value={user.name} disabled>
			</div>
			<div class="grid grid-cols-3">
				<span class="font-bold p-1.5">Last name:</span>
				<input type="text" class="input col-span-2" value={user.surname} disabled/>
			</div>
			<div class="grid grid-cols-3">
				<span class="font-bold p-1.5">Phone number:</span>
				<input type="tel" class="input col-span-2" value={user.phone} disabled/>
			</div>
			<div class="grid grid-cols-3">
				<span class="font-bold p-1.5">City:</span>
				<input type="tel" class="input col-span-2" value={user.city} disabled/>
			</div>
			<div class="grid grid-cols-3">
				<span class="font-bold p-1.5">Email:</span>
				<input type="email" class="input col-span-2" value={user.email} disabled/>
			</div>
		</div>
	</div>
</div>
