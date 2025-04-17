<script lang="ts">
    export let form;

    import { getToastStore, type ToastSettings } from '@skeletonlabs/skeleton';
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
</script>

<div class="flex justify-center h-full pt-5">
	<div class="card p-4 w-modal shadow-xl space-y-4">
		<form class="modal-form border border-surface-500 p-4 space-y-4 rounded-container-token"
		      method="POST" action="?/createSchool">
			<label class="label">
				<span>Name</span>
				<input class="input" type="text" name="name" value={form?.name ?? ''} placeholder="Enter name..." />
			</label>
			<label class="label">
				<span>Description</span>
				<input class="input" type="text" name="description" value={form?.description ?? ''} placeholder="Enter description..." />
			</label>
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
