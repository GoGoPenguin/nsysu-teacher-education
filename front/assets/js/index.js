const TYPE = {
    'both': '同時認列教育實習服務暨志工服務',
    'internship': '實習服務',
    'volunteer': '志工服務',
}

const getCourses = () => {
    $.ajax({
        url: `${config.server}/v1/course`,
        type: 'GET',
        error: (xhr) => {
            if (xhr.status === 401) {
                removeCookie()
            } else {
            }
        },
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        success: (response) => {
            if (response.list.length == 0) {
                $('#course tbody').append(`
                    <tr>
                        <td scope="row" colspan="6" style="text-align: center">尚無資料</td>
                    </tr>
                `)
            } else {
                response.list.forEach((element, index) => {
                    let startDate = element.Start.substring(0, 10)
                    let startTime = element.Start.substring(11, 19)
                    let endDate = element.End.substring(0, 10)
                    let endTime = element.End.substring(11, 19)
                    let time = ""

                    if (startDate == endDate) {
                        time = `${startDate} ${startTime} ~ ${endTime}`
                    } else {
                        time = `${startDate} ${startTime} ~ ${endDate}  ${endTime}`
                    }

                    $('#course tbody').append(`
                    <tr>
                        <th scope="row">${index}</th>
                        <td>${element.Topic}</td>
                        <td>${time}</td>
                        <td class="info">${element.Information}</td>
                        <td>${element.Type}</td>
                        <td><button class="btn btn-primary" onclick="signUpCourse(${element.ID})">報名</button></td>
                    </tr>
                `)
                })
            }
        }
    });
}

const getServiceLearning = () => {
    $.ajax({
        url: `${config.server}/v1/service-learning`,
        type: 'GET',
        error: (xhr) => {
            if (xhr.status === 401) {
                removeCookie()
            } else {

            }
        },
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        success: (response) => {
            if (response.list.length == 0) {
                $('#service-learning tbody').append(`
                    <tr>
                        <td scope="row" colspan="7" style="text-align: center">尚無資料</td>
                    </tr>
                `)
            } else {
                response.list.forEach((element, index) => {
                    let startDate = element.Start.substring(0, 10)
                    let endDate = element.End.substring(0, 10)

                    $('#service-learning tbody').append(`
                    <tr>
                        <th scope="row">${index}</th>
                        <td>${TYPE[element.Type]}</td>
                        <td>${element.Content}</td>
                        <td>${startDate} ~ ${endDate}</td>
                        <td>${element.Session}</td>
                        <td>${element.Hours}</td>
                        <td><button class="btn btn-primary" onclick="signUpServiceLearning(${element.ID})">報名</button></td>
                    </tr>
                `)
                })
            }
        }
    });
}


$(document).ready(() => {
    getCourses()
    getServiceLearning()
})

$('#course tbody').on('click', 'td.info', function () {
    let filename = $(this).text()

    $.ajax({
        url: `${config.server}/v1/course/${filename}`,
        type: 'GET',
        xhrFields: {
            responseType: "blob"
        },
        error: (xhr) => {
            if (xhr.status === 401) {
                removeCookie()
            } else {
                swal({
                    title: '',
                    text: '失敗',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
            }
        },
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        success: (response) => {
            if (response.code === 0) {
                let a = document.createElement('a');
                let url = window.URL.createObjectURL(response);
                a.href = url;
                a.download = filename;
                document.body.append(a);
                a.click();
                a.remove();
                window.URL.revokeObjectURL(url);
            } else {
                swal({
                    title: '',
                    text: '失敗',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
            }
        }
    });
})

const signUpCourse = (id) => {
    $('input.course-id').val(id)
    $('#courseModal').modal('show')
}

$('#courseModal form').on('submit', (e) => {
    e.preventDefault();

    $.ajax({
        url: `${config.server}/v1/course/sign-up`,
        type: 'POST',
        data: {
            'Account': $.cookie('account'),
            'CourseID': $('input.course-id').val(),
            'Meal': $('#meal').val(),
        },
        error: (xhr) => {
            if (xhr.status === 401) {
                removeCookie()
            } else {
                swal({
                    title: '',
                    text: '報名失敗',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
            }
        },
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        success: (response) => {
            $('#courseModal').modal('hide')

            if (response.code === 0) {
                swal({
                    title: '',
                    text: '報名成功',
                    icon: "success",
                    timer: 1500,
                    buttons: false,
                })
            } else {
                swal({
                    title: '',
                    text: '報名失敗',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
            }
        }
    });
})

const signUpServiceLearning = (id) => {
    $.ajax({
        url: `${config.server}/v1/service-learning/sign-up`,
        type: 'POST',
        data: {
            'Account': $.cookie('account'),
            'ServiceLearningID': id,
        },
        error: (xhr) => {
            if (xhr.status === 401) {
                removeCookie()
            } else {
                swal({
                    title: '',
                    text: '報名失敗',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
            }
        },
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        success: (response) => {
            $('#signUpModal').modal('hide')

            if (response.code === 0) {
                swal({
                    title: '',
                    text: '報名成功',
                    icon: "success",
                    timer: 1500,
                    buttons: false,
                })
            } else {
                swal({
                    title: '',
                    text: '報名失敗',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
            }
        }
    });
}