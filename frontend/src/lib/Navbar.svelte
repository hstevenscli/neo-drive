<script>
    import NeoDriveLogo from '../assets/NeoDriveLogo.png';
    import { getTheme, setTheme } from './state.svelte';

    function toggleTheme() {
        if (getTheme() === 'dark') {
            setTheme('light');
        } else {
            setTheme('dark')
        }
        document.documentElement.setAttribute('data-theme', getTheme());
    }




</script>

<nav class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="image is-64x64">
        <img src={NeoDriveLogo} alt="NeoDrive Logo"/>
    </a>


    <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>

  <div id="navbarBasicExample" class="navbar-menu">
    <div class="navbar-start">
    </div>

    <div class="navbar-end">

      <div class="navbar-item">
        <div class="buttons">

            <button class="button" onclick={ toggleTheme }>
                <span class="icon">
                {#if getTheme() === "dark"}
                    <i class="fa-solid fa-sun"></i>
                {:else}
                    <i class="fa-solid fa-moon"></i>
                {/if}
                </span>
            </button>

        </div>
      </div>
    </div>
  </div>
</nav>



<style>

</style>
