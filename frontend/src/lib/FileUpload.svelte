<script>
    let filename = $state('');
    function handleFileChange(event) {
        const files = event.target.files;
        if (files && files.length > 0) {
            filename = files[0].name;
            console.log("files", files);
            if (files.length > 1) {
                filename = files[0].name + " + ...";
            }
        } else {
            filename = '';
        }
    }
</script>

    <div class="file has-name is-boxed is-flex is-justify-content-center">
      <label class="file-label">
        <input onchange="{handleFileChange}" class="file-input" id="file-input" type="file" multiple />
        <span class="file-cta">
          <span class="file-icon">
            <i class="fas fa-upload"></i>
          </span>
          <span class="file-label"> Choose a file… </span>
        </span>
        <span class="file-name"> {filename} </span>
      </label>
    </div>



<style>

</style>
