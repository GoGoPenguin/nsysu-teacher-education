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

                array[index].Information = `<div onclick="getInformation(${element.ID}, '${element.Information}')">${element.Information}</div>`
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
                $('#detailModal .modal-title').html(leture.Name)
                $('#detailModal #name').html(leture.Name)
                $('#detailModal #min_credit').html(leture.MinCredit)
                $('#detailModal #comment').html(leture.Comment)

                let categories = ''
                leture.Categories.forEach(category => {
                    let types = ''
                    let groups = ''

                    category.Types.forEach(type => {
                        types += `
                            <div class='row'>
                                <div class='col border border-dark'>${type.Name}</div>
                            </div>
                        `

                        type.Groups.forEach(group => {
                            let subjects = ``
                            group.Subjects.forEach(subject => {
                                subjects += `
                                    <div class='row'>
                                        <div class='col border border-dark'>${subject.Name}</div>
                                    </div>
                                `
                            })

                            groups += `
                                <div class='row'>
                                    <div class='col'>
                                        <div class='row'>
                                            <div class='col'>${subjects}</div>
                                        </div>
                                    </div>
                                </div>
                            `
                        })
                    })



                    categories += `
                        <div class='row'>
                            <div class='col'>
                                <div class='row'>
                                    <div class='col border border-dark'>${category.Name}</div>
                                    <div class='col'>${types}</div>
                                </div>
                            </div>
                            <div class='col'>${groups}</div>
                            <div class='col-2'></div>
                            <div class='col'></div>
                        </div>
                    `
                })

                $('#detailModal #categories').html(categories)
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