function selectImage(elem) {
	document.getElementById("Popup-Text").innerHTML = elem.getAttribute("title");
	document.getElementById("Popup-Image").src = elem.src;
	document.getElementById("Popup").style.display = "block";
}

function closePopup() {
	document.getElementById("Popup").style.display = "none";
}

document.addEventListener("keyup", (e) => {
	switch (e.key) {
		case "Escape":
			closePopup()
			return
		case "ArrowLeft":
			const prev = document.getElementById("Prev")
			if (prev == null) {
				return
			}
			window.location.replace(prev.href)
			return
		case "ArrowRight":
			window.location.replace(document.getElementById("Next").href) 
			return
		default:
	}
}, false)