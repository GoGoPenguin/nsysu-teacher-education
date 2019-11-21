let editedItem = null

loading()
$(document).ready(() => {
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
                $('#student-leture tbody').html(`
                    <tr>
                        <td scope="row" colspan="8" style="text-align: center">尚無資料</td>
                    </tr>
                `)
            } else {
                let html = ''
                response.list.forEach((element, index) => {
                    html += `
                        <tr>
                            <th scope="row">${index}</th>
                            <td class="${element.Pass ? 'text-success' : 'text-danger'}">${element.Pass ? '通過' : '未通過'}</td>
                            <td>${element.Leture.Name}</td>
                            <td>${element.Leture.MinCredit}</td>
                            <td>${element.Leture.Comment}</td>
                            <td><button class="btn btn-primary" onclick="getStudentLetureDetail(${element.ID})">編輯</button></td>
                        </tr>
                    `
                })
                $('#student-leture tbody').html(html)
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

                                temp += `<td colspan="1" class="align-middle">${subject.Compulsory ? "必修" : "選修"}</td>`

                                if (group.MinCredit > 0) {
                                    temp += `<td colspan="3" class="align-middle">${subject.Name}</td>`

                                    if (condition3) {
                                        temp += `<td colspan="1" rowspan="${group.Subjects.length}" class="align-middle vericaltext" id="group${group.ID}">至少${group.MinCredit}學分</td>`
                                    }
                                } else {
                                    temp += `<td colspan="4" class="align-middle" id="subject${subject.ID}">${subject.Name}</td>`
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

                let buttons = '<button class="btn btn-primary" type="button" onclick="check()">審核</button>'
                if (response.data.Pass) {
                    buttons += '<button class="btn btn-secondary ml-3" onclick="applictionForm(this)">下載申請書</button>'
                    $('#detailModal .modal-footer').html()
                }

                $('#detailModal .modal-footer').html(buttons)
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
        updateStudentLeturePass(leture.ID, true)
    } else {
        swal({
            title: '',
            text: '未通過',
            icon: "error",
            timer: 1500,
            buttons: false,
        })
        updateStudentLeturePass(leture.ID, false)
    }
}

const updateStudentLeturePass = (letureID, pass) => {
    $.ajax({
        url: `${config.server}/v1/leture/student/pass`,
        type: 'PATCH',
        data: {
            'LetureID': letureID,
            'Pass': pass,
        },
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.code !== 0) {
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
            getStudentLeture()
        }
    })
}

