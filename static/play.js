document.addEventListener("DOMContentLoaded", () => {
    const circles = document.querySelectorAll(".circle");

    circles.forEach(c => {
        c.addEventListener("click", () => {
            const col = c.dataset.col;

            fetch("/play", {
                method: "POST",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                body: "column=" + col
            }).then(() => {
                window.location.reload();
            });
        });
    });
});
