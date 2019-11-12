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
        getStudentCourses(),
    ]).then(() => {
        unloading()
    }).catch(() => {
        setTimeout(() => {
            removeCookie()
        }, 1500)
    })
})

const getStudentCourses = () => {
    return $.ajax({
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

                    let startDate = element.Course.Start.substring(0, 10)
                    let startTime = element.Course.Start.substring(11, 19)
                    let endDate = element.Course.End.substring(0, 10)
                    let endTime = element.Course.End.substring(11, 19)
                    let time = ""

                    if (startDate == endDate) {
                        time = `${startDate} ${startTime} ~ ${endTime}`
                    } else {
                        time = `${startDate} ${startTime} ~ ${endDate} ${endTime}`
                    }

                    let color = 'class="text-dark"'
                    if (element.Status === 'pass') {
                        color = 'class="text-success"'
                    } else if (element.Status === 'failed') {
                        color = 'class="text-danger"'
                    }

                    let result = `
                        <tr>
                            <th scope="row">${index}</th>\
                            <td ${color}>${STATUS[element.Status]}</td>\
                            <td>${element.Course.Topic}</td>\
                            <td>${time}</td>\
                            <td>${element.Course.Type}</td>\
                            <td>${MEAL[element.Meal]}</td>\
                            <td>${element.Comment}</td>\
                            <td>${element.Review}</td>\
                    `

                    if (element.Status !== 'pass') {
                        result = `${result}<td><button class="btn btn-primary" onclick="edit(${index})">編輯</button></td></tr>`
                    } else {
                        result = `${result}<td></td></tr>`
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
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            $('#updateReviewModal').modal('hide')
            $(`button#${studentCourseID}`).parent().prev().html(review)

            swal({
                title: '',
                text: '成功',
                icon: "success",
                timer: 1500,
                buttons: false,
            })
        }
    })
})