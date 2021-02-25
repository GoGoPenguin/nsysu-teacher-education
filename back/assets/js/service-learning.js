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
const STATUS_COLOR = {
    '': 'text-dark',
    'pass': 'text-success',
    'failed': 'text-danger',
}

let serviceLearningID = null
let serviceLearnings = []

let studentServiceLearningIndex = undefined
let studentServiceLearnings = []

$(document).ready(() => {
})

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

                let today = dayjs(new Date())
                let checked = element.Show || (element.Show == null && dayjs(element.Start).isAfter(today)) ? 'checked' : ''

                array[index].Type = TYPE[element.Type];
                array[index].Date = `${dayjs(element.Start).format('YYYY-MM-DD')} ~ ${dayjs(element.End).format('YYYY-MM-DD')}`
                array[index].Button = `
                    <button class="btn btn-primary mr-1" onclick="update(${index})">編輯</button>
                    <button class="btn btn-danger" onclick="$('#deleteModal').modal('show'); serviceLearningID=${element.ID}">刪除</button>
                `
                array[index].CheckBox = element.CreatedBy == "" ? `<input id="checkbox-${element.ID}" class="form-check-input" type="checkbox" style="margin: auto" ${checked} onclick="showOrNotShow(${element.ID})"></input>` : ""
                array[index].CreatedBy = element.CreatedBy == "" ? "管理者" : element.CreatedBy
            })
            return d.list
        },
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr, error, thrown) => {
            errorHandle(xhr, xhr.responseText)
        }
    },
    columns: [
        { data: "CheckBox" },
        { data: "Type" },
        { data: "Content" },
        { data: "Date" },
        { data: "Session" },
        { data: "Hours" },
        { data: "CreatedBy" },
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
    // searching: false,
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
                array[index].Status = `<span class="${STATUS_COLOR[array[index].Status]}">${STATUS[array[index].Status]}</span>`
                array[index].Date = `${dayjs(element.ServiceLearning.Start).format('YYYY-MM-DD')} ~ ${dayjs(element.ServiceLearning.End).format('YYYY-MM-DD')}`
                array[index].Hours = array[index].Hours == null ? array[index].ServiceLearning.Hours : array[index].Hours;

                studentServiceLearnings.push(element)
            })
            return d.list
        },
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr, error, thrown) => {
            errorHandle(xhr, xhr.responseText)
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
        { data: "Hours" },
        { data: "Button" },
    ],
    language: {
        url: '/assets/languages/chinese.json'
    },
});

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
            $('#service-learning-form button.btn.btn-primary').html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp載入中...')
            $('#service-learning-form button.btn.btn-primary').attr("disabled", true)
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, '失敗')
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
                serviceLearningTable.ajax.reload(null, false)
            } else {
                swal({
                    title: '',
                    text: '失敗',
                    icon: "error",
                    timer: 1000,
                    buttons: false,
                })
            }
            $('#service-learning-form')[0].reset()
        },
        complete: () => {
            $('#service-learning-form button.btn.btn-primary').html('送出')
            $('#service-learning-form button.btn.btn-primary').attr("disabled", false)
        }
    });
})

