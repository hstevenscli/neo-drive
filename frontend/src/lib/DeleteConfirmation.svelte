<script>
    let { deleteModalActive = $bindable(''), files = $bindable(), fileToDelete, lastKeyPressed } = $props();

    $effect(() => {
        if (deleteModalActive === 'is-active') {
            console.log("inside delete component", lastKeyPressed);
            handleKey();
        }
    })

    function handleKey() {
        if (lastKeyPressed === 'y' || lastKeyPressed === 'Y') {
            console.log("Yes");
            deleteModalActive = '';
            let index = files.indexOf(fileToDelete);
            if (index > -1) {
                files.splice(index, 1);
            }
        }
        if (lastKeyPressed === 'n' || lastKeyPressed === 'N') {
            console.log("No");
            deleteModalActive = '';
        }
    }

</script>


<div class="modal { deleteModalActive }">
  <div class="modal-background"></div>
  <div class="modal-content">
      <p>Are you sure you want to delete: {fileToDelete}</p>
      <div class="buttons">
          <div class="button is-danger">
              (Y)es
          </div>
          <div class="button is-link">
              (N)o
          </div>
      </div>
  </div>
  <button onclick="{() => { deleteModalActive = '' }}" class="modal-close is-large" aria-label="close"></button>
</div>


<style>
    .modal {
        background-color: grey;
    }

</style>
