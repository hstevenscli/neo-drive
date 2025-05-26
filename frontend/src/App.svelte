<script>
    import Counter from './lib/Counter.svelte';
    import Login from './lib/Login.svelte';
    import FileList from './lib/FileList.svelte';
    import NeoDriveLogo from './assets/NeoDriveLogo.png';

    let loggedIn = false;

</script>


{#if !loggedIn}
    <Login  bind:loggedIn />
{:else}
    <figure class="image is-64x64">
        <img src={NeoDriveLogo} alt="NeoDrive Logo"/>
    </figure>

    <Counter />
    <FileList />
{/if}



<p>logged In {loggedIn}</p>

<style>
    .cool {
        color: #00F0FF;
    }
</style>
