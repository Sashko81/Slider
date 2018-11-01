
$(document).ready(function() {
	$("#btn").click(function() {
		$.get( "/", function( data ) {
		    console.log( "number = " + data );
		    $(".check").each(function( index ) {
			    $( this ).removeAttr("checked");
		    });
		    $("#f"+data).attr("checked", "checked")
		});
	});
});