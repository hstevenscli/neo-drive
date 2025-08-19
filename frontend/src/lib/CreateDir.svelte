<script>
    import {onMount, tick} from "svelte";
    import {getLKP, setLKP} from "./state.svelte";

    let { mkdir, path, activeModalsObj = $bindable() } = $props();
    let input = $state("");
    let inputEl;
    onMount(async () => {
        setTimeout(() => {
            inputEl.focus();
        }, 1)
    })
    function closeModal() {
        setLKP('P');
        activeModalsObj.createDirModal = false;
    }
    $effect(() => {
        if (activeModalsObj.createDirModal) {
            handleKey();
        }
    })
    async function handleKey() {
        if (getLKP() === 'Escape') {
            closeModal();
        }
        if (getLKP() === 'Enter') {
            confirmDir();
        }
    }
    function confirmDir() {
        let p = path.join("")
        console.log("confirming");
        closeModal();
        console.log("Path:", path)
        console.log("p:", p);
        console.log("input:", input);
        console.log("P+input:", p + input);
        let string = p + "/" + input;
        console.log("EXACT STRIGN", string);
        mkdir(string);
    }

</script>

<div class="modal is-active has-text-centered">
  <div class="modal-background"></div>
  <div class="modal-content">
    <p class="title is-3">Name of new directory to create:<br></p>
    <input bind:this={inputEl} class="input" type="text" bind:value={input}>
    <div class="buttons is-flex is-justify-content-center mt-3">
        <button onclick={confirmDir} class="button is-link">Confirm</button>
        <button onclick={(event) => {
            event.stopPropagation();
            closeModal();
            }} class="button is-danger">Cancel</button>
    </div>
  </div>
  <button onclick={closeModal} class="modal-close is-large" aria-label="close"></button>
</div>


<style>

</style>
