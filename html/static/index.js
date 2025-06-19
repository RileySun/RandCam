let state = "main";
let popupCount = 0;

//Popup
function selectImage(elem) {
	document.getElementById("Popup-Text").innerHTML = elem.getAttribute("title");
	document.getElementById("Popup-Image").src = elem.src;
	document.getElementById("Popup").style.display = "block";
	state = "popup"
}

function closePopup() {
	document.getElementById("Popup").style.display = "none";
	state = "main"
}

//Controls
function controls(e) {
	switch (state) {
		case "main":
			mainControls(e);
		case "popup":
			popupControls(e);
		default:
	}
}

function mainControls(e) {
	switch (e.key) {
		case "ArrowLeft":
			const prev = document.getElementById("Prev");
			if (prev == null) {
				return;
			}
			window.location.replace(prev.href);
			return;
		case "ArrowRight":
			window.location.replace(document.getElementById("Next").href);
			return;
		default:
	}
}

function popupControls(e) {
	const cameras = document.getElementsByClassName("Camera-Src")
	const cams = Array.from(cameras)

	switch (e.key) {
		case "Escape":
			closePopup();
			return
		case "ArrowLeft":
			if (popupCount > 0) {
				popupCount -= 1
			} else {
				popupCount = cameras.length
			}
			document.getElementById("Popup-Text").innerHTML = cameras[popupCount].getAttribute("title");
			document.getElementById("Popup-Image").src = cameras[popupCount].src;
			return;
		case "ArrowRight":
			if (popupCount < cameras.length) {
				popupCount += 1
			} else {
				popupCount = 0
			}
			document.getElementById("Popup-Text").innerHTML = cameras[popupCount].getAttribute("title");
			document.getElementById("Popup-Image").src = cameras[popupCount].src;
			return;
		default:
	}
}

document.addEventListener("keyup", (e) => {controls(e)}, false)