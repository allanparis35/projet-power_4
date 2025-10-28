const audio = document.getElementById('game-audio');

// Sauvegarder l'état avant le refresh
window.addEventListener('beforeunload', function() {
    if (!audio.paused) {
        localStorage.setItem('musicPlaying', 'true');
        localStorage.setItem('musicTime', audio.currentTime);
    } else {
        localStorage.setItem('musicPlaying', 'false');
    }
});

// Restaurer l'état après le refresh
window.addEventListener('load', function() {
    if (localStorage.getItem('musicPlaying') === 'true') {
        const savedTime = parseFloat(localStorage.getItem('musicTime') || 0);
        audio.currentTime = savedTime;
        audio.play();
    }
});