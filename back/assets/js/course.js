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
        columnDefs: [
            { className: "info", targets: [2] },
        ],
        language: {
            url: '/assets/languages/chinese.json'
        },
    });

    $('#start').datetimepicker({
        format: 'YYYY-MM-DD HH:mm:00',
        locale: 'zh-tw',
        initialDate: new Date(),
        autoclose: true,
        icons: {
            time: "fas fa-clock",
            date: "fa fa-calendar",
            up: "fas fa-angle-up",
            down: "fas fa-angle-down",
        }
    })

    $('#end').datetimepicker({
        format: 'YYYY-MM-DD HH:mm:00',
        locale: 'zh-tw',
        initialDate: new Date(),
        autoclose: true,
        icons: {
            time: "fas fa-clock",
            date: "fa fa-calendar",
            up: "fas fa-angle-up",
            down: "fas fa-angle-down",
        }
    })
})

$('table#course').on('click', 'td.info', function () {
    let filename = $(this).text()

    $.ajax({
        url: config.server + '/v1/course/' + filename,
        type: 'GET',
        xhrFields: {
            responseType: "blob"
        },
        error: function (xhr) {
            alert('Unexcepted Error')
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
            let a = document.createElement('a');
            let url = window.URL.createObjectURL(response);
            a.href = url;
            a.download = filename;
            document.body.append(a);
            a.click();
            a.remove();
            window.URL.revokeObjectURL(url);
        }
    });
})

$("#info").fileinput({
    language: 'zh-TW',
    theme: "fas",
    showUpload: false,
    uploadUrl: config.server + '/v1/course',
    ajaxSettings: {
        headers: {
            'Authorization': 'Bearer ' + $.cookie('token'),
        }
    },
    uploadExtraData: function (previewId, index) {
        return {
            'Topic': $('#topic').val(),
            'Type': $('#type').val(),
            'Start': $('#start input').val(),
            'End': $('#end input').val(),
        }
    }
}).on('fileuploaderror', function (event, data, msg) {
    alert(msg)
    $('div.kv-upload-progress.kv-hidden').css({ 'display': 'none' })
})

$('#course-form').on('submit', function (e) {
    e.preventDefault();
    $("#info").fileinput('upload')
})
