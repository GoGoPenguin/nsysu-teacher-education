$(document).ready(function () {
    $.ajax({
        url: config.server + '/v1/course',
        type: 'GET',
        error: function (xhr) {
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
            if (response.list.length == 0) {
                $('#course tbody').append('\
                        <tr>\
                            <td scope="row" colspan="5" style="text-align: center">尚無資料</td>\
                        </tr>\
                    ')
            } else {
                response.list.forEach(function (element, index) {
                    let startDate = element.Start.substring(0, 10)
                    let startTime = element.Start.substring(11, 19)
                    let endDate = element.End.substring(0, 10)
                    let endTime = element.End.substring(11, 19)
                    let time = ""

                    if (startDate == endDate) {
                        time = startDate + ' ' + startTime + ' ~ ' + endTime
                    } else {
                        time = startDate + ' ' + startTime + ' ~ ' + endDate + ' ' + endTime
                    }

                    $('#course tbody').append('\
                        <tr>\
                            <th scope="row">'+ index + '</th>\
                            <td>'+ element.Topic + '</td>\
                            <td>'+ time + '</td>\
                            <td class="info">'+ element.Information + '</td>\
                            <td>'+ element.Type + '</td>\
                            <td><button class="btn btn-primary" onclick="signUp(' + element.ID + ')">報名</button></td>\
                        </tr>\
                    ')
                })
            }
        }
    });
})

$('#course tbody').on('click', 'td.info', function () {
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

function signUp(id) {
    $('input.course-id').val(id)
    $('#signUpModal').modal('show')
}

$('#signUpModal form').on('submit', function (e) {
    e.preventDefault();

    $.ajax({
        url: config.server + '/v1/course/sign-up',
        type: 'POST',
        data: {
            'Account': $.cookie('account'),
            'CourseID': $('input.course-id').val(),
            'Meal': $('#meal').val(),
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
            $('#signUpModal').modal('hide')
            $('section#course div.alert').show('fast')
        }
    });
})

$('section#course span').click(function () {
    $('section#course div.alert').hide('slow')
})