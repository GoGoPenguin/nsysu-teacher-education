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
                xhr.setRequestHeader('Authorization', 'Bearer ' + token);
            }
        },
        columns: [
            { data: "Name" },
            { data: "Account" },
            { data: "CreatedAt" },
        ],
        language: {
            url: 'assets/languages/chinese.json'
        },
    });
})