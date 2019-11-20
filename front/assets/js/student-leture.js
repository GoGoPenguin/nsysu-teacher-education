let editedItem = null

$(document).ready(() => {
    loading()

    Promise.all([
        getStudentLeture()
    ]).then(() => {
        unloading()
    }).catch(() => {
        setTimeout(() => {
            removeCookie()
        }, 1500)
    })
})

const getStudentLeture = () => {
    $.ajax({
        url: `${config.server}/v1/leture/student`,
        type: 'GET',
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.list.length == 0) {
                $('#student-leture tbody').append(`
                    <tr>
                        <td scope="row" colspan="8" style="text-align: center">尚無資料</td>
                    </tr>
                `)
            } else {
                response.list.forEach((element, index) => {
                    $('#student-leture tbody').append(`
                        <tr>
                            <th scope="row">${index}</th>
                            <td class="${element.Pass ? 'text-success' : 'text-danger'}">${element.Pass ? '通過' : '未通過'}</td>
                            <td>${element.Leture.Name}</td>
                            <td>${element.Leture.MinCredit}</td>
                            <td>${element.Leture.Comment}</td>
                            <td><button class="btn btn-primary" onclick="getStudentLetureDetail(${element.ID})">編輯</button></td>
                        </tr>
                    `)
                })
            }
        }
    });
}

const getStudentLetureDetail = (id) => {
    $.ajax({
        url: `${config.server}/v1/leture/student/detail/${id}`,
        type: 'GET',
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.code === 0) {
                editedItem = response.data.Leture
                let leture = response.data.Leture
                let html = ''

                for (let category of leture.Categories) {
                    let content = ''
                    let comment = ''
                    let subjects = 0

                    if (category.MinCredit > 1) {
                        comment += `總共至少修習${category.MinCredit}學分<br>`
                    }

                    if (category.MinType > 1) {
                        comment += `總共至少修習${category.MinType}專長<br>`
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
                                    temp += `<td colspan="2" rowspan="${subjects}" class="align-middle" id="category${category.ID}">${category.Name}</td>`
                                }

                                if (condition2 && condition3) {
                                    temp += `<td colspan="2" rowspan="${subjectGroups}" class="align-middle" id="type${type.ID}">${type.Name}</td>`
                                }

                                temp += `<td colspan="1" class="align-middle" id="subject${subject.ID}">${subject.Compulsory ? "必修" : "選修"}</td>`

                                if (group.MinCredit > 0) {
                                    temp += `<td colspan="3" class="align-middle">${subject.Name}</td>`

                                    if (condition3) {
                                        temp += `<td colspan="1" rowspan="${group.Subjects.length}" class="align-middle vericaltext" id="group${group.ID}">至少${group.MinCredit}學分</td>`
                                    }
                                } else {
                                    temp += `<td colspan="4" class="align-middle">${subject.Name}</td>`
                                }

                                temp += `<td colspan="1" class="align-middle">${subject.Credit}</td>`
                                temp += `
                                    <td colspan="1" class="align-middle">
                                        <input type="number" class="form-control form-control-sm" max="100" min="0" value="${subject.Score}" onblur="updateSubject(${id}, ${subject.ID})" id="score${subject.ID}">
                                    </td>`
                                temp += `
                                    <td colspan="1" class="align-middle">
                                        <div class="custom-control custom-checkbox">
                                            <input type="checkbox" class="custom-control-input" id="check${subject.ID}" onclick="updateSubject(${id}, ${subject.ID})" ${subject.Pass == 1 ? 'checked' : ''} ${subject.Score == null ? 'disabled' : ''}>
                                            <label class="custom-control-label" for="check${subject.ID}"></label>
                                        </div>
                                    </td>`

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
        }
    });
}

const updateSubject = (studentLetureID, subjectID) => {
    if ($(`#score${subjectID}`).val() > 100 || $(`#score${subjectID}`).val() < 0) {
        swal({
            title: '',
            text: '成績應介於0到100之間',
            icon: "error",
            timer: 1500,
            buttons: false,
        })
        $(`#score${subjectID}`).val(null)
        return
    }
    if ($(`#score${subjectID}`).val() == '') {
        $(`#check${subjectID}`).prop("disabled", true)
        $(`#check${subjectID}`).prop("checked", false)
    } else {
        $(`#check${subjectID}`).prop("disabled", false)
    }

    $.ajax({
        url: `${config.server}/v1/leture/student/subject`,
        type: 'PATCH',
        data: {
            'StudentLetureID': studentLetureID,
            'SubjectID': subjectID,
            'Score': $(`#score${subjectID}`).val(),
            'Pass': $(`#check${subjectID}`).prop("checked"),
        },
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.code === 0) {
                reloadLeture(studentLetureID)
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
    })
}

const reloadLeture = (id) => {
    $.ajax({
        url: `${config.server}/v1/leture/student/detail/${id}`,
        type: 'GET',
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.code === 0) {
                editedItem = response.data.Leture
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
    })
}

const check = () => {
    let leture = editedItem
    let pass = true

    for (let category of leture.Categories) {
        let categoryCredit = 0
        let categoryTypes = 0

        for (let type of category.Types) {
            let typeCredit = 0

            for (let group of type.Groups) {
                let groupCredit = 0

                for (let subject of group.Subjects) {
                    if (subject.Pass) {
                        groupCredit += subject.Credit
                    }

                    if (subject.Compulsory && !subject.Pass) {
                        if (!$(`#subject${subject.ID}`).hasClass('text-danger')) {
                            $(`#subject${subject.ID}`).addClass('text-danger')
                        }
                        pass = false
                    } else {
                        $(`#subject${subject.ID}`).removeClass('text-danger')
                    }
                }

                typeCredit += groupCredit

                if (groupCredit < group.MinCredit) {
                    if (!$(`#group${group.ID}`).hasClass('text-danger')) {
                        $(`#group${group.ID}`).addClass('text-danger')
                    }
                    pass = false
                } else {
                    $(`#group${group.ID}`).removeClass('text-danger')
                }
            }

            categoryCredit += typeCredit
            if (typeCredit != 0) {
                categoryTypes++
            }

            if (typeCredit < type.MinCredit) {
                if (!$(`#type${type.ID}`).hasClass('text-danger')) {
                    $(`#type${type.ID}`).addClass('text-danger')
                }
                pass = false
            } else {
                $(`#type${type.ID}`).removeClass('text-danger')
            }
        }

        if (categoryCredit < category.MinCredit || categoryTypes < category.MinType) {
            if (!$(`#category${category.ID}`).hasClass('text-danger')) {
                $(`#category${category.ID}`).addClass('text-danger')
            }
            pass = false
        } else {
            $(`#category${category.ID}`).removeClass('text-danger')
        }
    }

    if (pass) {
        swal({
            title: '',
            text: '通過',
            icon: "success",
            timer: 1500,
            buttons: false,
        })
    } else {
        swal({
            title: '',
            text: '未通過',
            icon: "error",
            timer: 1500,
            buttons: false,
        })
    }
}

$('#detailModal').on('hidden.bs.modal', () => {
    editedItem = null
})