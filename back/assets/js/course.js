$(document).ready(function () {
    $('table#course').DataTable({
        processing: true,
        serverSide: true,
        ordering: false,
        searching: false,
        ajax: {
            url: config.server + '/v1/course',
            type: 'GET',
            dataSrc: function (d) {
                d.list.forEach(function (element, index, array) {
                    let startDate = array[index].Start.substring(0, 10)
                    let startTime = array[index].Start.substring(11, 19)
                    let endDate = array[index].End.substring(0, 10)
                    let endTime = array[index].End.substring(11, 19)

                    if (startDate == endDate) {
                        array[index].Time = startDate + ' ' + startTime + ' ~ ' + endTime
                    } else {
                        array[index].Time = startDate + ' ' + startTime + ' ~ ' + endDate + ' ' + endTime
                    }
                })
                return d.list
            },
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
            { data: "Topic" },
            { data: "Time" },
            { data: "Information" },
            { data: "Type" },
        ],
        language: {
            url: '/assets/languages/chinese.json'
        },
    });
})