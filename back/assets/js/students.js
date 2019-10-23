const studentTable = $('table#students').DataTable({
    processing: true,
    serverSide: true,
    ordering: false,
    searching: false,
    ajax: {
        url: `${config.server}/v1/users`,
        type: 'GET',
        dataSrc: function (d) {
            d.list.forEach(function (element, index, array) {
                array[index].CreatedAt = element.CreatedAt.substring(0, 19)
            })
            return d.list
        },
        beforeSend: function (xhr) {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        error: function (xhr, error, thrown) {
            if (xhr.status == 401) {
                let cookies = $.cookie()
                for (var cookie in cookies) {
                    $.removeCookie(cookie)
                }

                location.href = '/login.html'
            } else {
                swal({
                    title: '',
                    text: xhr.responseText,
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
            }
        }
    },
    columns: [
        { data: "Name" },
        { data: "Account" },
        { data: "Major" },
        { data: "Number" },
        { data: "CreatedAt" },
    ],
    language: {
        url: '/assets/languages/chinese.json'
    },
});

$('input#upload').fileinput({
    language: 'zh-TW',
    theme: "fas",
    allowedFileExtensions: ['csv'],
    uploadUrl: `${config.server}/v1/users`,
    ajaxSettings: {
        headers: {
            'Authorization': `Bearer ${$.cookie('token')}`,
        }
    },
}).on('fileuploaded', (event, previewId, index, fileId) => {
    swal({
        title: '',
        text: '成功',
        icon: "success",
        timer: 1500,
        buttons: false,
    })
    studentTable.ajax.reload()
}).on('fileuploaderror', function (event, data, msg) {
    swal({
        title: '',
        text: '新增失敗',
        icon: "error",
        timer: 1500,
        buttons: false,
    })
    $('div.kv-upload-progress.kv-hidden').css({ 'display': 'none' })
})