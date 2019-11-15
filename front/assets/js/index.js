const TYPE = {
    'both': '同時認列教育實習服務暨志工服務',
    'internship': '實習服務',
    'volunteer': '志工服務',
}

let studentCourses = []
let studentServiceLearnings = []

let editedID = null
let editedItem = null

$(document).ready(() => {
    loading()

    Promise.all([
        getStudentCourses(),
        getStudentServiceLearning(),
        getLetures()
    ]).then(() => {
        unloading()
    }).catch(() => {
        setTimeout(() => {
            removeCookie()
        }, 1500)
    })
})

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
            if (response.list.length > 0) {
                response.list.forEach((element, index) => {
                    studentCourses.push(Object.assign({}, element))
                })
            }
        },
        complete: () => {
            getCourses()
        }
    });
}

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
            if (response.list.length > 0) {
                response.list.forEach((element, index) => {
                    studentServiceLearnings.push(Object.assign({}, element))
                })
            }
        },
        complete: () => {
            getServiceLearning()
        }
    });
}

const getCourses = () => {
    $.ajax({
        url: `${config.server}/v1/course`,
        type: 'GET',
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
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
                    let startDate = dayjs(element.Start).format('YYYY-MM-DD')
                    let startTime = dayjs(element.Start).format('HH:mm:ss')
                    let endDate = dayjs(element.End).format('YYYY-MM-DD')
                    let endTime = dayjs(element.End).format('HH:mm:ss')
                    let time = ""
                    let action = ''

                    if (startDate == endDate) {
                        time = `${startDate} ${startTime} ~ ${endTime}`
                    } else {
                        time = `${startDate} ${startTime} ~ ${endDate}  ${endTime}`
                    }

                    let studentCourse = studentCourses.find(studentCourse => {
                        return element.ID === studentCourse.Course.ID
                    })

                    if (studentCourse !== undefined) {
                        action = `<button class="btn btn-primary" disabled>已報名</button>`
                    } else {
                        action = `<button class="btn btn-primary" onclick="signUpCourse(${element.ID}, this)">報名</button>`
                    }

                    $('#course tbody').append(`
                        <tr>
                            <th scope="row">${index}</th>
                            <td>${element.Topic}</td>
                            <td>${time}</td>
                            <td class="info" onclick="getInformation(${element.ID}, '${element.Information}')">${element.Information}</td>
                            <td>${element.Type}</td>
                            <td>${action}</td>
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
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
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
                    let startDate = dayjs(element.Start).format('YYYY-MM-DD')
                    let endDate = dayjs(element.End).format('YYYY-MM-DD')
                    let action = ''

                    let studentServiceLearning = studentServiceLearnings.find(studentServiceLearning => {
                        return element.ID === studentServiceLearning.ServiceLearning.ID
                    })

                    if (studentServiceLearning !== undefined) {
                        action = `<button class="btn btn-primary" disabled>已報名</button>`
                    } else {
                        action = `<button class="btn btn-primary" onclick="signUpServiceLearning(${element.ID}, this)">報名</button>`
                    }

                    $('#service-learning tbody').append(`
                        <tr>
                            <th scope="row">${index}</th>
                            <td>${TYPE[element.Type]}</td>
                            <td>${element.Content}</td>
                            <td>${startDate} ~ ${endDate}</td>
                            <td>${element.Session}</td>
                            <td>${element.Hours}</td>
                            <td>${action}</td>
                        </tr>
                    `)
                })
            }
        }
    });
}

const getLetures = () => {
    return $.ajax({
        url: `${config.server}/v1/leture`,
        type: 'GET',
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.list.length == 0) {
                $('#leture tbody').append(`
                    <tr>
                        <td scope="row" colspan="7" style="text-align: center">尚無資料</td>
                    </tr>
                `)
            } else {
                response.list.forEach((element, index) => {
                    $('#leture tbody').append(`
                    <tr>
                        <th scope="row">${index}</th>
                        <td>${element.Name}</td>
                        <td>${element.MinCredit}</td>
                        <td>${element.Comment}</td>
                        <td><button class="btn btn-secondary mr-3" onclick="detail(${element.ID}, this)">查看</button><button class="btn btn-primary">報名</button></td>
                    </tr>
                `)
                })
            }
        }
    });
}

const getInformation = (id, filename) => {
    $.ajax({
        url: `${config.server}/v1/course/${id}`,
        type: 'GET',
        xhrFields: {
            responseType: "blob"
        },
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
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

const signUpCourse = (id, el) => {
    editedID = id
    editedItem = el
    $('#courseModal').modal('show')
}

$('#courseModal form').on('submit', (e) => {
    e.preventDefault();

    $.ajax({
        url: `${config.server}/v1/course/sign-up`,
        type: 'POST',
        data: {
            'CourseID': editedID,
            'Meal': $('#meal').val(),
        },
        beforeSend: (xhr) => {
            $('#courseModal div.modal-footer button.btn.btn-primary').html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp載入中...')
            $('#courseModal div.modal-footer button.btn.btn-primary').attr("disabled", true)
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.code === 0) {
                swal({
                    title: '',
                    text: '報名成功',
                    icon: "success",
                    timer: 1500,
                    buttons: false,
                })
                $(editedItem).html('已報名')
                $(editedItem).attr("disabled", true)
            } else {
                swal({
                    title: '',
                    text: '報名失敗',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
                $(editedItem).html('報名')
            }
        },
        complete: () => {
            $('#courseModal div.modal-footer button.btn.btn-primary').html('送出')
            $('#courseModal div.modal-footer button.btn.btn-primary').attr("disabled", false)
            $('#courseModal').modal('hide')

            editedID = null
            editedItem = null
        }
    });
})

const signUpServiceLearning = (id, el) => {
    $.ajax({
        url: `${config.server}/v1/service-learning/sign-up`,
        type: 'POST',
        data: {
            'ServiceLearningID': id,
        },
        beforeSend: (xhr) => {
            $(el).html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp載入中...')
            $(el).attr("disabled", true)
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.code === 0) {
                swal({
                    title: '',
                    text: '報名成功',
                    icon: "success",
                    timer: 1500,
                    buttons: false,
                })
                $(el).html('已報名')
            } else {
                swal({
                    title: '',
                    text: '報名失敗',
                    icon: "error",
                    timer: 1500,
                    buttons: false,
                })
                $(el).html('報名')
                $(el).attr("disabled", false)
            }
        }
    });
}

const detail = (id, el) => {
    $.ajax({
        url: `${config.server}/v1/leture/${id}`,
        type: 'GET',
        beforeSend: (xhr) => {
            $(el).html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp載入中...')
            $(el).attr("disabled", true)
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, '失敗')
        },
        success: (response) => {
            if (response.code === 0) {
                let leture = response.data
                let html = ''

                for (let category of leture.Categories) {
                    let content = ''
                    let comment = ''
                    let subjects = 0

                    if (category.MinCredit > 1) {
                        comment += `總共至少修習${category.MinCredit}學分<br>`
                    }

                    if (category.MinType > 1) {
                        comment += `總共至少修習${category.MinType}類別<br>`
                    }

                    for (let type of category.Types) {
                        let condition1 = category.Types.indexOf(type) === category.Types.length - 1
                        let subjectGroups = 0

                        if (type.MinCredit > 1) {
                            comment += `${type.Name}至少修習${type.MinCredit}學分<br>`
                        }

                        for (let group of type.Groups) {
                            let condition2 = type.Groups.indexOf(group) === type.Groups.length - 1

                            subjects += group.Subjects.length
                            subjectGroups += group.Subjects.length

                            for (let subject of group.Subjects) {
                                let condition3 = group.Subjects.indexOf(subject) === group.Subjects.length - 1
                                let temp = '<tr>'

                                if (condition1 && condition2 && condition3) {
                                    temp += `<td colspan="2" rowspan="${subjects}" class="align-middle">${category.Name}</td>`
                                }

                                if (condition2 && condition3) {
                                    temp += `<td colspan="2" rowspan="${subjectGroups}" class="align-middle">${type.Name}</td>`
                                }

                                temp += `<td colspan="1" class="align-middle">${subject.Compulsory ? "必修" : "選修"}</td>`

                                if (group.MinCredit > 0) {
                                    temp += `<td colspan="3" class="align-middle">${subject.Name}</td>`

                                    if (condition3) {
                                        temp += `<td colspan="1" rowspan="${group.Subjects.length}" class="align-middle vericaltext">至少${group.MinCredit}學分</td>`
                                    }
                                } else {
                                    temp += `<td colspan="4" class="align-middle">${subject.Name}</td>`
                                }

                                temp += `<td colspan="1" class="align-middle">${subject.Credit}</td>`

                                if (condition1 && condition2 && condition3) {
                                    temp += `<td colspan="2" rowspan="${subjects}" class="align-middle">${comment}</td>`
                                }

                                temp += `</tr>`
                                content = `${temp}${content}`
                            }
                        }
                    }

                    html += content
                }

                $('#detailModal .modal-title').html(leture.Name)
                $('#detailModal #name').html(leture.Name)
                $('#detailModal #min_credit').html(leture.MinCredit)
                $('#detailModal #comment').html(leture.Comment)
                $('#detailModal #categories').html(html)
                $('#detailModal').modal('show')
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
            $(el).html('查看')
            $(el).attr("disabled", false)
        }
    });
}