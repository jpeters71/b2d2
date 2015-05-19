
$(document).ready(function(){
    $('li').click(function(){
      
      $(this).addClass('active')
           .siblings()
           .removeClass('active');
    });                

    $("#mnu_drinks").click(function(){
        var html = "";
        $.getJSON("/svc/available-recipes", function(data) {
            html="<ul>";
            for (var i=0; i < data.length; i++) {
                var item = data[i];

                html += "<li class='drink_item' id='" + item.Id + "''><img height='100' width='100' src='data:image/png;base64," + item.Image + "'/>" + item.Title + "</li>";
            }
            html += "</ul>";
            $("#content").html(html);
        });
    });

    $("#mnu_recipes").click(function(){
        var html = "";
        $.getJSON("/svc/recipes", function(data) {
            html="<ul>";
            for (var i=0; i < data.length; i++) {
                var item = data[i];

                html += "<li class='drink_item' id='" + item.Id + "''><img height='100' width='100' src='data:image/png;base64," + item.Image + "'/>" + item.Title + "</li>";
            }
            html += "</ul>";
            $("#content").html(html);
        });
    });


    $("#mnu_ingredients").click(function(){
        var html = "";
        $.getJSON("/svc/ingredients", function(data) {
            html="<ul>";
            for (var i=0; i < data.length; i++) {
                var item = data[i];

                html += "<li class='drink_item' id='" + item.Id + "''>" + item.Title + "</li>";
            }
            html += "</ul>";
            $("#content").html(html);
        });
    });

});
