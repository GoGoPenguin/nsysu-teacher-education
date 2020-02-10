const STATUS = {
    '': '未審核',
    'pass': '通過',
    'failed': '未通過',
}
const MEAL = {
    'vegetable': '素',
    'meat': '葷',
}

let studentCourses = []

$(document).ready(() => {
    loading()

    Promise.all([
        getStudentInformation(),
        getStudentCourses(),
    ]).then(() => {
        unloading()
    }).catch(() => {
        setTimeout(() => {
            removeCookie()
        }, 1500)
    })
})

const getStudentInformation = () => {
    $.ajax({
        url: `${config.server}/v1/user`,
        type: 'GET',
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            student = Object.assign({}, response.data)
            $('div.greeting').html(`Hi, ${student.Name}同學`)
        }
    });
}

const getStudentCourses = () => {
    $.ajax({
        url: `${config.server}/v1/course/student`,
        type: 'GET',
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.list.length == 0) {
                $('#student-course tbody').append(`
                    <tr>
                        <td scope="row" colspan="8" style="text-align: center">尚無資料</td>
                    </tr>
                `)
            } else {
                studentCourses = []
                response.list.forEach((element, index) => {
                    studentCourses.push(Object.assign({}, element))

                    let startDate = dayjs(element.Course.Start).format('YYYY-MM-DD')
                    let startTime = dayjs(element.Course.Start).format('HH:mm')
                    let endDate = dayjs(element.Course.End).format('YYYY-MM-DD')
                    let endTime = dayjs(element.Course.End).format('HH:mm')
                    let time = ""

                    if (startDate == endDate) {
                        time = `${startDate} ${startTime} ~ ${endTime}`
                    } else {
                        time = `${startDate} ${startTime} ~ ${endDate} ${endTime}`
                    }

                    let color = 'class="waiting"'
                    if (element.Status === 'pass') {
                        color = 'class="success"'
                    } else if (element.Status === 'failed') {
                        color = 'class="danger"'
                    }

                    let result = `
                        <tr>
                            <td data-title="審核情況" ${color}><span>●</span>${STATUS[element.Status]}</td>\
                            <td data-title="類別">${element.Course.Type}</td>\
                            <td data-title="研習主題">${element.Course.Topic}</td>\
                            <td data-title="研習時段">${time}</td>\
                            <td data-title="審核結果說明">${element.Comment == "" ? "無" : element.Comment}</td>\
                    `

                    if (element.Status !== 'pass') {
                        result = `${result}<td><a class="btn_table" onclick="edit(${index})" id="${element.ID}">編輯</a></td></tr>`
                    } else {
                        result = `${result}<td><a class="btn_table disabled">編輯</a></td></tr>`
                    }

                    $('#student-course tbody').append(result)
                })
            }
        }
    });
}

const edit = (index) => {
    let id = studentCourses[index].ID
    let review = studentCourses[index].Review

    $('#updateReviewModal textarea').val(review)
    $('#updateReviewModal input').val(id)
    $('#updateReviewModal').modal('show')
}

$('#updateReviewModal form').on('submit', (e) => {
    e.preventDefault()

    let studentCourseID = $('#updateReviewModal input').val()
    let review = $('#updateReviewModal textarea').val()

    $.ajax({
        url: `${config.server}/v1/course/student/review`,
        type: 'PATCH',
        data: {
            'StudentCourseID': studentCourseID,
            'Review': review,
        },
        beforeSend: (xhr) => {
            $('#updateReviewModal .btn_table').addClass("disabled")
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.code === 0) {
                swal({
                    title: '',
                    text: '成功',
                    icon: "success",
                    timer: 1500,
                    buttons: false,
                })
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
        complete: () => {
            $('#updateReviewModal .btn_table').removeClass("disabled")
            $('#updateReviewModal').modal('hide')
        }
    })
})