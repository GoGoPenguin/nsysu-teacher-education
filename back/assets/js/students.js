$(document).ready(function () {
    var index = 0

    $('table#students').DataTable({
        processing: true,
        serverSide: true,
        ajax: {
            url: config.server + '/v1/users',
            type: 'GET',
            data: function (d) {
                d.Index = index
                d.Count = 2
            },
            dataSrc: function (response) {
                index = response.data.LastID
                return response.data.List;
            },
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
    });
})