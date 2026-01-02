<script>
    import FileUpload from './FileUpload.svelte';
    import DeleteConfirmation from './DeleteConfirmation.svelte';
    import FilePreview from './FilePreview.svelte';
    import HelpMenu from './HelpMenu.svelte';
    import PathView from './PathView.svelte';
    import { tick, untrack, onMount } from 'svelte';
    import { toggleTheme, setLKP, getLKP, getTheme, setTheme, toggleHelpModal, getShowHelpModal } from './state.svelte';
    import CreateDir from './CreateDir.svelte';
    import FileEdit from './FileEdit.svelte';

    let { keyPressedCounter} = $props();
    let deleteModalActive = $state('');
    let fileToDelete = $state('');
    let files = $state([]);
    let fileIndex = $state(0);
    let previewFileType = $state('');
    let previewFile = $state();
    let previewFileName = $state();
    let errorIndex = $state(-1)
    let fileEditModalActive = $state(false);
    let fileToEdit = $state('');

    // Path  and displayPath contain the same items, the only difference
    // being that each item in path starts with a / while displayPath elements
    // Don't have a / in them
    let path = $state([])
    let displayPath = $state([])

    let activeModalsObj = $state({
        fileEditModal: false,
        createDirModal: false,
    })

    // Returns true if any key is set to true in the activeModalsObj
    function checkActiveModal() {
        let keys = Object.keys(activeModalsObj);
        for (let i = 0; i < keys.length; i++) {
            if (activeModalsObj[keys[i]]) {
                return true
            }
        }
        return false
    }

    function genericOpenModal(targetModal) {
        // toggle modal states
        activeModalsObj[targetModal] = !activeModalsObj[targetModal];
    }

    function toggleCreateDirModalActive() {
        activeModalsObj.createDirModal = !activeModalsObj.createDirModal;
    }

    onMount( async () => {
        //let response = await fetch('http://localhost:8080/files');
        //let message = await response.json();
        await getDirectory();
        await tick();
        fileIndex = 0;
    })

    // Helper function for resetting the paths and getting the root directory
    function flushPath() {
        path = [];
        displayPath = [];
        getDirectory();
    }

    async function getDirectory(path="") {
        let url = `/dir${path}`
        console.log(url);
        let response = await fetch(url);
        let message = await response.json();
        console.log(message);
        files = message;
        files.sort((a, b) => {
            if (a.IsDir && !b.IsDir) return -1;
            if (!a.IsDir && b.IsDir) return 1;
            return 0;
        });
        console.log(files);
    }

    function moveIndexHighlight() {
        if (files.length === 1) {
            fileIndex = 0;
        }
        if (fileIndex <= files.length-1) {
            fileIndex = fileIndex
        } else {
            fileIndex = files.length -1;
        }
    }

    $effect(() => {
        // If the delete modal is open don't do key presses at this level
        if (deleteModalActive === 'is-active') {
            return
        }
        // If the create dir modal is open don't do key presses at this level
        if (activeModalsObj.createDirModal || activeModalsObj.fileEditModal) {
            return
        }
        // If the file preview is open don't do key presses at this level
        if (previewFileType !== '') {
            return
        }
        // prevent effect from activating itself and infinite looping
        if (keyPressedCounter) {
            untrack(() => handleKey());
        }
        // keep highlighted file in center screen
        setTimeout(() => {
            let selected = document.querySelector(`.index-${fileIndex}`);
            if (selected) {
                selected.scrollIntoView({ block: "center" });
            }
        })
        // so that itll update everytime even if key is the same
    })

    async function handleKey() {
        switch (getLKP()) {
            case "Enter":
                if (files[fileIndex].IsDir) {
                    // If dir, we need to:
                    // Add the directory name to the path
                    // Make a request for that path/dir
                    // make an actual apth and a display path so that you can do this better
                    path.push("/" + files[fileIndex].Name);
                    let p = path.join("");
                    displayPath.push(files[fileIndex].Name);
                    await getDirectory(p);
                    await tick();
                    fileIndex = 0;
                }
                break;
            case "-":
                // Go up  asingle directory by popping 1 elt from path
                // get directory using path
                path.pop();
                displayPath.pop();
                let p = path.join("/");
                console.log("Path", displayPath);
                getDirectory(p);
                break;
            case "Backspace":
                // Go up  asingle directory by popping 1 elt from path
                // get directory using path
                path.pop();
                displayPath.pop();
                let pbackspace = path.join("/");
                console.log("Path", displayPath);
                getDirectory(pbackspace);
                break;
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
                fileToDelete = files[fileIndex].Name;
                deleteModalActive = 'is-active';
                break;
            case "d":
                if (!files[fileIndex].IsDir) {
                    let p2 = path.join("")
                    downloadFile(p2 + "/" + files[fileIndex].Name);
                }
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
                toggleHelpModal();
                break;
            case "t":
                toggleTheme();
                break;
            case "v":
                setLKP('p');
                let p1 = path.join("")
                console.log(p1)
                readFile(p1 + "/" + files[fileIndex].Name);
                break;
            case ":":
                await tick();
                activeModalsObj.createDirModal = true;
                console.log("Command mode");
                break;
        }
        adjustIndex();
    }

    async function mkdir(path) {
        let url = "/dir/";
        let response = await fetch(url + path, {
            method: "POST"
        })
        let json = response.json();
        console.log(json);
        let p = this.path.join("");
        await getDirectory(p);
        await tick();
        console.log("Index before:", fileIndex);
        moveIndexHighlight();
        console.log("Index after:", fileIndex);
    }

    async function editFileName(path, newName) {
        // convert path into something usable
        console.log("Path:", path)
        console.log("NewName:", newName)
        console.log("New Name");
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
        tempLink.href = `/files/${filename}`;
        tempLink.download = filename;
        document.body.appendChild(tempLink);
        tempLink.click();
        document.body.removeChild(tempLink);
    }
    function deleteFile(filename) {
        console.log(`Deleting ${filename}`);
    }


    async function readFile(filename) {
        let response = await fetch("/view" + filename);
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
            default:
                if (contentType.slice(0,4) === "text") {
                    let othertext = await response.text();
                    previewFile = othertext;
                    previewFileType = "text"
                } else {
                    console.log("some kind of data");
                }
        }
        previewFileName = filename;

        if(response.status === 415) {
            setTimeout(() => {
                errorIndex = -1;
            }, 2000)
            errorIndex = fileIndex;
        }



    }
