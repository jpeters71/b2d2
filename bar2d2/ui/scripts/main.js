// Setup everyting.
$(document).ready(function(){
    // This is the dialog for when a user selects a drink.
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
    $("#ingredient_popup").dialog({
        autoOpen: false,
        modal: true,
        width: 600,
        height: 410,
        buttons: {
                  "Update": function() {
                      $(this).dialog("close");
                  },
                  Cancel: function() {
                      $(this).dialog("close");
                  }
        },        
    });
    
    // Handles setting the active menu item.
    $("li").click(function(){
      $(this).addClass("active")
           .siblings()
           .removeClass("active");
    }); 
    
    // When the drinks menu item is selected.
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
    
    // When the recipe menu item is selected.
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
    
    /*    
            <div id="ingredient_popup" title="Ingredient!">
            <div id="ingredient_popup_brand"/>
            <div id="ingredient_popup_proof"/>
            <div id="ingredient_popup_brand"/>
            <div id="ingredient_popup_description"/>
        </div>
*/

    // When the ingredients menu item is selected
    $("#mnu_ingredients").click(function(){
        var html = "";
        $.getJSON("/svc/ingredients", function(data) {
            html="<ul>";
            for (var i=0; i < data.length; i++) {
                var item = data[i];

                html += "<li class='ingredient_item ' id='" + item.Id + "''>" + item.Title + "</li>";
            }
            html += "</ul>";
            $("#title").html("Ingredients");
            $("#list").html(html);

            $(".ingredient_item").click(function() {
                // Set things up.  Start by clearing stuff
                $("#ingredient_popup_brand").val("");
                $("#ingredient_popup_proof").val("");
                $("#ingredient_popup_pump").val("");
                $("#ingredient_popup_description").val("");
                $("#ingredient_popup").dialog("open");
                
                // Run ajax to get this item
                $.getJSON("/svc/ingredients/" + this.id, function(data) {
                    $("#ingredient_popup").dialog({
                        title: data.Title
                    });
                    $("#ingredient_popup_brand").val(data.Brand);
                    $("#ingredient_popup_proof").val(data.AlcoholContent * 2);
                    $("#ingredient_popup_pump").val(data.PumpId);
                    $("#ingredient_popup_description").val(data.Description);
                });
            });
        });
    });

    // When the ingredients menu item is selected
    $("#mnu_pumps").click(function(){
        var html = "";
        $("#list").html("TODO!");
    });

    // Start it up
    $("#mnu_drinks").click();

});
