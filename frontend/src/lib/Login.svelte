<script>
    import {getLKP, setLKP} from "./state.svelte";

    let { loggedIn = $bindable() } = $props();
    let password = $state("monke");
    let showNotification = $state(false);
    let message = $state({
        1: "" ,
        2: "" ,
        3: "" ,
    });

    async function handleLogin() {
        console.log("Last key pressed", getLKP());
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
        password = '';
        console.log("Last key pressed", getLKP());
    }
    let i = 0;
    function addLetter() {
        let letters = "Welcome to NeoDrive";
        if (i < 11) {
            message[1] += letters[i];
        } else if ( i <= 13 ) {
            message[2] += letters[i];
        } else {
            message[3] += letters[i];
        }
        i++;
    }

    let id = setInterval(() => {

        if (i > 18) {
            clearInterval(id);
            return
        }
        addLetter();

    }, 100)
</script>

<p class="title is-1 ml-6 mt-6">{ message[1] } <span class="neo">{ message[2] } </span><span class="drive">{ message[3] } </span> <span class="cursor">&#9611;</span></p>
<div class="section is-medium">
    <p class="title is-1 has-text-centered">Login.</p>
</div>

<div class="columns">
    <div class="column is-4 is-offset-4">
        <div class="field">
            <label class="label">Password {#if showNotification }<span class="is-size-7 has-text-danger">INVALID PASSWORD</span>{/if}</label>
            <div class="control">
                <input onkeydown={(event) => {
                    if (event.key === "Enter") {
                        event.stopPropagation();
                        console.log("setting lkp to blank");
                        setLKP('');
                        handleLogin();
                    }
                }}
                class="input" type="password" bind:value={password} autofocus />
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
