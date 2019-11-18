let studentLetures = []

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
                studentLetures = []
                response.list.forEach((element, index) => {
                    studentLetures.push(Object.assign({}, element))

                    $('#student-leture tbody').append(`
                        <tr>
                            <th scope="row">${index}</th>
                            <td>${element.Pass}</td>
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
                            temp += `<td colspan="1" class="align-middle"><input type="number" class="form-control form-control-sm" max="100" min="0" value="${subject.Score}"></td>`
                            temp += `<td colspan="1" class="align-middle"><div class="custom-control custom-checkbox"><input type="checkbox" class="custom-control-input" id="check${subject.ID}" ${subject.Pass == 1 ? 'checked' : ''}><label class="custom-control-label" for="check${subject.ID}"></label></div></td>`

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
        }
    });
}