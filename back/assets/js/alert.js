$('div#alert').on('click', 'button.close', () => {
    $('div#alert').hide('slow')
})

const alert = (msg) => {
    $('div#alert').addClass('alert alert-danger')
    $('div#alert').html(`${msg}<button type="button" class="close"><span aria-hidden="true">&times;</span></button>`)
    $('div#alert').show('fast')
}