// Setup everyting.
var iCurrRepId = -1;

$(document).ready(function(){
    // This is the dialog for when a user selects a drink.
    $("#drink_popup").dialog({
        autoOpen: false,
        modal: true,
        width: 600,
        height: 350,
        buttons: {
                  "Make It": function() {
                        var title = $(this).dialog( "option", "title" );
                        $(this).dialog("close");
                        var repId = iCurrRepId;
                        $("#make_it_popup").dialog({
                            title: "Making a " + title
                        });

                        $("#make_it_popup_text").html("");
                        $("#make_it_popup_img").show();
                        $("#make_it_popup_img").attr("src", "styles/ajax-loader.gif");
                        $("#make_it_popup").dialog("open");
                        
                        // Run ajax to get this item
                        var jqxhr = $.ajax({
                            url: "/svc/recipes/" + iCurrRepId,
                            type: "POST", 
                            statusCode: {
                                201: function() {
                                    $("#make_it_popup_img").hide();
                                    $("#make_it_popup_text").html("Done!<br/><small>Share and Enjoy!&trade;</small>");
                                }
                            }                            
                        })
                        .fail(function(data) {
                            $("#make_it_popup_img").hide();
                            $("#make_it_popup_text").html("ERROR!<br/><em>" + data.responseText + "</em>");
                            
                        });
                  },
                  Cancel: function() {
                      $(this).dialog("close");
                  }
        },        
    });

    $("#make_it_popup").dialog({
        autoOpen: false,
        modal: true,
        width: 300,
        height: 300
    });

    $("#ingredient_popup").dialog({
        autoOpen: false,
        modal: true,
        width: 600,
        height: 435,
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

                html += "<li class='drink_item' id='" + item.ID + "''><img src='data:image/png;base64," + item.Image + "'/>" + item.Title + "</li>";
            }
            html += "</ul>";
            $("#title").html("Select Drink");
            $("#list").html(html);

            $(".drink_item").click(function() {
                // Set things up.  Start by clearing stuff
                $("#drink_popup_text").html("");
                $("#drink_popup_img").attr("src", "styles/ajax-loader.gif");
                iCurrRepId = this.id;
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

                html += "<li class='drink_item' id='" + item.ID + "''><img src='data:image/png;base64," + item.Image + "'/>" + item.Title + "</li>";
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

                html += "<li class='ingredient_item ' id='" + item.ID + "''>" + item.Title + "</li>";
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
                    $("#ingredient_popup_pump").val(data.PumpID);
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
