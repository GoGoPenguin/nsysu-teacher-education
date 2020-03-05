$(document).ready(function () {
    // Add smooth scrolling to all links
    $("a").on('click', function (event) {
        event.stopPropagation()

        // Make sure this.hash has a value before overriding default behavior
        if (this.hash !== "") {
            // Prevent default anchor click behavior
            event.preventDefault();

            // Store hash
            var hash = this.hash;

            // Using jQuery's animate() method to add smooth page scroll
            // The optional number (800) specifies the number of milliseconds it takes to scroll to the specified area
            $(".mobile_menu").removeClass("open");
            $(".burger_bar").removeClass("open");

            $('html, body').animate({
                scrollTop: $(hash).offset().top - $('#mainNav').height()
            }, 800);
        } // End if
        else if ($(this).attr('href') != undefined) {
            window.location = $(this).attr('href');
        }
    });

    $('.main_menu li').click(function () {
        $(this).children().click();
    });

    $('.mobile_menu li').click(function () {
        $(this).children().click();
    });
});