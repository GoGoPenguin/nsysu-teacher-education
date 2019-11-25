$('#login-form').on('submit', (e) => {
    e.preventDefault();

    let account = $('input.input100.account').val()
    let password = $('input.input100.password').val()

    $.ajax({
        url: `${config.server}/v1/login`,
        type: 'POST',
        data: {
            'Account': account,
            'Password': password,
            'Role': 'student',
        },
        beforeSend: () => {
            $('#login').html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp載入中...')
            $('#login').attr("disabled", true)
        },
        error: (xhr) => {
            $('div.flex-sb-m div.alert-danger').show('fast')
            setTimeout(() => {
                $('div.flex-sb-m div.alert-danger').hide('slow')
            }, 2000)
        },
        success: (response) => {
            if (response.code != 0) {
                $('div.flex-sb-m div.alert-danger').show('fast')
                setTimeout(() => {
                    $('div.flex-sb-m div.alert-danger').hide('slow')
                }, 2000)
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
        },
        complete: (jqXHR, textStatus) => {
            $('#login').html('送出')
            $('#login').attr("disabled", false)
        }
    });
})

$('div.flex-sb-m button.close').click(() => {
    $('div.flex-sb-m div.alert-danger').hide('slow')
})

$('#logoutModal button.btn.btn-primary').click(() => {
    $.ajax({
        url: `${config.server}/v1/logout`,
        type: 'POST',
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            removeCookie()
        },
        success: (response) => {
            removeCookie()
        }
    });
})

const errorHandle = (xhr, msg) => {
    if (xhr.status == 401) {
        setTimeout(removeCookie, 2000);

        swal({
            title: '',
            text: '登入逾時，或是已從其他裝置登入，即將跳回登入頁面。',
            icon: 'warning',
            timer: 2000,
            buttons: false,
        })
    } else {
        swal({
            title: '',
            text: msg,
            icon: "error",
            timer: 1000,
            buttons: false,
        })
    }
}

const setHeader = (xhr) => {
    let token = $.cookie('token')
    if (token == undefined) {
        renewToken()
        token = $.cookie('token')
    }

    xhr.setRequestHeader('Authorization', `Bearer ${token}`);
}