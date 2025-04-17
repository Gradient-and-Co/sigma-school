<script lang="ts">
    import type {Course, School} from "$lib/types.ts";

    export let data;
    const courses = data.courses as Course[]
	const schools = data.schools as Map<string, School>
</script>

<div class="container h-full mx-auto w-5/6 text-token py-3">
	<div class="text-4xl">
		<h1>MY COURSES</h1>
	</div>
	<div class="grid grid-cols-1 md:grid-cols-3 gap-4">
		{#each courses as course}
		<a class="card bg-initial card-hover overflow-hidden my-3" href={"/me/courses/" + course.id}>
			<header>
				<div class="grid grid-cols-3 p-4">
					<div class="col-span-2 p-1">
						<p class="text-2xl">{course.name}</p>
						<a class="text-1xl" href="/">{schools.get(course.id)?.name}</a>
					</div>
					<img src={course.image_url} class="bg-black/50 w-full aspect-[9/9]" alt="course avatar" />
				</div>
			</header>
			<hr class="opacity-50" />
			<footer class="p-4 flex justify-start items-center space-x-4">
				<div class="flex-auto flex justify-between items-center">
					<small>Language: {course.language}</small>
				</div>
			</footer>
		</a>
		{/each}
	</div>
</div>
