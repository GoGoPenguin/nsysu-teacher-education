$(document).ready(function () {
    $('table#students').DataTable({
        processing: true,
        serverSide: true,
        ordering: false,
        searching: false,
        ajax: {
            url: config.server + '/v1/users',
            type: 'GET',
            dataSrc: 'list',
            beforeSend: function (xhr) {
                let token = $.cookie('token')
                if (token == undefined) {
                    renewToken()
                    token = $.cookie('token')
                }

                xhr.setRequestHeader('Authorization', 'Bearer ' + token);
            },
            error: function (xhr, error, thrown) {
                if (xhr.status == 401) {
                    let cookies = $.cookie()
                    for (var cookie in cookies) {
                        $.removeCookie(cookie)
                    }

                    location.href = '/login.html'
                } else {
                    alert(xhr.responseText)
                }
            }
        },
        columns: [
            { data: "Name" },
            { data: "Account" },
            { data: "CreatedAt" },
        ],
        language: {
            url: '/assets/languages/chinese.json'
        },
    });
})

$('input#upload').fileinput({
    language: 'zh-TW',
    theme: "fas",
    allowedFileExtensions: ['csv'],
    uploadUrl: config.server + '/v1/users',
    ajaxSettings: {
        headers: {
            'Authorization': 'Bearer ' + $.cookie('token')
        }
    },
}).on('fileuploaded', function (e, params) {
    if (params.response.code != 0) {
        $('div.kv-fileinput-error.file-error-message').html('\
            <button type="button" class="close kv-error-close" aria-label="Close">\
                <span aria-hidden="true">×</span>\
            </button>\
            <ul>\
                <li data-thumb-id="thumb-upload-20_students.csv" data-file-id="20_students.csv">\
                    <pre>'+ params.response.message + '</pre>\
                </li>\
            </ul>\
        ').show('fast')
    }
});

$('body').on('click', 'button.close.kv-error-close', function () {
    $('div.kv-fileinput-error.file-error-message').hide('slow')
})