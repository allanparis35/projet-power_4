const audio = document.getElementById('game-audio');


window.addEventListener('beforeunload', function() {
    if (!audio.paused) {
        localStorage.setItem('musicPlaying', 'true');
        localStorage.setItem('musicTime', audio.currentTime);
    } else {
        localStorage.setItem('musicPlaying', 'false');
    }
});


window.addEventListener('load', function() {
    if (localStorage.getItem('musicPlaying') === 'true') {
        const savedTime = parseFloat(localStorage.getItem('musicTime') || 0);
        audio.currentTime = savedTime;
        audio.play();
    }
});