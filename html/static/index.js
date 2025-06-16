function selectImage(elem) {
	document.getElementById("Popup-Text").innerHTML = elem.getAttribute("title");
	document.getElementById("Popup-Image").src = elem.src;
	document.getElementById("Popup").style.display = "block";
}

function closePopup() {
	document.getElementById("Popup").style.display = "none";
}