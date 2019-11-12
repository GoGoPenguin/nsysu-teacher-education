const loading = () => {
    $('#overlay').css('visibility', 'visible')
    $('body').toggleClass('overlay-open')
}

const unloading = () => {
    $('#overlay').css('visibility', 'hidden')
    $('body').toggleClass('overlay-open')
}