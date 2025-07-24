<script>
    import {getTheme} from "./state.svelte";

    let { displayPath = $bindable(), path = $bindable(), flushPath, getDirectory, toggleCreateDirModalActive } = $props()
    console.log("PATH", displayPath)
</script>

<div class="pathview-master-container { getTheme() === "light" ? 'background-light': 'background-dark' }">

{#each displayPath as pathNode}
    <span>{pathNode}/ </span>
{/each}


<div class="container is-flex is-align-items-center is-justify-content-space-between { getTheme() === "light" ? 'border-light': 'border-dark' }">
    <nav class="breadcrumb" aria-label="breadcrumbs">
        <ul class="is-flex is-align-items-center">
            <li><a onclick={ flushPath }>Home</a></li>
    {#each displayPath as pathNode, index}
            <li><a onclick="{() => {
                    let subarr = displayPath.slice(0, index + 1);
                    let newpath = path.slice(0, index+1);
                    path = newpath;
                    let s = subarr.join("/");
                    displayPath = subarr;
                    getDirectory("/" + s);
                }}" href="#">{pathNode}</a></li>
    {/each}
        </ul>
    </nav>

    <button onclick={toggleCreateDirModalActive} class="button">
        <span class="icon">
            <i class="fa-solid fa-folder-plus"></i>
        </span>
    </button>
</div>

</div>



<style>

    .background-dark {
        background: #15161B;
    }
    .background-light {
        background: #FEFFFE;
    }
    .border-light {
        border-bottom: 2px inset black;
    }
    .border-dark {
        border-bottom: 2px inset white;
    }
    .pathview-master-container {
        position: sticky;
        top: 0;
        z-index: 1000; /* Ensure it sits above content */
    }
    .container {
        width: 50%;
        padding: 10px 10px;
        margin-bottom: 1rem;
        border-radius: 5px;
    }
    .breadcrumb {
        margin-bottom: 0;
    }
    
</style>
