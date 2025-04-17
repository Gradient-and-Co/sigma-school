<script lang="ts">
    export let form;

    import { getToastStore, type ToastSettings } from '@skeletonlabs/skeleton';
    const toastStore = getToastStore();

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
</script>

<div class="flex justify-center h-full pt-16">
	<div class="card p-4 w-modal shadow-xl space-y-4">
		<form class="modal-form border border-surface-500 p-4 space-y-4 rounded-container-token"
		      method="POST" action="?/login">
			<label class="label">
				<span>Email</span>
				<input class="input" type="email" name="email" value={form?.email ?? ''} placeholder="Enter email..." />
			</label>
			<label class="label">
				<span>Password</span>
				<input class="input" type="password" name="password" placeholder="Enter password..." />
			</label>
			<button class="input btn btn-md variant-ghost-primary" type="submit" name="submit">Log In</button>
			<p>Have no account? <a class="anchor" href="/auth/register">Register here</a>.</p>
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
