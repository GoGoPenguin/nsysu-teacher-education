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

const serviceLearningTable = $('table#service-learning').DataTable({
    processing: true,
    serverSide: true,
    ordering: false,
    searching: false,
    ajax: {
        url: config.server + '/v1/service-learning',
        type: 'GET',
        dataSrc: (d) => {
            d.list.forEach((element, index, array) => {
                array[index].Type = TYPE[element.Type];
                array[index].Date = element.Start.substring(0, 10) + ' ~ ' + element.End.substring(0, 10)
            })
            return d.list
        },
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', 'Bearer ' + token);
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
        url: config.server + '/v1/service-learning/sign-up',
        type: 'GET',
        dataSrc: (d) => {
            d.list.forEach((element, index, array) => {
                array[index].ServiceLearning.Type = TYPE[element.ServiceLearning.Type];
                array[index].Status = STATUS[array[index].Status]
                array[index].Date = element.ServiceLearning.Start.substring(0, 10) + ' ~ ' + element.ServiceLearning.End.substring(0, 10)
                array[index].Button = '<button class="btn btn-primary" onclick="check(' + index + ')">審核</button>'
            })
            return d.list
        },
        beforeSend: (xhr) => {
            let token = $.cookie('token')
            if (token == undefined) {
                renewToken()
                token = $.cookie('token')
            }

            xhr.setRequestHeader('Authorization', 'Bearer ' + token);
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
})

$('#service-learning-form').on('submit', (e) => {
    e.preventDefault();

    $.ajax({
        url: config.server + '/v1/service-learning',
        type: 'POST',
        data: {
            'Type': $('#type').val(),
            'Content': $('#content').val(),
            'Session': $('#start-time input').val() + ' ~ ' + $('#end-time input').val(),
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

            xhr.setRequestHeader('Authorization', 'Bearer ' + token);
        },
        error: function (xhr) {
            alert('Unexcepted Error')
            console.error(xhr);
        },
        success: function (response) {
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