const STATUS = {
    '': '未審核',
    'pass': '通過',
    'failed': '未通過',
}
const MEAL = {
    'vegetable': '素',
    'meat': '葷',
}

let courseID = null
let courses = []

let studentCourses = []
let studentCoursesIndex = null

const courseTable = $('table#course').DataTable({
    processing: true,
    serverSide: true,
    ordering: false,
    ajax: {
        url: `${config.server}/v1/course`,
        type: 'GET',
        dataSrc: (d) => {
            courses = []

            d.list.forEach((element, index, array) => {
                courses.push(Object.assign({}, element))

                let startDate = array[index].Start.substring(0, 10)
                let startTime = array[index].Start.substring(11, 19)
                let endDate = array[index].End.substring(0, 10)
                let endTime = array[index].End.substring(11, 19)

                if (startDate == endDate) {
                    array[index].Time = `${startDate} ${startTime} ~ ${endTime}`
                } else {
                    array[index].Time = `${startDate} ${startTime} ~ ${endDate} ${endTime}`
                }

                array[index].Button = `
                    <button class="btn btn-primary mr-1" onclick="update(${index})">編輯</button>
                    <button class="btn btn-danger" onclick="$('#deleteModal').modal('show'); courseID=${element.ID}">刪除</button>
                `

                array[index].Information = `<div onclick="getInformation(${element.ID}, '${element.Information}')">${element.Information}</div>`
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
        { data: "Button" },
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
        url: `${config.server}/v1/course/student`,
        type: 'GET',
        dataSrc: (d) => {
            studentCourses = []

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
                    array[index].Button = `<button class="btn btn-primary" onclick="check(${index}, false)">審核</button>`
                } else {
                    array[index].Button = `<button class="btn btn-secondary" onclick="check(${index}, true)">查看</button>`
                }

                array[index].Status = STATUS[array[index].Status]
                array[index].Meal = MEAL[array[index].Meal]

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

    $('#update-start').datetimepicker({
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

    $('#update-end').datetimepicker({
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

const getInformation = (id, filename) => {
    $.ajax({
        url: `${config.server}/v1/course/${id}`,
        type: 'GET',
        xhrFields: {
            responseType: "blob"
        },
        error: (xhr) => {
            swal({
                title: '',
                text: '失敗',
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
}

$("#info").fileinput({
    language: 'zh-TW',
    theme: "fas",
    showUpload: false,
    uploadUrl: `${config.server}/v1/course`,
    ajaxSettings: {
        headers: {
            'Authorization': `Bearer ${$.cookie('token')}`,
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
}).on('filepreupload', function (event, data, previewId, index) {
    let filename = data.files[0].name
    let length = (new TextEncoder().encode(filename)).length

    if (length > 36) {
        swal({
            title: '',
            text: '檔名太長',
            icon: "error",
            timer: 1500,
            buttons: false,
        })

        // this doesn't work at all
        // https://plugins.krajee.com/file-input/plugin-events#event-manipulation
        return {
            message: '檔名太長',
        }
    }
}).on('filecustomerror', function (event, params) {
    $("#info").fileinput('enable')
    $("#info").fileinput('reset')
})

$('#course-form').on('submit', (e) => {
    e.preventDefault();
    $("#info").fileinput('upload')
})

$("#update-info").fileinput({
    language: 'zh-TW',
    theme: "fas",
    showUpload: false,
    uploadUrl: `${config.server}/v1/course`,
    ajaxSettings: {
        headers: {
            'Authorization': `Bearer ${$.cookie('token')}`,
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
}).on('filepreupload', function (event, data, previewId, index) {
    let filename = data.files[0].name
    let length = (new TextEncoder().encode(filename)).length

    if (length > 36) {
        swal({
            title: '',
            text: '檔名太長',
            icon: "error",
            timer: 1500,
            buttons: false,
        })

        // this doesn't work at all
        // https://plugins.krajee.com/file-input/plugin-events#event-manipulation
        return {
            message: '檔名太長',
        }
    }
}).on('filecustomerror', function (event, params) {
    $("#info").fileinput('enable')
    $("#info").fileinput('reset')
})

$('#course-form').on('submit', (e) => {
    e.preventDefault();
    $("#info").fileinput('upload')
})

const check = (index, readonly) => {
    studentCoursesIndex = index

    if (readonly) {
        $('#checkModal .modal-footer').hide()
        $('#comment').attr('readonly', true)
        $('#comment').addClass('form-control-plaintext')
        $('#comment').removeClass('form-control')
    } else {
        $('#checkModal .modal-footer').show()
        $('#comment').attr('readonly', false)
        $('#comment').removeClass('form-control-plaintext')
        $('#comment').addClass('form-control')
    }

    $('#checkModal .status p').html(studentCourses[index].Status)
    $('#checkModal .name input').val(studentCourses[index].Student.Name)
    $('#checkModal .major input').val(studentCourses[index].Student.Major)
    $('#checkModal .account input').val(studentCourses[index].Student.Account)
    $('#checkModal .number input').val(studentCourses[index].Student.Number)
    $('#checkModal .course-topic input').val(studentCourses[index].Course.Topic)
    $('#checkModal .course-type input').val(studentCourses[index].Course.Type)
    $('#checkModal .course-review').val(studentCourses[index].Review)
    $('#comment').val(studentCourses[index].Comment)
    $('#checkModal').modal('show')
}

$('#checkModal .btn-primary').click(() => {
    $.ajax({
        url: `${config.server}/v1/course/student/status`,
        type: 'PATCH',
        data: {
            StudentCourseID: studentCourses[studentCoursesIndex].ID,
            Status: 'pass',
            Comment: $('#comment').val(),
        },
        error: (xhr) => {
            console.error(xhr);
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
            if (response.code === 0) {
                swal({
                    title: '',
                    text: '修改成功',
                    icon: "success",
                    timer: 1500,
                    buttons: false,
                })
                studentCourseTable.ajax.reload()
            } else {
                swal({
                    title: '',
                    text: '修改失敗',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
            }
        },
        complete: (data) => {
            $('#checkModal').modal('hide')
        }
    });
})

$('#checkModal .btn-danger').click(() => {
    $.ajax({
        url: `${config.server}/v1/course/student/status`,
        type: 'PATCH',
        data: {
            StudentCourseID: studentCourses[studentCoursesIndex].ID,
            Status: 'failed',
            Comment: $('#comment').val(),
        },
        error: (xhr) => {
            console.error(xhr);
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
            if (response.code === 0) {
                swal({
                    title: '',
                    text: '修改成功',
                    icon: "success",
                    timer: 1500,
                    buttons: false,
                })
                studentCourseTable.ajax.reload()
            } else {
                swal({
                    title: '',
                    text: '修改失敗',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
            }
        },
        complete: (data) => {
            $('#checkModal').modal('hide')
        }
    });
})

const update = (index) => {
    let course = courses[index]
    courseID = course.ID

    $('#update-topic').val(course.Topic)
    $('#update-type').val(course.Type)
    $('#update-start input').val(course.Start.substring(0, 19))
    $('#update-end input').val(course.End.substring(0, 19))

    $('#updateModal').modal('show')
}

const editCourse = () => {
    $('#update-submit').click()
}

$('#update-form').on('submit', (e) => {
    e.preventDefault()

    let filestack = $('#update-info').fileinput('getFileStack')
    let fstack = []
    $.each(filestack, (fileID, fileObj) => {
        if (fileObj !== undefined) {
            fstack.push(fileObj)
        }
    })

    let form = new FormData()
    form.append("CourseID", courseID)
    form.append("Topic", $('#update-topic').val())
    form.append("Type", $('#update-type').val())
    form.append("Start", $('#update-start input').val())
    form.append("End", $('#update-end input').val())

    if (fstack.length > 0) {
        form.append("Information", fstack[0].file)
    }

    $.ajax({
        url: `${config.server}/v1/course`,
        type: 'PATCH',
        data: form,
        contentType: false,
        processData: false,
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        error: (xhr) => {
            swal({
                title: '',
                text: '失敗',
                icon: "error",
                timer: 1000,
                buttons: false,
            })
            console.error(xhr);
        },
        success: (response) => {
            if (response.code === 0) {
                swal({
                    title: '',
                    text: '成功',
                    icon: "success",
                    timer: 1000,
                    buttons: false,
                })
                courseTable.ajax.reload()
            } else {
                swal({
                    title: '',
                    text: '失敗',
                    icon: "error",
                    timer: 1000,
                    buttons: false,
                })
            }
            $('#update-form')[0].reset()
        },
        complete: (data) => {
            $('#updateModal').modal('hide')
        }
    });
})

const deleteCourse = () => {
    $.ajax({
        url: `${config.server}/v1/course/${courseID}`,
        type: 'DELETE',
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', `Bearer ${token}`);
        },
        error: (xhr) => {
            swal({
                title: '',
                text: '失敗',
                icon: "error",
                timer: 1000,
                buttons: false,
            })
            console.error(xhr);
        },
        success: (response) => {
            if (response.code === 0) {
                swal({
                    title: '',
                    text: '成功',
                    icon: "success",
                    timer: 1000,
                    buttons: false,
                })
                courseTable.ajax.reload()
            } else {
                swal({
                    title: '',
                    text: '失敗',
                    icon: "error",
                    timer: 1000,
                    buttons: false,
                })
            }
        },
        complete: (data) => {
            $('#deleteModal').modal('hide')
        }
    });
}