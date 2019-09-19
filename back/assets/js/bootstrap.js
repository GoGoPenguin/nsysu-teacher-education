let token = $.cookie('token')
let refreshToken = $.cookie('refresh-token')

if (token == undefined) {
    if (refreshToken != undefined) {
        renewToken()
        token = $.cookie('token')
    } else if (window.location.pathname != '/login.html') {
        location.href = '/login.html'
    }
} else if (window.location.pathname == '/login.html') {
    location.href = '/'
}