<script lang="ts">
    import type {School} from "$lib/types.ts";
	import {apiRequest} from "$lib/api/api.ts";
	import {error} from "@sveltejs/kit";

    export let data;
	let schools = data.schools as School[]

	async function deleteSchool(index: number) {
		const token = data.token;
		const {error: err} = await apiRequest('/schools/' + schools[index].id, 'delete', undefined, token);
		if (err) {
			throw error(500, {message: err.error})
		} else {
			schools.splice(index, 1);
			schools = [...schools];
		}
	}
</script>

<div class="container h-full mx-auto w-5/6 text-token py-3">
	<div class="text-4xl">
		<h1>MY SCHOOLS</h1>
	</div>
	<a class="btn btn-md variant-ghost-surface my-5" href="/me/schools/create">+ CREATE</a>
	<div class="grid grid-cols-1 md:grid-cols-1 gap-4">
		{#each schools as school, index}
			<div class="card bg-initial card-hover overflow-hidden">
				<header>
					<div class="grid grid-cols-3 p-4">
						<a href={'/me/schools/' + school.id} class="col-span-2 p-1">
							<p class="text-2xl">{school.name}</p>
						</a>
						<div class="col-span-1 flex justify-end">
							<button type="button" class="btn btn-md variant-ghost-error"
									on:click={() => deleteSchool(index)}>✖</button>
						</div>
					</div>
				</header>
				<hr class="opacity-50" />
				<footer class="p-4 flex justify-start items-center space-x-4">
					<p class="text-1xl">{school.description}</p>
				</footer>
			</div>
		{/each}
	</div>
</div>
