let theme = $state(document.documentElement.getAttribute('data-theme'));


export function getTheme() {
    return theme;
}

export function setTheme(mode) {
    theme = mode;
}
