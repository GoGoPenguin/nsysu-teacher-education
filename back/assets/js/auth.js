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
        beforeSend: (xhr) => {
            $('#loginform button').html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp載入中...')
            $('#loginform button').attr("disabled", true)
        },
        error: (xhr) => {
            swal({
                title: '',
                text: '系統錯誤',
                icon: "error",
                timer: 2000,
                buttons: false,
            })
        },
        success: (response) => {
            if (response.code != 0) {
                swal({
                    title: '',
                    text: '帳號或密碼錯誤',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
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
        complete: () => {
            $('#loginform button').html('送出')
            $('#loginform button').attr("disabled", false)
        }
    })
});

$('button#logout').click(() => {
    $('#logoutModal').show()
})

$('#logoutModal button.btn.btn-primary').click(() => {
    $.ajax({
        url: `${config.server}/v1/logout`,
        type: 'POST',
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
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