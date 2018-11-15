$(document).ready (function(){
     $("#apibtn").click(function() {
        $.ajax({
            url: "/api/test"
        }).then(function(data) {
        $('.apishow').append(data.show);
        });
    });
})