$(document).ready(() => {
    let converter = new showdown.Converter()

    $.ajax({
        url: "https://raw.githubusercontent.com/SuperEdge/nsysu-teacher-education/master/front/CHANGELOG.md",
        type: 'GET',
        success: (response) => {
            let html = converter.makeHtml(response)
            $('div#front').html(html)
        }
    });

    $.ajax({
        url: "https://raw.githubusercontent.com/SuperEdge/nsysu-teacher-education/master/back/CHANGELOG.md",
        type: 'GET',
        success: (response) => {
            let html = converter.makeHtml(response)
            $('div#back').html(html)
        }
    });

    $.ajax({
        url: "https://raw.githubusercontent.com/SuperEdge/nsysu-teacher-education/master/api/CHANGELOG.md",
        type: 'GET',
        success: (response) => {
            let html = converter.makeHtml(response)
            $('div#api').html(html)
        }
    });
})