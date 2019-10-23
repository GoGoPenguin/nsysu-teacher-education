$('#loginform').on('submit', (e) => {
    e.preventDefault();

    let account = $('input#account').val()
    let password = $('input#password').val()

    $.ajax({
        url: `${config.server}/v1/login`,
        type: 'POST',
        data: {
            'Account': account,
            'Password': password,
            'Role': 'admin',
        },
        error: (xhr) => {
            alert('Unexcepted Error')
            console.error(xhr);
        },
        success: (response) => {
            if (response.code != 0) {
                alert(response.message)
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
});

$('button#logout').click(() => {
    $('#logoutModal').show()
})

$('#logoutModal button.btn.btn-primary').click(() => {
    $.ajax({
        url: `${config.server}/v1/logout`,
        type: 'POST',
        error: (xhr) => {
            console.error(xhr);
        },
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        success: (response) => {
            if (response.code != 0) {
                console.error(response.message)
            } else {
                let date = new Date()
                date.setTime(date.getTime() + (response.data.Expire * 1000));

                let cookies = $.cookie()
                for (var cookie in cookies) {
                    $.removeCookie(cookie)
                }

                location.href = '/login.html'
            }
        }
    });
})