const applictionForm = (el) => {
    $(el).attr('disabled', true)
    $(el).html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>&nbsp下載中...')

    setTimeout(() => {
        let doc = new jsPDF()
        doc.setTextColor(0);
        doc.setFont('kaiu')
        doc.setFontStyle('normal');
        doc.setFontSize(10);

        const white = 255
        const black = 0
        const gray = 230
        const cellWidth = 25
        const styleDef = {
            font: 'kaiu',
            fontStyle: 'normal',
            fillColor: white,
            textColor: black,
            lineColor: black,
            lineWidth: 0.1,
            valign: 'middle',
        }
        const didDrawPage = (HookData) => {
            let header = '國立中山大學中等學校各任教學科專門科目學分認定申請表'
            doc.setFontSize(16);
            doc.center(header, 15);

            let footer = '【請單面列印，未申請採認之科目請刪除(含表格)】'
            let pageSize = doc.internal.pageSize;
            let pageHeight = pageSize.height ? pageSize.height : pageSize.getHeight();
            doc.setFontSize(10);
            doc.center(footer, pageHeight - 5);
        }

        let information = {
            theme: 'grid',
            headStyles: styleDef,
            footStyle: styleDef,
            styles: styleDef,
            margin: { top: 20 },
            body: [
                [
                    { content: '姓名', styles: { cellWidth: cellWidth } },
                    { content: '' },
                    { content: '聯絡電話', styles: { cellWidth: cellWidth } },
                    { content: '' },
                    { content: '身分證字號', styles: { cellWidth: cellWidth } },
                    { content: '', colSpan: 2 },
                ],
                [
                    { content: '畢業學校系所', styles: { cellWidth: cellWidth } },
                    { content: '', colSpan: 3 },
                    { content: '學號', styles: { cellWidth: cellWidth } },
                    { content: '', colSpan: 2 },
                ],
                [
                    { content: '申請任教科別', styles: { cellWidth: cellWidth } },
                    { content: '', colSpan: 6 },
                ],
                [
                    { content: '資格', styles: { cellWidth: cellWidth } },
                    { content: '□ 參加教師資格考試\n□ 加科/加另一類科（請附教師證書影本）', colSpan: 3 },
                    { content: '學程編號', styles: { cellWidth: cellWidth } },
                    { content: '', colSpan: 2 },
                ],
                [
                    { content: '修習起訖期間', styles: { cellWidth: cellWidth } },
                    { content: '   年   月   日～    年   月   日（   年   月   日～   年   月   日他校學分採認）', colSpan: 6 },
                ],
                [
                    { content: '教育部核定專門課程文號：   年   月   日臺教師(   )字第                    號函', colSpan: 7, styles: { cellWidth: cellWidth } },
                ],
            ],
        }
        doc.autoTable(information);

        let leture = {
            startY: doc.autoTable.previous.finalY + 5,
            headStyles: styleDef,
            footStyle: styleDef,
            margin: { top: 20 },
            showHead: 'everyPage',
            rowPageBreak: 'avoid',
            styles: styleDef,
            didDrawPage: didDrawPage,
            head: [
                [
                    { content: '師資生自行填寫（請用電腦打字）', colSpan: 10, styles: { halign: 'center' } },
                    { content: '系所審查意見', colSpan: 3, rowSpan: 2, styles: { halign: 'center' } },
                ],
                [
                    { content: '編號', colSpan: 1, rowSpan: 2, styles: { cellWidth: 7, fillColor: gray, halign: 'center' } },
                    { content: '課程類別', colSpan: 2, rowSpan: 2, styles: { fillColor: gray, halign: 'center' } },
                    { content: '教育部核定科目', colSpan: 2, cellWidth: 15, styles: { fillColor: gray, halign: 'center' } },
                    { content: '師資生已修習科目（依成績單確實填寫）', colSpan: 5, styles: { fillColor: gray, halign: 'center' } },
                ],
                [
                    { content: '科目名稱', cellWidth: 15, styles: { fillColor: gray, halign: 'center' } },
                    { content: '學分', styles: { cellWidth: 7, fillColor: gray, halign: 'center' } },
                    { content: '學年', styles: { cellWidth: 7, fillColor: gray, halign: 'center' } },
                    { content: '學期', styles: { cellWidth: 7, fillColor: gray, halign: 'center' } },
                    { content: '科目名稱', cellWidth: 15, styles: { fillColor: gray, halign: 'center' } },
                    { content: '學分', styles: { cellWidth: 7, fillColor: gray, halign: 'center' } },
                    { content: '成績', styles: { cellWidth: 7, fillColor: gray, halign: 'center' } },
                    { content: '完全採認', styles: { cellWidth: 12, fillColor: gray, halign: 'center' } },
                    { content: '不能採認', styles: { cellWidth: 12, fillColor: gray, halign: 'center' } },
                    { content: '系主任簽章', styles: { cellWidth: 12, fillColor: gray, halign: 'center' } },
                ],
            ],
            body: [],
        }

        for (let category of editedItem.Categories) {
            for (let type of category.Types) {
                for (let group of type.Groups) {
                    for (let subject of group.Subjects) {
                        leture.body.push([
                            { content: '', styles: { fillColor: white } },
                            { content: '', styles: { fillColor: white, cellWidth: 7 } },
                            { content: subject.Compulsory ? '必修' : '選修', styles: { fillColor: white, cellWidth: 7 } },
                            { content: subject.Name, styles: { fillColor: white } },
                            { content: subject.Credit, styles: { fillColor: white } },
                            { content: '', styles: { fillColor: white } },
                            { content: '', styles: { fillColor: white } },
                            { content: '', styles: { fillColor: white } },
                            { content: '', styles: { fillColor: white } },
                            { content: '', styles: { fillColor: white } },
                            { content: '', styles: { fillColor: white } },
                            { content: '', styles: { fillColor: white } },
                            { content: '', styles: { fillColor: white } },
                        ])
                    }
                }
            }
        }

        doc.autoTable(leture);
        doc.output('dataurlnewwindow')

        $(el).attr('disabled', false)
        $(el).html('下載申請書')
    }, 1)
}

$('#detailModal').on('hidden.bs.modal', () => {
    editedItem = null
});

(function (API) {
    API.center = function (txt, y) {
        // Get current font size
        let fontSize = this.internal.getFontSize();

        // Get page width
        let pageWidth = this.internal.pageSize.width;

        // Get the actual text's width
        // You multiply the unit width of your string by your font size and divide
        // by the internal scale factor. The division is necessary
        // for the case where you use units other than 'pt' in the constructor
        // of jsPDF.
        //
        txtWidth = this.getStringUnitWidth(txt) * fontSize / this.internal.scaleFactor;

        // Calculate text's x coordinate
        x = (pageWidth - txtWidth) / 2;

        // Draw text at x,y
        this.text(txt, x, y);
    }
})(jsPDF.API);