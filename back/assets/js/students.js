const studentTable = $('table#students').DataTable({
    processing: true,
    serverSide: true,
    ordering: false,
    ajax: {
        url: `${config.server}/v1/users`,
        type: 'GET',
        dataSrc: (d) => {
            d.list.forEach((element, index, array) => {
                array[index].CreatedAt = dayjs(element.CreatedAt).format('YYYY-MM-DD HH:mm:ss')
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
        { data: "Account" },
        { data: "Major" },
        { data: "Number" },
        { data: "CreatedAt" },
    ],
    language: {
        url: '/assets/languages/chinese.json'
    },
});

$('input#upload').fileinput({
    language: 'zh-TW',
    theme: "fas",
    allowedFileExtensions: ['csv'],
    uploadUrl: `${config.server}/v1/users`,
    ajaxSettings: {
        headers: {
            'Authorization': `Bearer ${$.cookie('token')}`,
        }
    },
}).on('fileuploaded', (event, previewId, index, fileId) => {
    swal({
        title: '',
        text: '成功',
        icon: "success",
        timer: 1500,
        buttons: false,
    })
    studentTable.ajax.reload(null, false)
    $('input#upload').fileinput('clear')
}).on('fileuploaderror', (event, data, msg) => {
    swal({
        title: '',
        text: '新增失敗',
        icon: "error",
        timer: 1500,
        buttons: false,
    })
    $('div.kv-upload-progress.kv-hidden').css({ 'display': 'none' })
})