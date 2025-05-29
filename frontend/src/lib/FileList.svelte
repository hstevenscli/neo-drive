<script>
    import FileUpload from './FileUpload.svelte';
    import DeleteConfirmation from './DeleteConfirmation.svelte';
    import { untrack } from 'svelte';

    let { lastKeyPressed , keyPressedCounter} = $props();

    let deleteModalActive = $state('');
    let fileIndex = $state(0);
    let fileToDelete = $state('');
    let files = $state([
        "file1.txt",
        "banana.txt",
        "receipts.png",
        "video.mp4",
        "myfile.wav",
        "anotherone.png",
        "anotherone(1).png",
        "wacky.pdf",
        "file.txt",
    ])

    $effect(() => {
        if (deleteModalActive === 'is-active') {
            return
        }
        if (keyPressedCounter) {
            untrack(() => handleKey());
        }
        // so that itll update everytime even if key is the same
    })

    function handleKey() {
        switch (lastKeyPressed) {
            case "j":
                decreaseIndex();
                break;
            case "ArrowDown":
                decreaseIndex();
                break;
            case "k":
                increaseIndex();
                break;
            case "ArrowUp":
                increaseIndex();
                break;
            case "d":
                fileToDelete = files[fileIndex];
                deleteModalActive = 'is-active';
                deleteFile(files[fileIndex]);
                break;
            case "Enter":
                downloadFile(files[fileIndex]);
                break;
            case "u":
                // Gets the file upload form from FileUpload.svelte
                let fileInput = document.getElementById("file-input");
                fileInput.click();
                break;
            case "?":
                console.log("Opening Help Menu");
                break;
        }
    }

    function decreaseIndex() {
        fileIndex = Math.min(files.length - 1, fileIndex + 1);
    }
    function increaseIndex() {
        fileIndex = Math.max(0, fileIndex - 1);
    }

    function downloadFile(filename) {
        console.log(`Downloading ${filename}`);
    }
    function deleteFile(filename) {
        console.log(`Deleting ${filename}`);
    }

</script>

    <DeleteConfirmation bind:deleteModalActive { files } { fileToDelete} { lastKeyPressed } />

    <FileUpload></FileUpload>

    {#each files as file, i }
        <div class="box is-flex is-align-items-center is-justify-content-space-between file-box index-{i} {fileIndex === i ? 'has-background-link': ''}"> <b class="has-text-white">{file}</b>
            <div class="buttons">
                <button onclick={() => {
                    downloadFile(file);
                    document.activeElement.blur();
                    }}
                class="button" aria-label="Download">
                    <span class="icon">
                        <i class="fa-solid fa-download"></i>
                    </span>
                </button>
                <button onclick={() => {
                    deleteFile(file);
                    document.activeElement.blur();
                }} 
                class="button is-danger" aria-label="Delete">
                    <span class="icon has-text-white">
                        <i class="fa-solid fa-x"></i>
                    </span>
                </button>
            </div>
        </div>
    {/each}
    


<style>
    .file-box {
        width: 50%;
        margin-left: 25%;
    }
</style>