const check = (index, readonly) => {
    studentServiceLearningIndex = index

    if (readonly) {
        $('#checkModal .modal-footer').hide()

        $('#comment').attr({
            'readonly': true,
            'class': 'form-control-plaintext',
        })

        $('#checkModal .hours input').attr({
            'readonly': true,
            'class': 'form-control-plaintext',
        })
    } else {
        $('#checkModal .modal-footer').show()

        $('#comment').attr({
            'readonly': false,
            'class': 'form-control',
        })

        $('#checkModal .hours input').attr({
            'readonly': false,
            'class': 'form-control',
        })
    }

    $('#checkModal .hours input').attr({
        'max': studentServiceLearnings[index].ServiceLearning.Hours,
        'min': 0,
    })

    $('#checkModal .status p').html(studentServiceLearnings[index].Status)
    $('#checkModal .name input').val(studentServiceLearnings[index].Student.Name)
    $('#checkModal .major input').val(studentServiceLearnings[index].Student.Major)
    $('#checkModal .account input').val(studentServiceLearnings[index].Student.Account)
    $('#checkModal .number input').val(studentServiceLearnings[index].Student.Number)
    $('#checkModal .type input').val(studentServiceLearnings[index].ServiceLearning.Type)
    $('#checkModal .hours input').val(studentServiceLearnings[index].Hours)
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
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, '失敗')
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
    let hours = $('#checkModal .hours input').val()
    let max = studentServiceLearnings[studentServiceLearningIndex].ServiceLearning.Hours
    let min = 0

    if (hours > max || hours < min) {
        swal({
            title: '',
            text: `時數必須介於${min}到${max}之間`,
            icon: "error",
            timer: 1500,
            buttons: false,
        })
        return
    }

    $.ajax({
        url: `${config.server}/v1/service-learning/student/status`,
        type: 'PATCH',
        data: {
            'StudentServiceLearningID': ID,
            'Status': status,
            "Hours": hours,
            'Comment': $('#comment').val(),
        },
        beforeSend: (xhr) => {
            if (status == 'pass') {
                $('#checkModal div.modal-footer button.btn.btn-primary').html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp載入中...')
                $('#checkModal div.modal-footer button.btn.btn-primary').attr("disabled", true)
            } else if (status == 'failed') {
                $('#checkModal div.modal-footer button.btn.btn-danger').html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp載入中...')
                $('#checkModal div.modal-footer button.btn.btn-danger').attr("disabled", true)
            }
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, '失敗')
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
                studentServiceLearnings[studentServiceLearningIndex].Hours = hours;
                studentServiceLearningTable.ajax.reload(null, false)
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
            if (status == 'pass') {
                $('#checkModal div.modal-footer button.btn.btn-primary').html('通過')
                $('#checkModal div.modal-footer button.btn.btn-primary').attr("disabled", false)
            } else if (status == 'failed') {
                $('#checkModal div.modal-footer button.btn.btn-danger').html('未通過')
                $('#checkModal div.modal-footer button.btn.btn-danger').attr("disabled", false)
            }
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
    $('#update-start-date input').val(dayjs(serviceLearning.Start).format('YYYY-MM-DD'))
    $('#update-end-date input').val(dayjs(serviceLearning.End).format('YYYY-MM-DD'))
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
            $('#updateModal div.modal-footer button.btn.btn-primary').html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp載入中...')
            $('#updateModal div.modal-footer button.btn.btn-primary').attr("disabled", true)
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, '失敗')
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
                serviceLearningTable.ajax.reload(null, false)
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
            $('#updateModal div.modal-footer button.btn.btn-primary').html('送出')
            $('#updateModal div.modal-footer button.btn.btn-primary').attr("disabled", false)
            $('#updateModal').modal('hide')
        }
    });
})

const deleteServiceLearning = () => {
    $.ajax({
        url: `${config.server}/v1/service-learning/${serviceLearningID}`,
        type: 'DELETE',
        beforeSend: (xhr) => {
            $('#deleteModal div.modal-footer button.btn.btn-danger').html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp載入中...')
            $('#deleteModal div.modal-footer button.btn.btn-danger').attr("disabled", true)
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, '失敗')
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
                serviceLearningTable.ajax.reload(null, false)
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
            $('#deleteModal div.modal-footer button.btn.btn-danger').html('送出')
            $('#deleteModal div.modal-footer button.btn.btn-danger').attr("disabled", false)
            $('#deleteModal').modal('hide')
        }
    });
}

const showOrNotShow = (id) => {
    let checkbox = $(`#checkbox-${id}`)

    $.ajax({
        url: `${config.server}/v1/service-learning/show/${id}`,
        type: 'PATCH',
        data: {
            Show: checkbox.prop('checked'),
        },
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            swal({
                title: '',
                text: '失敗',
                icon: "error",
                timer: 1000,
                buttons: false,
            })
        },
        success: (response) => {
            if (response.code === 0) {
                serviceLearningTable.ajax.reload(null, false)
            } else {
                swal({
                    title: '',
                    text: '失敗',
                    icon: "error",
                    timer: 1000,
                    buttons: false,
                })
            }
        }
    });
}