
$(document).ready(function(){
    $("#drink_popup").dialog({
        autoOpen: false
    });
    
    $("li").click(function(){
      $(this).addClass("active")
           .siblings()
           .removeClass("active");
    }); 
    
    $("#mnu_drinks").click(function(){
        var html = "";
        $.getJSON("/svc/available-recipes", function(data) {
            html="<ul>";
            for (var i=0; i < data.length; i++) {
                var item = data[i];

                html += "<li class='drink_item' id='" + item.Id + "''><img src='data:image/png;base64," + item.Image + "'/>" + item.Title + "</li>";
            }
            html += "</ul>";
            $("#title").html("Select Drink");
            $("#list").html(html);

            $(".drink_item").click(function() {
                alert("Test");
                alert("ID: " + this.id);
            });

        }); 
    });

    $("#mnu_recipes").click(function(){
        var html = "";
        $.getJSON("/svc/recipes", function(data) {
            html="<ul>";
            for (var i=0; i < data.length; i++) {
                var item = data[i];

                html += "<li class='drink_item' id='" + item.Id + "''><img src='data:image/png;base64," + item.Image + "'/>" + item.Title + "</li>";
            }
            html += "</ul>";
            $("#title").html("Recipes");
            $("#list").html(html);
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
            $("#title").html("Ingredients");
            $("#list").html(html);
        });
    });

    // Start it up
    $("#mnu_drinks").click();

});
