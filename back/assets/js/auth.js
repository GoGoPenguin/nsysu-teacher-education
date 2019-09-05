$(document).ready(function () {
    let refreshToken = $.cookie('refresh-token')

    if (refreshToken == undefined && window.location.pathname != '/login.html') {
        location.href = '/login.html'
    } else if (refreshToken != undefined && window.location.pathname == '/login.html') {
        location.href = '/'
    }
})

$('button#login').click(function () {
    let account = $('input#account').val()
    let password = $('input#password').val()

    $.ajax({
        url: config.server + '/v1/login',
        type: 'POST',
        data: {
            'Account': account,
            'Password': password,
            'Role': 'admin',
        },
        error: function (xhr) {
            error('Unexcepted Error')
            console.error(xhr);
        },
        success: function (response) {
            if (response.code != 0) {
                error(response.message)
            } else {
                let date = new Date()
                date.setTime(date.getTime() + (response.data.Expire * 1000));

                $.cookie('token', response.data.Token, {
                    expires: date,
                });
                $.cookie('account', account);
                $.cookie('refresh-token', response.data.RefreshToken)

                location.href = '/'
            }
        }
    });
})

$('button#logout').click(function () {
    $('#logoutModal').show()
})

$('#logoutModal button.btn.btn-primary').click(function () {
    $.ajax({
        url: config.server + '/v1/logout',
        type: 'POST',
        error: function (xhr) {
            console.error(xhr);
        },
        beforeSend: function (xhr) {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', 'Bearer ' + token);
        },
        success: function (response) {
            if (response.code != 0) {
                console.error(response.message)
            } else {
                let date = new Date()
                date.setTime(date.getTime() + (response.data.Expire * 1000));

                $.removeCookie('token')
                $.removeCookie('refresh-token')

                location.href = '/login.html'
            }
        }
    });
})

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
            let cookies = $.cookie();
            for (var cookie in cookies) {
                $.removeCookie(cookie);
            }

            location.href = '/login.html'
            console.error(xhr);
        },
        success: function (response) {
            if (response.code != 0) {
                let cookies = $.cookie();
                for (var cookie in cookies) {
                    $.removeCookie(cookie);
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