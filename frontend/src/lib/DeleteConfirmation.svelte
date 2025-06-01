<script>
    let { deleteModalActive = $bindable(''), files = $bindable(), fileToDelete, lastKeyPressed, getFiles } = $props();

    $effect(() => {
        if (deleteModalActive === 'is-active') {
            console.log("inside delete component", lastKeyPressed);
            handleKey();
        }
    })

    async function deleteFile(filename) {
        let response = await fetch("http://localhost:8080/files/" + filename, {
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
        getFiles();
    }


    async function handleKey() {
        if (lastKeyPressed === 'Escape') {
            deleteModalActive = '';
        }
        if (lastKeyPressed === 'y' || lastKeyPressed === 'Y') {
            confirmDelete();
        }
        if (lastKeyPressed === 'n' || lastKeyPressed === 'N') {
            deleteModalActive = '';
        }
    }

</script>


{#if deleteModalActive === 'is-active'}
<div class="modal is-active has-text-centered">
  <div class="modal-background"></div>
  <div class="modal-content">
    <p class="title is-3">Are you sure you want to delete:<br><span class="has-text-primary">{fileToDelete}</span></p>
    <div class="buttons is-flex is-justify-content-center">
      <button on:click={confirmDelete} class="button is-danger">(Y)es</button>
      <button on:click={() => deleteModalActive = ''} class="button is-link">(N)o</button>
    </div>
  </div>
  <button on:click={() => deleteModalActive = ''} class="modal-close is-large" aria-label="close"></button>
</div>
{/if}


<style>
    .modal {
        background-color: grey;
    }

</style>
