const TYPE = {
    'both': '同時認列教育實習服務暨志工服務',
    'internship': '實習服務',
    'volunteer': '志工服務',
}
const STATUS = {
    '': '未審核',
    'pass': '通過',
    'failed': '未通過',
}

let serviceLearningID = null
let serviceLearnings = []

let studentServiceLearningIndex = undefined
let studentServiceLearnings = []

const serviceLearningTable = $('table#service-learning').DataTable({
    processing: true,
    serverSide: true,
    ordering: false,
    ajax: {
        url: `${config.server}/v1/service-learning`,
        type: 'GET',
        dataSrc: (d) => {
            serviceLearnings = []

            d.list.forEach((element, index, array) => {
                serviceLearnings.push(Object.assign({}, element))

                array[index].Type = TYPE[element.Type];
                array[index].Date = `${element.Start.substring(0, 10)} ~ ${element.End.substring(0, 10)}`
                array[index].Button = `
                    <button class="btn btn-primary mr-1" onclick="update(${index})">編輯</button>
                    <button class="btn btn-danger" onclick="$('#deleteModal').modal('show'); serviceLearningID=${element.ID}">刪除</button>
                `
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
        { data: "Type" },
        { data: "Content" },
        { data: "Date" },
        { data: "Session" },
        { data: "Hours" },
        { data: "Button" },
    ],
    language: {
        url: '/assets/languages/chinese.json'
    },
});

const studentServiceLearningTable = $('table#student-service-learning').DataTable({
    processing: true,
    serverSide: true,
    ordering: false,
    searching: false,
    ajax: {
        url: `${config.server}/v1/service-learning/student`,
        type: 'GET',
        dataSrc: (d) => {
            studentServiceLearnings = []

            d.list.forEach((element, index, array) => {
                if (array[index].Status !== 'pass') {
                    array[index].Button = `<button class="btn btn-primary" onclick="check(${index}, false)">審核</button>`
                } else {
                    array[index].Button = `<button class="btn btn-secondary" onclick="check(${index}, true)">查看</button>`
                }

                array[index].ServiceLearning.Type = TYPE[element.ServiceLearning.Type];
                array[index].Status = STATUS[array[index].Status]
                array[index].Date = `${element.ServiceLearning.Start.substring(0, 10)} ~ ${element.ServiceLearning.End.substring(0, 10)}`

                studentServiceLearnings.push(element)
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
        { data: "ServiceLearning.Type" },
        { data: "ServiceLearning.Content" },
        { data: "Date" },
        { data: "ServiceLearning.Session" },
        { data: "ServiceLearning.Hours" },
        { data: "Button" },
    ],
    language: {
        url: '/assets/languages/chinese.json'
    },
});

$(document).ready(() => {
    $('#start-date').datetimepicker({
        format: 'YYYY-MM-DD',
        locale: 'zh-tw',
        autoclose: true,
        icons: {
            time: "fas fa-clock",
            date: "fa fa-calendar",
            up: "fas fa-angle-up",
            down: "fas fa-angle-down",
        }
    })

    $('#end-date').datetimepicker({
        format: 'YYYY-MM-DD',
        locale: 'zh-tw',
        autoclose: true,
        icons: {
            time: "fas fa-clock",
            date: "fa fa-calendar",
            up: "fas fa-angle-up",
            down: "fas fa-angle-down",
        }
    })
    $('#start-time').datetimepicker({
        format: 'LT',
        locale: 'zh-tw',
        autoclose: true,
        icons: {
            time: "fas fa-clock",
            date: "fa fa-calendar",
            up: "fas fa-angle-up",
            down: "fas fa-angle-down",
        }
    })

    $('#end-time').datetimepicker({
        format: 'LT',
        locale: 'zh-tw',
        autoclose: true,
        icons: {
            time: "fas fa-clock",
            date: "fa fa-calendar",
            up: "fas fa-angle-up",
            down: "fas fa-angle-down",
        }
    })

    $('#update-start-date').datetimepicker({
        format: 'YYYY-MM-DD',
        locale: 'zh-tw',
        autoclose: true,
        icons: {
            time: "fas fa-clock",
            date: "fa fa-calendar",
            up: "fas fa-angle-up",
            down: "fas fa-angle-down",
        }
    })

    $('#update-end-date').datetimepicker({
        format: 'YYYY-MM-DD',
        locale: 'zh-tw',
        autoclose: true,
        icons: {
            time: "fas fa-clock",
            date: "fa fa-calendar",
            up: "fas fa-angle-up",
            down: "fas fa-angle-down",
        }
    })
    $('#update-start-time').datetimepicker({
        format: 'LT',
        locale: 'zh-tw',
        autoclose: true,
        icons: {
            time: "fas fa-clock",
            date: "fa fa-calendar",
            up: "fas fa-angle-up",
            down: "fas fa-angle-down",
        }
    })

    $('#update-end-time').datetimepicker({
        format: 'LT',
        locale: 'zh-tw',
        autoclose: true,
        icons: {
            time: "fas fa-clock",
            date: "fa fa-calendar",
            up: "fas fa-angle-up",
            down: "fas fa-angle-down",
        }
    })
})

$('#service-learning-form').on('submit', (e) => {
    e.preventDefault();

    $.ajax({
        url: `${config.server}/v1/service-learning`,
        type: 'POST',
        data: {
            'Type': $('#type').val(),
            'Content': $('#content').val(),
            'Session': `${$('#start-time input').val()} ~ ${$('#end-time input').val()}`,
            'Hours': $('#hours').val(),
            'Start': $('#start-date input').val(),
            'End': $('#end-date input').val(),
        },
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
            swal({
                title: '',
                text: '成功',
                icon: "success",
                timer: 1000,
                buttons: false,
            })
            serviceLearningTable.ajax.reload()
            $('#service-learning-form')[0].reset()
        }
    });
})

const check = (index, readonly) => {
    studentServiceLearningIndex = index

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

    $('#checkModal .status p').html(studentServiceLearnings[index].Status)
    $('#checkModal .name input').val(studentServiceLearnings[index].Student.Name)
    $('#checkModal .major input').val(studentServiceLearnings[index].Student.Major)
    $('#checkModal .account input').val(studentServiceLearnings[index].Student.Account)
    $('#checkModal .number input').val(studentServiceLearnings[index].Student.Number)
    $('#checkModal .type input').val(studentServiceLearnings[index].ServiceLearning.Type)
    $('#checkModal .content').val(studentServiceLearnings[index].ServiceLearning.Content)
    $('#checkModal .reference input').val(studentServiceLearnings[index].Reference)
    $('#checkModal .review input').val(studentServiceLearnings[index].Review)
    $('#comment').val(studentServiceLearnings[index].Comment)
    $('#checkModal').modal('show')
}

const getFile = (file) => {
    let ID = studentServiceLearnings[studentServiceLearningIndex].ID
    let filename = $(`#checkModal .${file} input`).val()

    if (filename === "") {
        return
    }

    $.ajax({
        url: `${config.server}/v1/service-learning/student/${file}`,
        type: 'GET',
        xhrFields: {
            responseType: "blob"
        },
        data: {
            'StudentServiceLearningID': ID,
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

const updateStatus = (status) => {
    let ID = studentServiceLearnings[studentServiceLearningIndex].ID

    $.ajax({
        url: `${config.server}/v1/service-learning/student/status`,
        type: 'PATCH',
        data: {
            'StudentServiceLearningID': ID,
            'Status': status,
            'Comment': $('#comment').val(),
        },
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
                studentServiceLearningTable.ajax.reload()
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
            $('#checkModal').modal('hide')
        }
    });
}

const update = (index) => {
    let serviceLearning = serviceLearnings[index]
    let startTime, endTime
    [startTime, endTime] = serviceLearning.Session.split(" ~ ")
    serviceLearningID = serviceLearning.ID

    $('#update-type').val(serviceLearning.Type)
    $('#update-content').val(serviceLearning.Content)
    $('#update-start-date input').val(serviceLearning.Start.substring(0, 10))
    $('#update-end-date input').val(serviceLearning.End.substring(0, 10))
    $('#update-start-time input').val(startTime)
    $('#update-end-time input').val(endTime)
    $('#update-hours').val(serviceLearning.Hours)

    $('#updateModal').modal('show')
}

const editServiceLearning = () => {
    $('#update-submit').click()
}

$('#update-form').on('submit', (e) => {
    e.preventDefault()

    $.ajax({
        url: `${config.server}/v1/service-learning`,
        type: 'PATCH',
        data: {
            'ServiceLearningID': serviceLearningID,
            'Type': $('#update-type').val(),
            'Content': $('#update-content').val(),
            'Session': `${$('#update-start-time input').val()} ~ ${$('#update-end-time input').val()}`,
            'Hours': $('#update-hours').val(),
            'Start': $('#update-start-date input').val(),
            'End': $('#update-end-date input').val(),
        },
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
                serviceLearningTable.ajax.reload()
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
            $('#updateModal').modal('hide')
        }
    });
})

const deleteServiceLearning = () => {
    $.ajax({
        url: `${config.server}/v1/service-learning/${serviceLearningID}`,
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
                serviceLearningTable.ajax.reload()
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