<script lang="ts">
    import {uploadFile} from "$lib/api/api.ts";
	import { onMount } from 'svelte';
	import {FileDropzone, getToastStore, type ToastSettings} from '@skeletonlabs/skeleton';
    const toastStore = getToastStore();
	export let form;

    function toastRegisterError(error: string): void {
        const t: ToastSettings = {
            message: error,
            background: 'variant-filled-error',
        };
        toastStore.trigger(t);
    }

    if (form?.error) {
        toastRegisterError(form?.error);
    }

	let name = '';
	let surname = '';
	let phone = '';
	let email = '';
	let avatarUrl = '';

	onMount(() => {
		if (typeof localStorage !== 'undefined') {
			name = localStorage.getItem('form_name') || '';
			surname = localStorage.getItem('form_surname') || '';
			phone = localStorage.getItem('form_phone') || '';
			email = localStorage.getItem('form_email') || '';
			avatarUrl = localStorage.getItem('form_avatarUrl') || '';
		}
	});

	function saveToLocalStorage(key: string, value: string) {
		localStorage.setItem(key, value);
	}

	async function fileHandler(e: Event) {
		const {data: url, error: err} = await uploadFile(e);
		if (err) {
			toastRegisterError('failed to upload file');
			return
		}
		avatarUrl = url;
		saveToLocalStorage('form_avatarUrl', avatarUrl);
		console.log(avatarUrl);
	}

	function resetFormData() {
		if (typeof localStorage !== 'undefined') {
			localStorage.removeItem('form_name');
			localStorage.removeItem('form_surname');
			localStorage.removeItem('form_phone');
			localStorage.removeItem('form_email');
			localStorage.removeItem('form_avatarUrl');
		}
	}

	function handleSubmit() {
		resetFormData();
	}
</script>

<div class="flex justify-center h-full pt-5">
	<div class="card p-4 w-modal shadow-xl space-y-4">
		<form class="modal-form border border-surface-500 p-4 space-y-4 rounded-container-token"
		      method="POST" action="?/register" on:submit={handleSubmit}>
			<label class="label">
				<span>Name</span>
				<input class="input" type="text" name="name" bind:value={name}
					   placeholder="Enter name..."
					   on:input={() => saveToLocalStorage('form_name', name)}/>
			</label>
			<label class="label">
				<span>Surname</span>
				<input class="input" type="text" name="surname" bind:value={surname}
					   placeholder="Enter surname..."
					   on:input={() => saveToLocalStorage('form_surname', surname)}/>
			</label>
			<label class="label">
				<span>Phone number</span>
				<input class="input" type="tel" name="phone" bind:value={phone}
					   placeholder="Enter phone number..."
					   on:input={() => saveToLocalStorage('form_phone', phone)}/>
			</label>
			<label class="label">
				<span>Email</span>
				<input class="input" type="email" name="email" bind:value={email}
					   placeholder="Enter email..."
					   on:input={() => saveToLocalStorage('form_email', email)}/>
			</label>
			<input type="hidden" name="avatar_url" bind:value={avatarUrl}
				   on:input={() => saveToLocalStorage('form_avatarUrl', avatarUrl)}/>
			<label class="label">
				<span>Password</span>
				<input class="input" type="password" name="password" value='' placeholder="Enter password..." />
			</label>
			<div class="label">
				{#if avatarUrl !== ''}
					<img src={avatarUrl} alt="avatar"/>
				{/if}
				<FileDropzone name="file" on:change={fileHandler}>
					<svelte:fragment slot="message">Upload or drag and drop user avatar</svelte:fragment>
				</FileDropzone>
			</div>
			<button class="input btn btn-md variant-ghost-primary" type="submit">Register</button>
		</form>
	</div>
</div>

<style>
    :global(body) {
        display: flex;
        flex-direction: column;
        min-height: 100vh;
    }
</style>
