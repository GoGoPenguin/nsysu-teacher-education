const STATUS = {
    '': '未審核',
    'pass': '通過',
    'failed': '未通過',
}

let studentCourses = []
let studentCoursesIndex = -1

const courseTable = $('table#course').DataTable({
    processing: true,
    serverSide: true,
    ordering: false,
    searching: false,
    ajax: {
        url: `${config.server}/v1/course`,
        type: 'GET',
        dataSrc: (d) => {
            d.list.forEach((element, index, array) => {
                let startDate = array[index].Start.substring(0, 10)
                let startTime = array[index].Start.substring(11, 19)
                let endDate = array[index].End.substring(0, 10)
                let endTime = array[index].End.substring(11, 19)

                if (startDate == endDate) {
                    array[index].Time = `${startDate} ${startTime} ~ ${endTime}`
                } else {
                    array[index].Time = `${startDate} ${startTime} ~ ${endDate} ${endTime}`
                }
            })
            return d.list
        },
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        error: (xhr, error, thrown) => {
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
                    timer: 1000,
                    buttons: false,
                })
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

const studentCourseTable = $('table#student-course').DataTable({
    processing: true,
    serverSide: true,
    ordering: false,
    searching: false,
    ajax: {
        url: `${config.server}/v1/course/sign-up`,
        type: 'GET',
        dataSrc: (d) => {
            d.list.forEach((element, index, array) => {
                let startDate = array[index].Course.Start.substring(0, 10)
                let startTime = array[index].Course.Start.substring(11, 19)
                let endDate = array[index].Course.End.substring(0, 10)
                let endTime = array[index].Course.End.substring(11, 19)

                if (startDate == endDate) {
                    array[index].Time = `${startDate} ${startTime} ~ ${endTime}`
                } else {
                    array[index].Time = `${startDate} ${startTime} ~ ${endDate} ${endTime}`
                }

                if (element.Status !== 'pass') {
                    array[index].Button = `<button class="btn btn-primary" onclick="check(${index})">審核</button>`
                } else {
                    array[index].Button = ''
                }

                array[index].Status = STATUS[array[index].Status]

                studentCourses.push(element)
            })

            return d.list
        },
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        error: (xhr, error, thrown) => {
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
                    timer: 1000,
                    buttons: false,
                })
            }
        }
    },
    columns: [
        { data: "Status" },
        { data: "Student.Account" },
        { data: "Student.Number" },
        { data: "Student.Major" },
        { data: "Student.Name" },
        { data: "Meal" },
        { data: "Course.Topic" },
        { data: "Course.Type" },
        { data: "Time" },
        { data: "Button" },
    ],
    language: {
        url: '/assets/languages/chinese.json'
    },
});

$(document).ready(() => {
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
        url: `${config.server}/v1/course/${filename}`,
        type: 'GET',
        xhrFields: {
            responseType: "blob"
        },
        error: (xhr) => {
            swal({
                title: '',
                text: '修改失敗',
                icon: "error",
                timer: 1000,
                buttons: false,
            })
            console.error(xhr);
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
    uploadUrl: `${config.server}/v1/course`,
    ajaxSettings: {
        headers: {
            'Authorization': 'Bearer ' + $.cookie('token'),
        }
    },
    uploadExtraData: (previewId, index) => {
        return {
            'Topic': $('#topic').val(),
            'Type': $('#type').val(),
            'Start': $('#start input').val(),
            'End': $('#end input').val(),
        }
    }
}).on('fileuploaded', (event, previewId, index, fileId) => {
    swal({
        title: '',
        text: '成功',
        icon: "success",
        timer: 1000,
        buttons: false,
    })
    courseTable.ajax.reload();
    $('#course-form')[0].reset()
}).on('fileuploaderror', (event, data, msg) => {
    swal({
        title: '',
        text: '新增失敗',
        icon: "error",
        timer: 1000,
        buttons: false,
    })
    $('div.kv-upload-progress.kv-hidden').css({ 'display': 'none' })
})

$('#course-form').on('submit', (e) => {
    e.preventDefault();
    $("#info").fileinput('upload')
})

const check = (index) => {
    studentCoursesIndex = index

    $('#checkModal .status p').html(studentCourses[index].Status)
    $('#checkModal .name input').val(studentCourses[index].Student.Name)
    $('#checkModal .major input').val(studentCourses[index].Student.Major)
    $('#checkModal .account input').val(studentCourses[index].Student.Account)
    $('#checkModal .number input').val(studentCourses[index].Student.Number)
    $('#checkModal .course-topic input').val(studentCourses[index].Course.Topic)
    $('#checkModal .course-type input').val(studentCourses[index].Course.Type)
    $('#checkModal .course-review').val(studentCourses[index].Review)
    $('#checkModal').modal('show')
}

$('#checkModal .btn-primary').click(() => {
    $.ajax({
        url: `${config.server}/v1/course/status`,
        type: 'PATCH',
        data: {
            StudentCourseID: studentCourses[studentCoursesIndex].ID,
            Status: 'pass',
        },
        error: (xhr) => {
            console.error(xhr);
            $('#checkModal').modal('hide')
            swal({
                title: '',
                text: '修改失敗',
                icon: "error",
                timer: 1500,
                buttons: false,
            })
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
            $('#checkModal').modal('hide')
            swal({
                title: '',
                text: '修改成功',
                icon: "success",
                timer: 1500,
                buttons: false,
            })

            studentCourseTable.ajax.reload()
        }
    });
})

$('#checkModal .btn-danger').click(() => {
    $.ajax({
        url: `${config.server}/v1/course/status`,
        type: 'PATCH',
        data: {
            StudentCourseID: studentCourses[studentCoursesIndex].ID,
            Status: 'failed',
        },
        error: (xhr) => {
            console.error(xhr);
            $('#checkModal').modal('hide')
            swal({
                title: '',
                text: '修改失敗',
                icon: "error",
                timer: 1500,
                buttons: false,
            })
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
            $('#checkModal').modal('hide')
            swal({
                title: '',
                text: '修改成功',
                icon: "success",
                timer: 1500,
                buttons: false,
            })

            let row = $('table#student-course tbody').children('tr').eq(0);
            let col = row.children('td').eq(0)
            col.html(STATUS['failed'])
        }
    });
})