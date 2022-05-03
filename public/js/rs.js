/**
 * randomslideshow js
 */
(function () {
    let defaultHeight = 200;
    let lightbox = {};
    let lazyLoadInstance = {};
    let currentImage = {};
    document.addEventListener("DOMContentLoaded", function () {
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

        initializeLazyLoader();
        initializeLightbox();
    });

    function initializeLazyLoader() {
        // Start lazyloader
        lazyLoadInstance = new LazyLoad({
            // Your custom settings go here
        });
    }

    function initializeLightbox() {
        if (lightbox.destroy !== undefined) {
            lightbox.destroy();
        }
        lightbox = {};
        // Start lightbox
        let lbopts = {
            overlayOpacity: 0.9,
            additionalHtml:
                "<div><i id='rs-lb-delete' class='fas fa-trash rs-lb-delete-icon'></i></div>",
        };
        lightbox = new SimpleLightbox("#rs-gallery-container a", lbopts);
        lightbox.refresh();
        lightbox.on("shown.simplelightbox", function (e) {
            setActiveLightboxImage(e);
            setupTrashButton();
        });
        lightbox.on("changed.simplelightbox", function (e) {
            setActiveLightboxImage(e);
        });
    }

    function shuffleThumbs() {
        // Scroll to top
        window.scrollTo(0, 0);

        var ul = document.querySelector("#rs-gallery-container");
        for (var i = ul.children.length; i >= 0; i--) {
            ul.appendChild(ul.children[(Math.random() * i) | 0]);
        }

        // Reset the lazyload and lightbox
        initializeLazyLoader();
        initializeLightbox();
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

    function setActiveLightboxImage(e) {
        let $el = document.getElementById(e.target.id);
        currentImage = {
            id: e.target.id,
            path: $el.dataset.fullpath,
        };
    }
    /**
     * Called after lightbox modal is shown
     */
    function setupTrashButton() {
        const $trash = document.getElementById("rs-lb-delete");

        $trash.addEventListener("click", function () {
            if (!window.confirm("Are you sure you want to delete this?")) {
                return;
            }
            const $el = document.getElementById(currentImage.id);
            let $next = $el.nextElementSibling;
            if ($next === null) {
                $next = $el.previousElementSibling;
            }

            apiDeletePicture(currentImage.path);
            lightbox.close();
            // Remove the element from the dom
            $el.remove();
            // Remove the trash icon; it will be rebuilt
            $trash.remove();

            if ($next !== null) {
                initializeLightbox();
                lightbox.open($next);
            }
        });
    }

    function apiDeletePicture(picPath) {
        const data = {
            action: "delete",
            picture_path: picPath,
        };
        const opts = {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        };
        const url = "/api/picture/";
        fetch(url, opts)
            .then((response) => response.json())
            .then((data) => console.log(data));
    }
})();
