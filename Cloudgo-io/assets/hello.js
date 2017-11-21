$(document).ready(function() {
    $.ajax({
        url: "/jstest"
    }).then(function(data) {
        $('.title').append(data.title);
    });
});

