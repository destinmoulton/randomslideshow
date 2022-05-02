/**
 * randomslideshow js
 */
(function () {
    let defaultHeight = 200;
    let lightbox = {};
    let lazyLoadInstance = {};
    document.addEventListener("DOMContentLoaded", function () {
        // code...
        // Disable lightbox thumbs
        // Start lightbox
        lightbox = new SimpleLightbox("#rs-gallery-container a", {});
        //fsLightboxInstances["gallery"].props.disableThumbs = true;
        const shuffle = document.getElementById("rs-icon-shuffle");
        const increase = document.getElementById("rs-icon-increase-thumb-size");
        const decrease = document.getElementById("rs-icon-decrease-thumb-size");
        shuffle.addEventListener("click", function (e) {
            shuffleThumbs();
        });
        increase.addEventListener("click", function (e) {
            resizeThumbs("plus");
        });
        decrease.addEventListener("click", function (e) {
            resizeThumbs("minus");
        });

        initializeImageLibs();
    });

    function initializeImageLibs() {
        // Start lazyloader
        lazyLoadInstance = new LazyLoad({
            // Your custom settings go here
        });
        lightbox.refresh();
    }

    function shuffleThumbs() {
        // Scroll to top
        window.scrollTo(0, 0);

        var ul = document.querySelector("#rs-gallery-container");
        for (var i = ul.children.length; i >= 0; i--) {
            ul.appendChild(ul.children[(Math.random() * i) | 0]);
        }

        // Reset the lazyload and lightbox
        initializeImageLibs();
    }
    function resizeThumbs(sizeDir) {
        const thumbs = document.querySelectorAll(".rs-thumb");

        if (sizeDir === "plus") {
            defaultHeight = defaultHeight + 50;
        } else {
            defaultHeight = defaultHeight - 50;
        }
        for (let thumb of thumbs) {
            thumb.style.height = defaultHeight + "px";
        }
    }
})();
