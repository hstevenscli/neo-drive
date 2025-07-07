<script>
    import {toggleHelpModal} from "./state.svelte";

</script>

<div class="modal is-active">
  <div onclick={() => {
      toggleHelpModal();
      }} class="modal-background"></div>
  <div class="modal-content">
    <!-- Any other Bulma elements you want -->
    <div class="box">
        <p>Movement: j,k, and arrow keys</p>
        <p>Download: d</p>
        <p>Delete: x</p>
        <p>Select files for upload: u</p>
        <p>Upload selected files: U</p>
        <p>Toggle file preview: v</p>
        <p>Toggle theme: t</p>
        <p>Close help menu: ?</p>
    </div>
  </div>
  <button onclick={() => {
      toggleHelpModal();
      }} class="modal-close is-large" aria-label="close"></button>
</div>


<style>
    .modal-background {
        opacity: 0.3;
    }
    .box {
        width: 45%;
        margin-left: 1em;
    }
    .modal-content {
        position: absolute;
        bottom: 2rem;
        left: 25%;
        transform: translateX(-50%);
    }
</style>
