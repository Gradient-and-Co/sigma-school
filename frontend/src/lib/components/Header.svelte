<script lang="ts">
	export let session;
	import {AppBar, Avatar} from "@skeletonlabs/skeleton";
</script>

<AppBar>
	<div slot="lead">
		<strong class="text-xl uppercase hidden md:flex">Eschool</strong>
		<div class="sm:hidden">
			<nav class="list-nav">
				<!-- (optionally you can provide a label here) -->
				<ul>
					<li>
						<a href="/">
							<span class="flex-auto">CATALOG</span>
						</a>
					</li>
					<li>
						<a href="/me/courses">
							<span class="flex-auto">LEARNING</span>
						</a>
					</li>
					<li>
						<a href="/me/schools">
							<span class="flex-auto">TEACHING</span>
						</a>
					</li>
					{#if !session}
						<li>
							<a href="/auth/login">
								<span class="flex-auto">LOGIN</span>
							</a>
						</li>
						<li>
							<a href="/auth/register">
								<span class="flex-auto">REGISTER</span>
							</a>
						</li>
					{:else}
						<form method="POST" action="/me?/logout">
							<button class="" type="submit">LOGOUT</button>
						</form>
						<a href="/me">
							<Avatar
									border="border-4 border-surface-300-600-token hover:!border-primary-500"
									cursor="cursor-pointer"
									width="w-10"
									initials="PS"
									fontSize="200"
							/>
						</a>
					{/if}
				</ul>
			</nav>
		</div>
	</div>
	<div slot="default" class="hidden md:flex">
		<a class="btn btn-md variant-soft-primary mx-0.5" href="/">
			CATALOG
		</a>
		<a class="btn btn-md variant-ghost-surface mx-0.5" href="/me/courses">
			LEARNING
		</a>
		<a class="btn btn-md variant-ghost-surface mx-0.5" href="/me/schools">
			TEACHING
		</a>
	</div>
	<div slot="trail" class="hidden md:flex">
		{#if !session}
			<a class="btn btn-md variant-ghost-surface mx-0.5" href="/auth/login">LOGIN</a>
			<a class="btn btn-md variant-ghost-surface mx-0.5" href="/auth/register">REGISTER</a>
		{:else}
			<form method="POST" action="/me?/logout">
				<button class="btn btn-md variant-ghost-surface mx-1" type="submit">LOGOUT</button>
			</form>
			<a href="/me">
				<Avatar
						border="border-4 border-surface-300-600-token hover:!border-primary-500"
						cursor="cursor-pointer"
						width="w-10"
						initials="PS"
						fontSize="200"
				/>
			</a>
		{/if}
	</div>
</AppBar>
