<script>
    import { onDestroy } from "svelte";
    let { previewFile, previewFileType = $bindable(), previewFileName } = $props();
    import { setLKP, getLKP } from "./state.svelte";
    let active = $state('is-active');
    let innerWidth = window.innerWidth;
    let innerHeight = window.innerHeight;
    console.log(innerWidth);

    $effect(() => {
        handleKey();
    })
    function handleKey() {
        let key = getLKP();
        if (key === 'Escape') {
            active = '';
            previewFileType = '';
        }
        if (key === 'v') {
            active = '';
            previewFileType = '';
            setLKP('p');
        }
    }

    onDestroy(() => {
        setLKP('p');
    })
</script>

<!-- 
make 2 different modals, one for images and pdfs
and one for text files

-->

{#snippet imgmodal()}
<p class="image is-4by3">
    <img src={previewFile} alt="Image Preview">
</p>
{/snippet}

{#snippet pdfmodal()}
<embed src={previewFile} width="{innerWidth}px" height="{innerHeight}px" />
{/snippet}

{#snippet htmlmodal()}
    <div class="box">
        {@html previewFile}
    </div>
{/snippet}

{#snippet textmodal()}
<div class="box">
    <div class="content">
        {#each previewFile.split("\n") as line }
            <p class="has-text-weight-bold">{line}</p>
        {/each}
    </div>
</div>
{/snippet}



<!-- CHANGE THE CONTENT BASED ON PREVIEW FILE TYPE -->
<div class="modal {active}">
    <div onclick={() => {active = ''; previewFileType = '';}} class="modal-background"></div>
    <div class="modal-content { previewFileType === "png" || previewFileType === "jpg" ? 'hide-scroll' : '' }">
        <p class="has-text-primary"><b>{ previewFileName }</b></p>
        <div class="box">
            {#if previewFileType === 'pdf'}
                {@render pdfmodal()}
            {:else if previewFileType === 'png' || previewFileType === 'jpg'}
                {@render imgmodal()}
            {:else if previewFileType === 'html'}
                {@render htmlmodal()}
            {:else}
                {@render textmodal()}
            {/if}

        </div>
    </div>
    <button onclick={() => {active = ''; previewFileType = '';}} class="modal-close is-large" aria-label="close"></button>
</div>

<style>
    .modal {
        width: 100%;
    }
    .modal-content {
        width: 80%;
        height: 90%;
    }
    .hide-scroll {
        overflow: hidden;
    }
</style>
