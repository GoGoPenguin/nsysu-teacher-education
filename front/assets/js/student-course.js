const STATUS = {
    '': '未審核',
    'pass': '通過',
    'failed': '未通過',
}

$(document).ready(function () {
    $.ajax({
        url: config.server + '/v1/course/sign-up',
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
                            <td scope="row" colspan="8" style="text-align: center">尚無資料</td>\
                        </tr>\
                    ')
            } else {
                response.list.forEach(function (element, index) {
                    let startDate = element.Course.Start.substring(0, 10)
                    let startTime = element.Course.Start.substring(11, 19)
                    let endDate = element.Course.End.substring(0, 10)
                    let endTime = element.Course.End.substring(11, 19)
                    let time = ""

                    if (startDate == endDate) {
                        time = startDate + ' ' + startTime + ' ~ ' + endTime
                    } else {
                        time = startDate + ' ' + startTime + ' ~ ' + endDate + ' ' + endTime
                    }

                    $('#student-course tbody').append('\
                        <tr>\
                            <th scope="row">'+ index + '</th>\
                            <td>'+ element.Course.Topic + '</td>\
                            <td>'+ time + '</td>\
                            <td>'+ element.Course.Type + '</td>\
                            <td>'+ STATUS[element.Status] + '</td>\
                            <td>'+ element.Comment + '</td>\
                            <td><button class="btn btn-primary">心得上傳</button></td>\
                        </tr>\
                    ')
                })
            }
        }
    });
})