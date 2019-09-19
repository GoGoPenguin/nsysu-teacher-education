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

function renewToken() {
    let account = $.cookie('account')
    let refreshToken = $.cookie('refresh-token')

    $.ajax({
        url: config.server + '/v1/renew-token',
        type: 'POST',
        async: false,
        cache: false,
        data: {
            Account: account,
            RefreshToken: refreshToken,
        },
        error: function (xhr) {
            let cookies = $.cookie()
            for (var cookie in cookies) {
                $.removeCookie(cookie)
            }

            location.href = '/login.html'
        },
        success: function (response) {
            if (response.code != 0) {
                let cookies = $.cookie();
                for (var cookie in cookies) {
                    $.removeCookie(cookie)
                }

                location.href = '/login.html'
            } else {
                let date = new Date()
                date.setTime(date.getTime() + (response.data.Expire * 1000));

                $.cookie('token', response.data.Token, {
                    expires: date,
                });
            }
        }
    });
}