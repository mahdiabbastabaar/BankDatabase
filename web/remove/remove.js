const buttonClicked = async (event) => {
    const safeBoxIdElement = document.getElementById("safebox-id");
    const safeBoxId = safeBoxIdElement.value;
    const body = {id: safeBoxId};
    const response = await fetch("https://httpbin.org/status/400", {
        method: "POST",
        body: JSON.stringify(body)
    });
    safeBoxIdElement.value = "";
    if (response.ok) {
        alert("evacuated successfully!");
    } else {
        alert(`an error occurred:\n${response.status}: ${response.statusText}`);
    }
}

const evacuationButton = document.getElementById("evacuation-button");
evacuationButton.onclick = async (event) => {
    await buttonClicked(event);
}