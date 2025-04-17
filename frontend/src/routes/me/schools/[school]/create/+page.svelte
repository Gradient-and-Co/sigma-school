<script lang="ts">
    import {uploadFile} from "$lib/api/api.ts";

	export let form;

	import {FileDropzone, getToastStore, type ToastSettings} from '@skeletonlabs/skeleton';
    const toastStore = getToastStore();

    function toastCreateError(error: string): void {
        const t: ToastSettings = {
            message: error,
            background: 'variant-filled-error',
        };
        toastStore.trigger(t);
    }

    if (form?.error) {
		toastCreateError(form?.error);
    }

	let imageUrl = '';
	async function fileHandler(e: Event) {
		const {data: url, error: err} = await uploadFile(e);
		if (err) {
			toastCreateError('failed to upload file');
			return
		}
		imageUrl = url;
	}
</script>

<div class="flex justify-center h-full pt-5">
	<div class="card p-4 w-modal shadow-xl space-y-4">
		<form class="modal-form border border-surface-500 p-4 space-y-4 rounded-container-token"
		      method="POST" action="?/createCourse">
			<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
				<div class="col-span-1">
					<label class="label">
						<input class="input" type="text" name="name" value={form?.name ?? ''} placeholder="Course name..." />
					</label>
				</div>
				<div class="col-span-1">
					<label class="label mx-1">
						<select class="select variant-ghost-surface" name="language">
							<option value="english">English</option>
							<option value="russian">Russian</option>
						</select>
					</label>
				</div>
				<div class="col-span-2">
					<label class="label">
						<input class="input" type="number" name="price" value={form?.price ?? ''} placeholder="Course price..." />
					</label>
				</div>
				<div class="col-span-2">
					<label class="label">
						<textarea class="textarea" name="description" rows="3" placeholder="Course description...">{form?.description ?? ''}</textarea>
					</label>
				</div>
				<input type="hidden" name="image_url" value={form?.image_url ?? imageUrl} />
				<div class="label col-span-2">
					{#if imageUrl !== ''}
						<img src={imageUrl} alt="course"/>
					{/if}
					<FileDropzone name="file" on:change={fileHandler}>
						<svelte:fragment slot="message">Upload or drag and drop course image</svelte:fragment>
					</FileDropzone>
				</div>
			</div>
			<button class="input btn btn-md variant-ghost-primary" type="submit">CREATE</button>
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
