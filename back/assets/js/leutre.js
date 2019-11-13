let STATUS = {
    'enable': '啟用',
    'disabled': '停用',
}
let letures = []

const courseTable = $('table#leture').DataTable({
    processing: true,
    serverSide: true,
    ordering: false,
    ajax: {
        url: `${config.server}/v1/leture`,
        type: 'GET',
        dataSrc: (d) => {
            letures = []

            d.list.forEach((element, index, array) => {
                letures.push(Object.assign({}, element))

                array[index].Button = `
                    <button class="btn btn-secondary mr-1" onclick="detail(${index}, this)">查看</button>
                    <button class="btn btn-primary mr-1">編輯</button>
                    <button class="btn btn-danger">刪除</button>
                `

                array[index].Status = STATUS[element.Status]
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
        { data: "Name" },
        { data: "MinCredit" },
        { data: "Comment" },
        { data: "Status" },
        { data: "Button" },
    ],
    language: {
        url: '/assets/languages/chinese.json'
    },
});

const detail = (index, el) => {
    let leture = letures[index]

    $.ajax({
        url: `${config.server}/v1/leture/${leture.ID}`,
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