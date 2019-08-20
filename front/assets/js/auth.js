$(document).ready(function () {
    let token = $.cookie('token')

    if (token == undefined && window.location.pathname != '/login.html') {
        location.href = '/login.html'
    } else if (token != undefined && window.location.pathname == '/login.html') {
        location.href = '/'
    }
})

$('button#login').click(function () {
    let account = $('input.input100.account').val()
    let password = $('input.input100.password').val()

    $.ajax({
        url: config.server + '/v1/login',
        type: 'POST',
        data: {
            'Account': account,
            'Password': password,
        },
        error: function (xhr) {
            console.error(xhr);
        },
        success: function (response) {
            if (response.code != 0) {
                console.error(response.message)
            } else {
                let date = new Date()
                date.setTime(date.getTime() + (response.data.Expire * 1000));

                $.cookie('token', response.data.Token, {
                    expires: date,
                });
                $.cookie('refresh-token', response.data.RefreshToken)

                location.href = '/'
            }
        }
    });
})

