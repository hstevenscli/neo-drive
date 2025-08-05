<script>

    import { setLKP, getLKP } from "./state.svelte";
    let { activeModalsObj = $bindable(), fileToEdit, editFileName, path } = $props();

    let newName = $state(fileToEdit);
    function closeModal() {
        setLKP('P');
        activeModalsObj.fileEditModal = false;
    }
    $effect(() => {
        if (activeModalsObj.fileEditModal) {
            handleKey();
        }
    })
    async function handleKey() {
        if (getLKP() === 'Escape') {
            closeModal();
        }
        if (getLKP() === 'Enter') {
            confirmFile();
        }
    }

    function confirmFile() {
        console.log("File name changed");
    }
</script>

<div class="modal is-active has-text-centered">
  <div class="modal-background"></div>
  <div class="modal-content">
    <p class="title is-3">New Filename:<br></p>
    <input class="input" type="text" bind:value={newName}>
    <div class="buttons is-flex is-justify-content-center mt-3">
        <button onclick={() => {
                console.log("Confirm");
                editFileName(path, newName);
            }} class="button is-link">Confirm</button>
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
