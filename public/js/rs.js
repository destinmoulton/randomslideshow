/**
 * randomslideshow js
 */
(function () {
    let defaultHeight = 200;
    document.addEventListener("DOMContentLoaded", function () {
        // code...
        const increase = document.getElementById("rs-icon-increase-thumb-size");
        const decrease = document.getElementById("rs-icon-decrease-thumb-size");
        increase.addEventListener("click", function (e) {
            resizeImages("plus");
        });
        decrease.addEventListener("click", function (e) {
            resizeImages("minus");
        });
    });

    function resizeImages(sizeDir) {
        console.log("resizeImages called");
        const thumbs = document.querySelectorAll(".rs-thumb");

        if (sizeDir === "plus") {
            defaultHeight = defaultHeight + 50;
        } else {
            defaultHeight = defaultHeight - 50;
        }
        for (let thumb of thumbs) {
            console.log(thumb);
            thumb.style.height = defaultHeight + "px";
        }
    }
})();
