let message  = "Hi";
let no_ofmessages = 5;
let i = 0;
let delay = setInterval(function(){
	document.getElementsByClassName('composer_rich_textarea')[0].innerText = message;		
	$(".im_submit_send").trigger("mousedown");	
	i++;
	if(i==no_ofmessages){
		clearInterval(delay)
	}
	
}, 1000);