$(".user").mouseover(
    function () {
        $(".sub_menu").css("height", "165px");
    }
)

$(".user").mouseout(
    function () {
        $(".sub_menu").css("height", "0");
    }
)

$(document).ready(function () {
    if ($(window).width() < 600) {
        $(".logo").css("width", "130px");
        $(".logo-img").attr("src", "/assets/img/logo-mobile.svg");
    }
});

$(".nav_burger").click(
    function () {
        $(".mobile_menu").toggleClass("open");
        $(".burger_bar").toggleClass("open");
    }
)