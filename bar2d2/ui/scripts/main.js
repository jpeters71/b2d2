
$(document).ready(function(){
    $("#drink_popup").dialog({
        autoOpen: false,
        modal: true,
        width: 600,
        height: 350,
        buttons: {
                  "Make It": function() {
                      $(this).dialog("close");
                  },
                  Cancel: function() {
                      $(this).dialog("close");
                  }
        },        
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
                // Set things up.  Start by clearing stuff
                $("#drink_popup_text").html("");
                $("#drink_popup_img").attr("src", "styles/ajax-loader.gif");
                $("#drink_popup").dialog("open");
                
                // Run ajax to get this item
                $.getJSON("/svc/recipes/" + this.id, function(data) {
                    $("#drink_popup").dialog({
                        title: data.Title
                    });
                    $("#drink_popup_text").html(data.Description);
                    $("#drink_popup_img").attr("src", "data:image/png;base64," + data.Image);
                });
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
