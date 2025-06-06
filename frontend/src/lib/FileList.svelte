<script>
    import FileUpload from './FileUpload.svelte';
    import DeleteConfirmation from './DeleteConfirmation.svelte';
    import FilePreview from './FilePreview.svelte';
    import { tick, untrack } from 'svelte';
    import { setLKP, getLKP, getTheme, setTheme } from './state.svelte';
    import { onMount } from 'svelte';
    import { toggleTheme } from './state.svelte';
    import { fade } from 'svelte/transition';
    import HelpMenu from './HelpMenu.svelte';

    let { keyPressedCounter} = $props();

    let deleteModalActive = $state('');
    let fileToDelete = $state('');
    let files = $state('');
    let fileIndex = $state(0);
    let previewFileType = $state('');
    let previewFile = $state();
    let showHelpModal = $state(false);
    //   let files = $state([
    //       "file1.txt",
    //       "banana.txt",
    //       "receipts.png",
    //       "video.mp4",
    //       "myfile.wav",
    //       "anotherone.png",
    //       "anotherone(1).png",
    //       "wacky.pdf",
    //       "file.txt",
    //   ])

    onMount( async () => {
        //let response = await fetch('http://localhost:8080/files');
        //let message = await response.json();
        await getFiles();
        await tick();
        fileIndex = 0;
    })

    async function getFiles(path="") {
        let response = await fetch(`http://localhost:8080/files${path}`);
        let message = await response.json();
        files = message.files;
    }

    $effect(() => {
        if (deleteModalActive === 'is-active') {
            return
        }
        if (previewFileType !== '') {
            return
        }
        if (keyPressedCounter) {
            untrack(() => handleKey());
        }
        setTimeout(() => {
            let selected = document.querySelector(`.index-${fileIndex}`);
            if (selected) {
                selected.scrollIntoView({ block: "center" });
            }
        })
        // so that itll update everytime even if key is the same
    })

    function handleKey() {
        switch (getLKP()) {
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
            case "g":
                fileIndex = 0;
                break;
            case "G":
                fileIndex = files.length - 1;
                break;
            case "x":
                fileToDelete = files[fileIndex];
                deleteModalActive = 'is-active';
                break;
            case "d":
                downloadFile(files[fileIndex]);
                break;
            case "u":
                // Gets the file upload form from FileUpload.svelte
                let fileInput = document.getElementById("file-input");
                fileInput.click();
                break;
            case "U":
                let uploadButton = document.getElementById("file-upload-button");
                uploadButton.click();
                break;
            case "?":
                console.log("Opening Help Menu");
                showHelpModal = !showHelpModal;
                break;
            case "t":
                toggleTheme();
                break;
            case "v":
                setLKP('p');
                readFile(files[fileIndex]);
                break;
        }
        adjustIndex();
    }

    function decreaseIndex() {
        fileIndex = Math.min(files.length - 1, fileIndex + 1);
    }
    function increaseIndex() {
        fileIndex = Math.max(0, fileIndex - 1);
    }
    function adjustIndex() {
        if (fileIndex < 0) {
            fileIndex = 0;
        }
        if (fileIndex > files.length - 1) {
            fileIndex = files.length -1;
        }
    }

    function downloadFile(filename) {
        const tempLink = document.createElement("a");
        tempLink.href = `http://localhost:8080/files/${filename}`;
        tempLink.download = filename;
        document.body.appendChild(tempLink);
        tempLink.click();
        document.body.removeChild(tempLink);
    }
    function deleteFile(filename) {
        console.log(`Deleting ${filename}`);
    }


    async function readFile(filename) {
        let response = await fetch("http://localhost:8080/view/" + filename);
        let contentType = response.headers.get("Content-Type");
        console.log("Type:", contentType);

        switch (contentType) {
            case "application/pdf":
                console.log("pdf");
                let pdfblob = await response.blob();
                previewFile = URL.createObjectURL(pdfblob);
                previewFileType = "pdf";
                break;
            case "image/jpeg":
                console.log("jpg");
                let jpgblob = await response.blob();
                previewFile = URL.createObjectURL(jpgblob);
                previewFileType = "jpg";
                break;
            case "text/html; charset=utf-8":
                console.log("html");
                let htmltext = await response.text();
                previewFile = htmltext;
                previewFileType = "html";
                break;
            case "text/plain; charset=utf-8":
                console.log("plain text");
                let plaintext = await response.text();
                console.log(plaintext);
                previewFile = plaintext;
                previewFileType = "txt";
                
                break;
            case "image/png":
                console.log("png");
                let pngblob = await response.blob();
                previewFile = URL.createObjectURL(pngblob);
                previewFileType = "png";
                break;
            case "text/markdown; charset=utf-8":
                console.log("md");
                console.log("plain text");
                let mdtext = await response.text();
                previewFile = mdtext;
                previewFileType = "md";
                break;

        }
    }
</script>

    {#if showHelpModal}
        <HelpMenu bind:showHelpModal></HelpMenu>
    {/if}
    <DeleteConfirmation bind:deleteModalActive { getFiles } { files } { fileToDelete} />
    <FileUpload { getFiles }></FileUpload>
    <br>

    <p>{previewFileType}</p>
    {#if previewFileType !== ''}
        <FilePreview {previewFile} bind:previewFileType />
    {/if}

    {#each files as file, i }
        <div class="box is-flex is-align-items-center is-justify-content-space-between file-box index-{i} {fileIndex === i ? getTheme() === 'light' ? 'has-background-light-blue': 'has-background-link': ''}"> 
            <p class=" {getTheme() === "light" ? 'has-text-dark': 'has-text-light'}"><b><a href="{"http://localhost:8080/files/" + file }" download>{ file }</a></b></p>
            <!-- consider making each box its own FileControl.svelte comoponent -->
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
                    setLKP('p');
                    fileToDelete = file;
                    deleteModalActive = 'is-active';
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
    .has-background-light-blue {
        background-color: #ADEAFF;
    }
    a {
        color: inherit;
        text-decoration: inherit;
    }
</style>
