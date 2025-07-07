<script>
    import { setLKP, getLKP } from "./state.svelte";
    let { path, deleteModalActive = $bindable(''), files = $bindable(), fileToDelete, getDirectory, adjustIndex } = $props();

    $effect(() => {
        if (deleteModalActive === 'is-active') {
            console.log("inside delete component", getLKP());
            handleKey();
        }
    })

    async function deleteFile(filename) {
        let url = "http://localhost:8080/files" + path.join("") + "/" + filename;
        console.log("D URL:", url);
        let response = await fetch(url, {
            method: "DELETE",
        })
        let message = await response.json();
        console.log(message);
    }

    async function confirmDelete() {
        deleteModalActive = '';
        let index = files.indexOf(fileToDelete);
        if (index > -1) {
            files.splice(index, 1);
        }
        console.log("Deleting file");
        await deleteFile(fileToDelete);
        getDirectory(path.join(""));
        adjustIndex();
    }


    async function handleKey() {
        if (getLKP() === 'Escape') {
            deleteModalActive = '';
        }
        if (getLKP() === 'y' || getLKP() === 'Y') {
            confirmDelete();
        }
        if (getLKP() === 'n' || getLKP() === 'N') {
            deleteModalActive = '';
        }
    }

    function closeModal() {
        // use p because it doesnt do anything
        setLKP('p');
        deleteModalActive = '';
    }

</script>


<div class="modal { deleteModalActive } has-text-centered">
  <div class="modal-background"></div>
  <div class="modal-content">
    <p class="title is-3">Are you sure you want to delete:<br><span class="has-text-primary">{fileToDelete}</span></p>
    <div class="buttons is-flex is-justify-content-center">
        <button onclick={() => {confirmDelete(); closeModal()}} class="button is-danger">(Y)es</button>
      <button onclick={closeModal} class="button is-link">(N)o</button>
    </div>
  </div>
  <button onclick={closeModal} class="modal-close is-large" aria-label="close"></button>
</div>


<style>
    .modal {
        background-color: grey;
    }

</style>
