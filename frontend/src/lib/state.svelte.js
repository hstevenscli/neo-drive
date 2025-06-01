let theme = $state(document.documentElement.getAttribute('data-theme'));
let lastKeyPressed = $state();
export let keyPressedCounter = 0;


export function getLKP() {
    return lastKeyPressed;
}

export function setLKP(key) {
    lastKeyPressed = key;
}

export function getTheme() {
    return theme;
}

export function setTheme(mode) {
    theme = mode;
}
