let STATUS = {
    'enable': '啟用',
    'disable': '停用',
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
                    <button class="btn btn-secondary mr-1" onclick="detail(${index})">查看</button>
                    <button class="btn btn-primary mr-1" onclick="update(${index})">編輯</button>
                    <button class="btn btn-danger" onclick="$('#deleteModal').modal('show'); courseID=${element.ID}">刪除</button>
                `

                array[index].Status = STATUS[element.Status]
            })
            return d.list
        },
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr, error, thrown) => {
            error(xhr, xhr.responseText)
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

const detail = (index) => {
    let leture = letures[index]

    $.ajax({
        url: `${config.server}/v1/leture/${leture.ID}`,
        type: 'GET',
        error: (xhr) => {
            error(xhr, '失敗')
        },
        beforeSend: (xhr) => {
            setHeader(xhr)
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

                        if (type.MinCredit > 1) {
                            comment += `${type.Name}至少修習${type.MinCredit}學分<br>`
                        }

                        for (let group of type.Groups) {
                            subjects += group.Subjects.length

                            for (let subject of group.Subjects) {
                                let temp = '<tr>'
                                let condition1 = category.Types.indexOf(type) === category.Types.length - 1
                                let condition2 = group.Subjects.indexOf(subject) === group.Subjects.length - 1

                                if (condition1 && condition2) {
                                    temp += `<td colspan="2" rowspan="${subjects}" class="align-middle">${category.Name}</td>`
                                }

                                if (condition2) {
                                    temp += `<td colspan="2" rowspan="${group.Subjects.length}" class="align-middle">${type.Name}</td>`
                                }

                                temp += `<td colspan="1" class="align-middle">${subject.Compulsory ? "必修" : "選修"}</td>`
                                temp += `<td colspan="4" class="align-middle">${subject.Name}</td>`
                                temp += `<td colspan="1" class="align-middle">${subject.Credit}</td>`

                                if (condition1 && condition2) {
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