<script>
    import { onMount } from 'svelte';
    import Counter from './lib/Counter.svelte';
    import Login from './lib/Login.svelte';
    import FileList from './lib/FileList.svelte';
    import Navbar from './lib/Navbar.svelte';
    import { setLKP, getLKP } from './lib/state.svelte';


    let loggedIn = false;
    let keyPressedCounter = 0;

    function handleKeyDown (event) {
        console.log("Key pressed", event.key);
        setLKP(event.key);
        keyPressedCounter++;
    }

    onMount(() => {
        window.addEventListener("keydown", handleKeyDown);
        return () => window.removeEventListener("keydown", handleKeyDown);
    })


</script>

<Navbar />

{#if !loggedIn}
    <Login  bind:loggedIn />

{:else}
    <Counter />
    <FileList {keyPressedCounter}/>
{/if}


<p>logged In {loggedIn}</p>

<style>
    .cool {
        color: #00F0FF;
    }
</style>
