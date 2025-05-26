<script>
    let { loggedIn = $bindable() } = $props();
    let password = $state("");
    let showNotification = $state(false);

    async function handleLogin() {
        console.log("logged in with password:", password);
        let body = JSON.stringify({ password: password });
        console.log(body);
        let response = await fetch("http://localhost:8080/login", {
            method: "POST",
            body: body,
        });
        let message = await response.json();
        console.log(message);
        if (response.status === 200) {
            loggedIn = true;
        }
        if (!response.ok) {
            console.log("Show notification")
            showNotification = true;
        }
    }
</script>

<p class="title is-1 ml-6 mt-6">Welcome to <span class="neo">Neo</span><span class="drive">Drive</span> <span class="cursor">&#9611;</span></p>
<div class="section is-medium">
    <p class="title is-1 has-text-centered">Login.</p>
</div>

<div class="columns">
    <div class="column is-4 is-offset-4">
        <div class="field">
            <label class="label">Password {#if showNotification }<span class="is-size-7 has-text-danger">INVALID PASSWORD</span>{/if}</label>
            <div class="control">
                <input onkeydown={(event) => {
                    if (event.key == "Enter") {
                        handleLogin();
                    }
                }}
                class="input" type="password" bind:value={password} autofocus />
<!-- {#if showNotification }
            <p class="help is-danger has-text-weight-bold">INVALID CREDENTIALS</p>
            {/if}
-->
            </div>
            <p class="control">
                <button
                    onclick={handleLogin}
                    class="button is-info is-fullwidth mt-2 is-link is-outlined">
                    <b>Submit</b>
                </button>
            </p>
        </div>
    </div>
</div>

<!--
{#if showNotification }
    <div class="notification is-danger">
        <button class="delete" onclick={() => { showNotification=false }}></button>
        <div class="is-flex is-justify-content-center">
            Invalid Credientials
        </div>
    </div>
{/if}
-->

<style>
    .button {
        width: 85%;
        margin-left: 7.5%;
    }
    .notification {
        width: 50%;
        margin-left: 25%;
    }
    .cursor {
        animation: blink 1.5s steps(1) infinite;
    }

    @keyframes blink {
        50% {
            opacity: 0;
        }

    }

    .neo {
        color: #00c5d1;

    }

    .drive {
        color: #ff3e7a;

    }


</style>
