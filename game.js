$(document).ready(function () {
    $("#button1").click(function (event) {
        $.ajax({
            type: 'GET',
            crossOrigin: true,
            url: 'http://localhost:8080/users/Tictactoe',
            success: function (result) {
                $("#target").html('<table class="table table-bordered"><tr>' +
                    '<td><gamebuttom val="1"  class="btn btn-default">' + result.field[0] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="2"  class="btn btn-default">' + result.field[1] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="3"  class="btn btn-default">' + result.field[2] + "</gamebuttom></td><tr>" +
                    '<td><gamebuttom val="4"  class="btn btn-default">' + result.field[3] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="5"  class="btn btn-default">' + result.field[4] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="6"  class="btn btn-default">' + result.field[5] + "</gamebuttom></td><tr>" +
                    '<td><gamebuttom val="7"  class="btn btn-default">' + result.field[6] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="8"  class="btn btn-default">' + result.field[7] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="9"  class="btn btn-default">' + result.field[8] + "</gamebuttom></td>"
                    + "</tr> </table><div>" + result.currentPlayer + "</div>");
                myfunction();

            },
            error: function () {

            }

        });


    });
});

function myfunction() {
    $("gamebuttom").click(function (event) {
        $.ajax({
            type: 'POST',
            crossOrigin: true,
            url: 'http://localhost:8080/users/Tictactoe/move/' + $(event.target).attr("val"),
            success: function (result) {
                $("#target").html('<table class="table table-bordered"><tr>' +
                    '<td><gamebuttom val="1"  class="btn btn-default">' + result.field[0] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="2"  class="btn btn-default">' + result.field[1] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="3"  class="btn btn-default">' + result.field[2] + "</gamebuttom></td><tr>" +
                    '<td><gamebuttom val="4"  class="btn btn-default">' + result.field[3] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="5"  class="btn btn-default">' + result.field[4] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="6"  class="btn btn-default">' + result.field[5] + "</gamebuttom></td><tr>" +
                    '<td><gamebuttom val="7"  class="btn btn-default">' + result.field[6] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="8"  class="btn btn-default">' + result.field[7] + "</gamebuttom></td>" +
                    '<td><gamebuttom val="9"  class="btn btn-default">' + result.field[8] + "</gamebuttom></td>"
                    + "</tr> </table><div>" + result.currentPlayer + "</div>" +"<div>"+result.message +"</div>" );
                myfunction();
            },
            error: function () {

            }
        });
    });
};