</script>

    {#if getShowHelpModal()}
        <HelpMenu ></HelpMenu>
    {/if}

    {#if activeModalsObj.createDirModal}
        <CreateDir bind:activeModalsObj { path } { mkdir }></CreateDir>
    {/if}
    {#if activeModalsObj.fileEditModal}
        <FileEdit bind:activeModalsObj { fileToEdit } { editFileName } { path } ></FileEdit>
    {/if}

    <DeleteConfirmation bind:deleteModalActive { path } { getDirectory } { files } { fileToDelete} { adjustIndex } />
    <FileUpload { getDirectory } { path }></FileUpload>
    <br>


    {#if previewFileType === '' && deleteModalActive != 'is-active' && !checkActiveModal()}
        <PathView bind:displayPath bind:path { flushPath } { getDirectory } { toggleCreateDirModalActive }></PathView>
    {/if}


    <p>{previewFileType}</p>
    {#if previewFileType !== ''}
        <FilePreview {previewFile} { previewFileName } bind:previewFileType />
    {/if}

    {#each files as file, i }
        <div onclick={() => {
            fileIndex = i;
            setLKP("Enter");
            handleKey();
            console.log("HEY");
            }} class="box is-flex is-align-items-center is-justify-content-space-between file-box index-{i} {fileIndex === i ? getTheme() === 'light' ? 'has-background-light-blue': 'has-background-link': ''}"> 
            <p class=" {getTheme() === "light" ? 'has-text-dark': 'has-text-light'}">
                {#if file.IsDir}
                    <span class="icon">
                        <i class="fa-solid fa-folder"></i>
                    </span>
                {/if}
                <b><i>{ file.Name }
                        <span onclick={(event) => { event.stopPropagation(); genericOpenModal("fileEditModal"); fileToEdit = files[i].Name;}} class="icon has-text-grey pen">
                            <i class="fa-solid fa-pen-nib"></i>
                        </span>
                </i></b></p>
            {#if errorIndex === i}
                <p class="help is-danger error-msg-{i} { getTheme() === 'light' ? 'has-background-white' : 'has-background-black' }"><b>Unsupported Media Type</b></p>
            {/if}
            <!-- consider making each box its own FileControl.svelte comoponent -->
            <div class="buttons">

            {#if !file.IsDir}
                <button onclick={() => {
                    setLKP('p');
                    let p1 = path.join("");
                    console.log(p1);
                    readFile(p1 + "/" + file.Name);
                }}
                class="button" aria-label="View">
                    <span class="icon">
                        <i class="fa-solid fa-eye"></i>
                    </span>
                </button>
                <button onclick={() => {
                    let p2 = path.join("")
                    downloadFile(p2 + "/" + file.Name);
                    document.activeElement.blur();
                    }}
                class="button" aria-label="Download">
                    <span class="icon">
                        <i class="fa-solid fa-download"></i>
                    </span>
                </button>
            {/if}

                <button onclick={() => {
                    setLKP('p');
                    fileToDelete = file.Name;
                    deleteModalActive = 'is-active';
                    document.activeElement.blur();
                }} 
                class="button is-danger" aria-label="Delete">
                    <span class="icon has-text-white">
                        <i class="fa-solid fa-trash-can"></i>
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
    .help {
        color: red;
        padding-left: 10px;
        padding-right: 10px;
        border-radius: 5px;
    }
    .fa-folder {
        color: #00c5d1;
    }

    .pen {
        cursor: pointer;
    }
</style>
