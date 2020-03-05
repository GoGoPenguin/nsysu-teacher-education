let editedItem = null
let student = null

loading()
$(document).ready(() => {
    Promise.all([
        getStudentLecture(),
        getStudentInformation()
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

const getStudentLecture = () => {
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
                $('#student-lecture tbody').html(`
                    <tr>
                        <td scope="row" colspan="8" style="text-align: center">尚無資料</td>
                    </tr>
                `)
            } else {
                let html = ''
                response.list.forEach((element) => {
                    html += `
                        <tr>
                            <td data-title="狀態" class="${element.Pass ? 'success' : 'danger'}"><span>●</span>${element.Pass ? '通過' : '未通過'}</td>
                            <td data-title="科目名稱">${element.Leture.Name}</td>
                            <td data-title="最低學分">${element.Leture.MinCredit}</td>
                            <td data-title="備註">${element.Leture.Comment}</td>
                            <td><a class="btn_table" onclick="getStudentLectureDetail(${element.ID})">編輯</a></td>
                        </tr>
                    `
                })
                $('#student-lecture tbody').html(html)
            }
        }
    });
}

const getStudentLectureDetail = (id) => {
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
                let lecture = response.data.Leture
                let html = ''

                for (const category of lecture.Categories) {
                    let content = ''
                    let comment = ''
                    let subjects = 0

                    if (category.MinCredit > 1) {
                        comment += `總共至少修習${category.MinCredit}學分<br><br>`
                    }

                    if (category.MinType > 1) {
                        comment += `總共至少修習${category.MinType}專長<br><br>`
                    }

                    for (const type of category.Types) {
                        let condition1 = category.Types.indexOf(type) === category.Types.length - 1
                        let subjectGroups = 0

                        if (type.MinCredit > 1) {
                            comment += `${type.Name}至少修習${type.MinCredit}學分<br><br>`
                        }

                        for (const group of type.Groups) {
                            let condition2 = type.Groups.indexOf(group) === type.Groups.length - 1

                            subjects += group.Subjects.length
                            subjectGroups += group.Subjects.length

                            for (const subject of group.Subjects) {
                                let condition3 = group.Subjects.indexOf(subject) === group.Subjects.length - 1
                                let temp = '<tr>'

                                if (condition1 && condition2 && condition3) {
                                    temp += `<td colspan="2" rowspan="${subjects}" id="category${category.ID}">${category.Name}</td>`
                                }

                                if (condition2 && condition3) {
                                    temp += `<td colspan="2" rowspan="${subjectGroups}" id="type${type.ID}">${type.Name}</td>`
                                }

                                temp += `<td colspan="1">${subject.Compulsory ? "必修" : "選修"}</td>`

                                if (group.MinCredit > 0) {
                                    temp += `<td colspan="3">${subject.Name}</td>`

                                    if (condition3) {
                                        temp += `<td colspan="1" rowspan="${group.Subjects.length}" class="align-middle vericaltext" id="group${group.ID}">至少${group.MinCredit}學分</td>`
                                    }
                                } else {
                                    temp += `<td colspan="4" id="subject${subject.ID}">${subject.Name}</td>`
                                }

                                temp += `<td colspan="1">${subject.Credit}</td>`
                                temp += `
                                    <td colspan="1">
                                        <input type="text" value="${subject.Year}" onblur="updateSubject(${id}, ${subject.ID})" id="year${subject.ID}">
                                    </td>`
                                temp += `
                                    <td colspan="1">
                                        <input type="text" value="${subject.Semester}" onblur="updateSubject(${id}, ${subject.ID})" id="semester${subject.ID}">
                                    </td>`
                                temp += `
                                    <td colspan="1">
                                        <input type="text" value="${subject.StudentName}" onblur="updateSubject(${id}, ${subject.ID})" id="studentName${subject.ID}">
                                    </td>`
                                temp += `
                                    <td colspan="1">
                                        <input type="text" value="${subject.StudentCredit}" onblur="updateSubject(${id}, ${subject.ID})" id="studentCredit${subject.ID}">
                                    </td>`
                                temp += `
                                    <td colspan="1">
                                        <input type="text" value="${subject.Score}" onblur="updateSubject(${id}, ${subject.ID})" id="score${subject.ID}">
                                    </td>`

                                if (condition1 && condition2 && condition3) {
                                    temp += `<td colspan="2" rowspan="${subjects}">${comment}</td>`
                                }

                                temp += `</tr>`
                                content = `${temp}${content}`
                            }
                        }
                    }

                    html += content
                }

                $('#detailModal .modal-title').html(lecture.Name)
                $('#detailModal #name').html(lecture.Name)
                $('#detailModal #min_credit').html(lecture.MinCredit)
                $('#detailModal #comment').html(lecture.Comment)
                $('#detailModal #categories').html(html)

                let buttons = '<button class="btn btn-primary" type="button" onclick="check()">審核</button>'
                if (response.data.Pass) {
                    buttons += '<button class="btn btn-secondary ml-3" onclick="applictionForm(this)">下載申請書</button>'
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

const updateSubject = (studentLectureID, subjectID) => {
    $.ajax({
        url: `${config.server}/v1/leture/student/subject`,
        type: 'PATCH',
        data: {
            'StudentLetureID': studentLectureID,
            'SubjectID': subjectID,
            'Name': $(`#studentName${subjectID}`).val(),
            'Year': $(`#year${subjectID}`).val(),
            'Semester': $(`#semester${subjectID}`).val(),
            'Credit': $(`#studentCredit${subjectID}`).val(),
            'Score': $(`#score${subjectID}`).val(),
        },
        beforeSend: (xhr) => {
            setHeader(xhr)
        },
        error: (xhr) => {
            errorHandle(xhr, "錯誤")
        },
        success: (response) => {
            if (response.code === 0) {
                reloadLecture(studentLectureID)
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

const reloadLecture = (id) => {
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
    let lecture = editedItem
    let pass = true

    for (const category of lecture.Categories) {
        let categoryCredit = 0
        let categoryTypes = 0

        for (const type of category.Types) {
            let typeCredit = 0

            for (const group of type.Groups) {
                let groupCredit = 0

                for (const subject of group.Subjects) {
                    if (
                        subject.Year !== '' &&
                        subject.Semester !== '' &&
                        subject.StudentName !== '' &&
                        subject.StudentCredit !== '' &&
                        subject.Score !== ''
                    ) {
                        groupCredit += subject.Credit
                    } else {
                        if (subject.Compulsory) {
                            if (!$(`#subject${subject.ID}`).hasClass('text-danger')) {
                                $(`#subject${subject.ID}`).addClass('text-danger')
                            }
                            pass = false
                        } else {
                            $(`#subject${subject.ID}`).removeClass('text-danger')
                        }
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
        updateStudentLecturePass(lecture.ID, true)

        let buttons = '<button class="btn btn-primary" type="button" onclick="check()">審核</button>'
        buttons += '<button class="btn btn-secondary ml-3" onclick="applictionForm(this)">下載申請書</button>'
        $('#detailModal .modal-footer').html(buttons)
    } else {
        swal({
            title: '',
            text: '未通過',
            icon: "error",
            timer: 1500,
            buttons: false,
        })
        updateStudentLecturePass(lecture.ID, false)

        let buttons = '<button class="btn btn-primary" type="button" onclick="check()">審核</button>'
        $('#detailModal .modal-footer').html(buttons)
    }
}

const updateStudentLecturePass = (lectureID, pass) => {
    $.ajax({
        url: `${config.server}/v1/leture/student/pass`,
        type: 'PATCH',
        data: {
            'LetureID': lectureID,
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
            getStudentLecture()
        }
    })
}

const applictionForm = (el) => {
    $(el).addClass('disabled')

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

            let footer = '【請單面列印】'
            let pageSize = doc.internal.pageSize;
            let pageHeight = pageSize.height ? pageSize.height : pageSize.getHeight();
            doc.setFontSize(10);
            doc.center(footer, pageHeight - 5);
        }

        doc.autoTable({
            theme: 'grid',
            headStyles: styleDef,
            footStyle: styleDef,
            styles: styleDef,
            margin: {
                top: 20
            },
            didDrawPage: didDrawPage,
            body: [
                [{
                    content: '姓名',
                    styles: {
                        cellWidth: cellWidth
                    }
                },
                {
                    content: student.Name,
                    styles: {
                        cellWidth: cellWidth
                    }
                },
                {
                    content: '聯絡電話',
                    styles: {
                        cellWidth: cellWidth
                    }
                },
                {
                    content: ''
                },
                {
                    content: '身分證字號',
                    styles: {
                        cellWidth: cellWidth
                    }
                },
                {
                    content: '',
                    colSpan: 2
                },
                ],
                [{
                    content: '畢業學校系所',
                    styles: {
                        cellWidth: cellWidth
                    }
                },
                {
                    content: student.Major,
                    colSpan: 3
                },
                {
                    content: '學號',
                    styles: {
                        cellWidth: cellWidth
                    }
                },
                {
                    content: student.Account,
                    colSpan: 2
                },
                ],
                [{
                    content: '申請任教科別',
                    styles: {
                        cellWidth: cellWidth
                    }
                },
                {
                    content: '',
                    colSpan: 6
                },
                ],
                [{
                    content: '資格',
                    styles: {
                        cellWidth: cellWidth
                    }
                },
                {
                    content: '□ 參加教師資格考試\n□ 加科/加另一類科（請附教師證書影本）',
                    colSpan: 3
                },
                {
                    content: '學程編號',
                    styles: {
                        cellWidth: cellWidth
                    }
                },
                {
                    content: student.Number,
                    colSpan: 2
                },
                ],
                [{
                    content: '修習起訖期間',
                    styles: {
                        cellWidth: cellWidth
                    }
                },
                {
                    content: '   年   月   日～    年   月   日（   年   月   日～   年   月   日他校學分採認）',
                    colSpan: 6
                },
                ],
                [{
                    content: '教育部核定專門課程文號：   年   月   日臺教師(   )字第                    號函',
                    colSpan: 7,
                    styles: {
                        cellWidth: cellWidth
                    }
                },],
            ],
        });

        let lecture = {
            startY: doc.autoTable.previous.finalY + 5,
            headStyles: styleDef,
            footStyle: styleDef,
            styles: styleDef,
            margin: {
                top: 20
            },
            showHead: 'everyPage',
            rowPageBreak: 'avoid',
            didDrawPage: didDrawPage,
            head: [
                [{
                    content: '師資生自行填寫（請用電腦打字）',
                    colSpan: 10,
                    styles: {
                        halign: 'center'
                    }
                },
                {
                    content: '系所審查意見',
                    colSpan: 3,
                    rowSpan: 2,
                    styles: {
                        halign: 'center'
                    }
                },
                ],
                [{
                    content: '編號',
                    colSpan: 1,
                    rowSpan: 2,
                    styles: {
                        cellWidth: 9,
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '課程類別',
                    colSpan: 2,
                    rowSpan: 2,
                    styles: {
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '教育部核定科目',
                    colSpan: 2,
                    cellWidth: 15,
                    styles: {
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '師資生已修習科目（依成績單確實填寫）',
                    colSpan: 5,
                    styles: {
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                ],
                [{
                    content: '科目名稱',
                    cellWidth: 15,
                    styles: {
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '學分',
                    styles: {
                        cellWidth: 9,
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '學年',
                    styles: {
                        cellWidth: 9,
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '學期',
                    styles: {
                        cellWidth: 9,
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '科目名稱',
                    cellWidth: 15,
                    styles: {
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '學分',
                    styles: {
                        cellWidth: 9,
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '成績',
                    styles: {
                        cellWidth: 9,
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '完全採認',
                    styles: {
                        cellWidth: 12,
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '不能採認',
                    styles: {
                        cellWidth: 12,
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '系主任簽章',
                    styles: {
                        cellWidth: 12,
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                ],
            ],
            body: [],
        }

        for (const category of editedItem.Categories) {
            for (const type of category.Types) {
                let temp = []
                for (const group of type.Groups) {
                    for (const subject of group.Subjects) {
                        if (subject.Year !== '' &&
                            subject.Semester !== '' &&
                            subject.StudentName !== '' &&
                            subject.StudentCredit !== '' &&
                            subject.Score !== ''
                        ) {
                            temp.push([{
                                content: '',
                                styles: {
                                    fillColor: white
                                }
                            },
                            {
                                content: subject.Compulsory ? '必修' : '選修',
                                styles: {
                                    fillColor: white,
                                    cellWidth: 7
                                }
                            },
                            {
                                content: subject.Name,
                                styles: {
                                    fillColor: white
                                }
                            },
                            {
                                content: subject.Credit,
                                styles: {
                                    fillColor: white,
                                    halign: 'center'
                                }
                            },
                            {
                                content: subject.Year.split(',').join('\n'),
                                styles: {
                                    fillColor: white,
                                    halign: 'center'
                                }
                            },
                            {
                                content: subject.Semester.split(',').join('\n'),
                                styles: {
                                    fillColor: white,
                                    halign: 'center'
                                }
                            },
                            {
                                content: subject.StudentName.split(',').join('\n'),
                                styles: {
                                    fillColor: white,
                                    halign: 'center'
                                }
                            },
                            {
                                content: subject.StudentCredit.split(',').join('\n'),
                                styles: {
                                    fillColor: white,
                                    halign: 'center'
                                }
                            },
                            {
                                content: subject.Score.split(',').join('\n'),
                                styles: {
                                    fillColor: white,
                                    halign: 'center'
                                }
                            },
                            {
                                content: '',
                                styles: {
                                    fillColor: white
                                }
                            },
                            {
                                content: '',
                                styles: {
                                    fillColor: white
                                }
                            },
                            {
                                content: '',
                                styles: {
                                    fillColor: white
                                }
                            },
                            ])
                        }
                    }
                }
                if (temp.length > 0) {
                    temp[0].splice(1, 0, {
                        content: type.Name,
                        rowSpan: temp.length,
                        styles: {
                            fillColor: white,
                            cellWidth: 7
                        }
                    })
                    lecture.body = lecture.body.concat(temp);
                }
            }
        }

        doc.autoTable(lecture);

        let result = {
            theme: 'plain',
            startY: doc.autoTable.previous.finalY + 5,
            styles: styleDef,
            margin: {
                top: 20
            },
            rowPageBreak: 'avoid',
            didDrawPage: didDrawPage,
            body: [],
        }

        let content = '【審查結果】\n'
        for (const category of editedItem.Categories) {
            for (const type of category.Types) {
                content += `${category.Name}${type.Name}____學分，`
            }
        }
        result.body.push([{
            content: content
        }])

        doc.autoTable(result);

        content = '附繳資料及注意事項：\n'
        content += '1.該科目之專門課程科目及學分一覽表（請將表訂科目編號對應填寫至一覽表上）。\n'
        content += '2.歷年成績單正本/第二專長學分證明（請將表訂科目編號對應填寫至成績單上）。\n'
        content += '3.畢業證書影本\n'
        content += '4.課程大綱（科目名稱不同者務必檢附）。\n'
        content += '5.申請加科/加另一類科者另須繳交審查費：1200元（請至本校線上收款系統列印繳付：\n'
        content += 'http://140.117.13.70/OLPRS/pay.asp，繳款師資培育中心師資職前教育專門課程審查費）'

        doc.autoTable({
            startY: doc.autoTable.previous.finalY + 5,
            styles: styleDef,
            margin: {
                top: 20
            },
            rowPageBreak: 'avoid',
            didDrawPage: didDrawPage,
            body: [
                [{
                    content: '審查順序：師培中心彙整表件、初審→各系所審查、採認→師培中心複審、建檔→教務處註冊組製證',
                    colSpan: 4,
                    styles: {
                        fillColor: white
                    }
                },],
                [{
                    content: '申請人簽章',
                    styles: {
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '師培中心承辦人核章',
                    styles: {
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '系所主任核章',
                    styles: {
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                {
                    content: '教務處註冊組核章',
                    styles: {
                        fillColor: gray,
                        halign: 'center'
                    }
                },
                ],
                [{
                    content: '',
                    rowSpan: 2,
                    styles: {
                        fillColor: white
                    }
                },
                {
                    content: '',
                    styles: {
                        fillColor: white
                    }
                },
                {
                    content: '',
                    rowSpan: 2,
                    styles: {
                        fillColor: white
                    }
                },
                {
                    content: '',
                    rowSpan: 2,
                    styles: {
                        fillColor: white
                    }
                },
                ],
                [{
                    content: '',
                    styles: {
                        fillColor: white
                    }
                },],
                [{
                    content: content,
                    colSpan: 4,
                    styles: {
                        fillColor: white
                    }
                },],
            ],
        })

        doc.save(`${editedItem.Name}申請表.pdf`)

        $(el).removeClass('disabled')
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