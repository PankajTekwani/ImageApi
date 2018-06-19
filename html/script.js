$(document).ready(function() {
  $('#btn').click(function() {
    let vTag = $('#tag').val();
    $('#location').val("");
    $.ajax({
      url: `http://localhost:8080/` + vTag,
      type: 'GET',
      data: {
        format: 'json'
      },
      success: function(jsonData) {
		  $(".heading").empty();
		  $(".wrapper").empty();
		if(jsonData.length > 0){
		  $( "<h1>" + vTag + "</h1><br />" ).appendTo('.heading');
		  console.log(jsonData);
				var rowsNum = jsonData.data.length / 5;
				for (var i = 0; i < jsonData.data.length; i++) {
					var image = jsonData.data[i];
					createImg(image.URL, image.probablity);
					console.log(image.URL);
					console.log(image.probablity);
				}
		}
		else
		{		
		  $( "<h3>No Image Corresponding to the Tag " + vTag + "</h3>" ).appendTo('.heading');
		}
      },
      error: function() {
        $('#errors').text("There was an error processing your request. Please try again.")
      }
    });
  });
});
function createImg(imgSrc, probability) {
		$( "<div class = 'box'><img id='Myid' src='" + imgSrc + "'  class='img-rounded' style='height: 150px;width: 150px;'><p>Value: " + probability + "</p></div>" ).appendTo('.wrapper');
};
