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
let StudentServiceLearningID = undefined

const getStudentServiceLearning = () => {
    $.ajax({
        url: `${config.server}/v1/service-learning/student`,
        type: 'GET',
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.list.length == 0) {
                $('#student-service-learning tbody').append(`
                        <tr>
                            <td scope="row" colspan="8" style="text-align: center">尚無資料</td>
                        </tr>
                    `)
            } else {
                response.list.forEach((element, index) => {
                    let color = 'class="text-dark"'
                    if (element.Status === 'pass') {
                        color = 'class="text-success"'
                    } else if (element.Status === 'failed') {
                        color = 'class="text-danger"'
                    }

                    let date = `${element.ServiceLearning.Start.substring(0, 10)} ~ ${element.ServiceLearning.End.substring(0, 10)}`
                    let result = `
                        <tr>
                            <th scope="row">${index}</th>
                            <td ${color}>${STATUS[element.Status]}</td>
                            <td>${TYPE[element.ServiceLearning.Type]}</td>
                            <td>${element.ServiceLearning.Content}</td>
                            <td>${date}</td>
                            <td>${element.ServiceLearning.Session}</td>
                            <td>${element.ServiceLearning.Hours}</td>
                    `

                    if (element.Status !== 'pass') {
                        result = `${result}<td><button class="btn btn-primary" onclick="edit(${element.ID})">編輯</button></td></tr>`
                    } else {
                        result = `${result}<td></td></tr>`
                    }

                    $('#student-service-learning tbody').append(result)
                })
            }
        }
    });
}

$(document).ready(() => {
    getStudentServiceLearning()

    $("#reference").fileinput({
        language: 'zh-TW',
        theme: "fas",
        showPreview: false,
        uploadUrl: `${config.server}/v1/service-learning/student`,
        ajaxSettings: {
            headers: {
                'Authorization': `Bearer ${$.cookie('token')}`,
            },
            method: "PATCH"
        },
        uploadExtraData: (previewId, index) => {
            return {
                'StudentServiceLearningID': StudentServiceLearningID,
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
        setTimeout(() => {
            $('#reference').fileinput('clear')
        }, 1000)
    }).on('fileuploaderror', (event, data, msg) => {
        swal({
            title: '',
            text: '失敗',
            icon: "error",
            timer: 1000,
            buttons: false,
        })
    })

    $("#review").fileinput({
        language: 'zh-TW',
        theme: "fas",
        showPreview: false,
        uploadUrl: `${config.server}/v1/service-learning/student`,
        ajaxSettings: {
            headers: {
                'Authorization': `Bearer ${$.cookie('token')}`,
            },
            method: "PATCH"
        },
        uploadExtraData: (previewId, index) => {
            return {
                'StudentServiceLearningID': StudentServiceLearningID,
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
        setTimeout(() => {
            $('#review').fileinput('clear')
        }, 1000)
    }).on('fileuploaderror', (event, data, msg) => {
        swal({
            title: '',
            text: '失敗',
            icon: "error",
            timer: 1000,
            buttons: false,
        })
    })
})

const edit = id => {
    StudentServiceLearningID = id
    $('#Modal').modal('show')
}