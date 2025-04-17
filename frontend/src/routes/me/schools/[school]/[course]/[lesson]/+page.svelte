<script lang="ts">
    import {FileDropzone, Tab, TabGroup} from "@skeletonlabs/skeleton";
    import { markedHighlight } from "marked-highlight";
    import hljs from 'highlight.js';

    const marked = new Marked(
        markedHighlight({
            emptyLangClass: 'hljs',
            langPrefix: 'hljs language-',
            highlight(code, lang) {
                const language = hljs.getLanguage(lang) ? lang : 'plaintext';
                return hljs.highlight(code, { language }).value;
            }
        })
    );

    import type {Lesson} from "$lib/types.ts";

    export let form;
    export let data;

    import { getToastStore, type ToastSettings } from '@skeletonlabs/skeleton';
    import {Marked} from "marked";
    import {uploadFile} from "$lib/api/api.ts";
    const toastStore = getToastStore();

    const lesson = data.lesson as Lesson;

    let markdown = '';
    if (lesson.type === 'theory') {
        markdown = data.theory;
    } else if (lesson.type === 'practice') {
        markdown = data.task;
    }

    let tabSet = 0;

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

    function createOption() {
        lesson.tests[0].options = [...lesson.tests[0].options, ''];
    }

    function deleteOption(index: number) {
        lesson.tests[0].options.splice(index, 1);
        lesson.tests[0].options = [...lesson.tests[0].options];
    }
</script>

{#if lesson.type === 'theory'}
<div class="flex justify-center h-full pt-5">
    <div class="card p-4 w-modal shadow-xl space-y-4">
        <form class="modal-form border border-surface-500 p-4 space-y-4 rounded-container-token"
              method="POST" action="?/editTheoryLesson">
            <div class="col-span-1">
                <label class="label">
                    <input class="input" type="text" name="title" value={form?.title ?? lesson.title}
                           placeholder="Lesson title..." />
                </label>
            </div>
            <TabGroup>
                <Tab bind:group={tabSet} name="tab1" value={0}>EDITOR</Tab>
                <Tab bind:group={tabSet} name="tab2" value={1}>PREVIEW</Tab>
                <svelte:fragment slot="panel">
                    {#if tabSet === 0}
                        <textarea class="textarea variant-ghost-surface" name="theory" rows="12" bind:value={markdown} placeholder="Enter markdown here"></textarea>
                    {:else if tabSet === 1}
                        <div class="preview">{@html marked.parse(markdown)}</div>
                    {/if}
                </svelte:fragment>
            </TabGroup>
            <button class="input btn btn-md variant-ghost-primary" type="submit">SAVE</button>
        </form>
    </div>
</div>

{:else if lesson.type === 'practice'}
<div class="flex justify-center h-full pt-5">
    <div class="card p-4 w-modal shadow-xl space-y-4">
        <form class="modal-form border border-surface-500 p-4 space-y-4 rounded-container-token"
              method="POST" action="?/editPracticeLesson">
            <div class="col-span-1">
                <label class="label">
                    <input class="input" type="text" name="title" value={form?.title ?? lesson.title}
                           placeholder="Lesson title..." />
                </label>
            </div>
            <TabGroup>
                <Tab bind:group={tabSet} name="tab1" value={0}>TASK EDITOR</Tab>
                <Tab bind:group={tabSet} name="tab2" value={1}>PREVIEW</Tab>
                <svelte:fragment slot="panel">
                    {#if tabSet === 0}
                        <textarea class="textarea variant-ghost-surface" name="task" rows="4" bind:value={markdown} placeholder="Enter markdown here"></textarea>
                    {:else if tabSet === 1}
                        <div class="preview">{@html marked.parse(markdown)}</div>
                    {/if}
                </svelte:fragment>
            </TabGroup>
            <div>
                {#each lesson.tests[0].options as option, index}
                    <div class="col-span-2 variant-ghost-surface flex justify-between items-center px-2 my-2">
                        <input class="radio" type="radio" name="answer" value={option} bind:group={lesson.tests[0].answer} />
                        <input name="options[]" class="input variant-ghost-surface mx-2 my-2" type="text" bind:value={lesson.tests[0].options[index]}/>
                        <div class="flex justify-center">
                            <button type="button" class="btn btn-md variant-ghost-error mx-1 right-0" on:click={() =>
                            deleteOption(index)}>✖</button>
                        </div>
                    </div>
                {/each}
            </div>
            <button type="button" class="btn btn-md variant-ghost-surface mx-1" on:click={createOption}>ADD OPTION</button>
            <button class="input btn btn-md variant-ghost-primary" type="submit">SAVE</button>
        </form>
    </div>
</div>

{:else if lesson.type === 'video'}
    <div class="flex justify-center h-full pt-5">
        <div class="card p-4 w-modal shadow-xl space-y-4">
            <form class="modal-form border border-surface-500 p-4 space-y-4 rounded-container-token"
                  method="POST" action="?/editVideoLesson">
                <div class="col-span-1">
                    <label class="label">
                        <input class="input" type="text" name="title" value={form?.title ?? lesson.title}
                               placeholder="Lesson title..." />
                    </label>
                </div>
                <div class="col-span-1">
                    <label class="label">
                        <input class="input" type="text" name="video_url" value={form?.video_url ?? lesson.video_url}
                               placeholder="Video link..." />
                    </label>
                </div>
                <button class="input btn btn-md variant-ghost-primary" type="submit">SAVE</button>
            </form>
        </div>
    </div>

{/if}

<style>
    :global(.hljs) {
        background-color: #2f2826;
        margin-top: 1rem;
        margin-bottom: 1rem;
    }

    :global(p) {
        margin-top: 1rem;
        margin-bottom: 1rem;
    }

    :global(h1) {
        font-size: 2rem;
    }

    :global(h2) {
        font-size: 1.5rem;
    }

    :global(h3) {
        font-size: 1.2rem;
    }

    :global(h4) {
        font-size: 1.0rem;
    }
</style>