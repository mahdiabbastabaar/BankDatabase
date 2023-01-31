const getData = async () => {
    const resp = await fetch('https://api.sampleapis.com/coffee/hot');
    const json = await resp.json();
    console.log(json.values())
    const table = document.getElementById("table");
    for(item in json){
      console.log(item);
      const tr = document.createElement("tr");
      const safebox_id = document.createElement("td");
      const status = document.createElement("td");
      const remove = document.createElement("td");
      safebox_id.innerHTML = json[item].safebox_id;
      status.innerHTML = json[item].status;
      remove.innerHTML = "delete";
      remove.classList.add("text-danger")
      remove.onclick = () => {
        removeData(json[item].id);
      }
      tr.appendChild(safebox_id);
      tr.appendChild(status);
      tr.appendChild(remove);
      table.appendChild(tr);
    }
  }

const removeData = (id) => {
    const resp = fetch(`https://api.sampleapis.com/coffee/hot?id=${id}`);
    resp.then(() => {
      const table = document.getElementById("table");
      table.innerHTML = "";
      getData();
    })
}

window.onload = () => {
  getData();
}