<script>
    let { getDirectory, path } = $props();

    let filename = $state('No File Selected');

    function handleFileChange(event) {
        const files = event.target.files;
        if (files && files.length > 0) {
            filename = files[0].name;
            console.log("files", files);
            if (files.length > 1) {
                filename = files[0].name + " + ...";
            }
        } else {
            filename = 'No File Selected';
        }
    }


    async function uploadFiles() {
        const input = document.getElementById("file-input");
        const files = input.files;
        if (!files.length) return;
            
        console.log("uploading files", files);

        const formData = new FormData();
        for (let i = 0; i < files.length; i++) {
            formData.append('files', files[i]);
        }


        if (path.length === 0) {
            path = ["/"]
        }
        let p = path.join("");
        let response = await fetch("/upload" + p, {
            method: "POST",
            body: formData,
        });
        let message = await response.json();
        console.log("Server Response", message);
        getDirectory(p);
        filename = 'No File Selected'
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
      <button id="file-upload-button" onclick={ uploadFiles } class="button">Submit</button>
      </label>
    </div>



<style>
    .file {
        margin-top: 2rem;
    }
</style>
