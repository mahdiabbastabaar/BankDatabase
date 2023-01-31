const getData = async () => {
    const resp = await fetch('https://api.sampleapis.com/coffee/hot');
    const json = await resp.json();
    console.log(json.values())
    const table = document.getElementById("table");
    for(item in json){
      console.log(item);
      const tr = document.createElement("tr");
      const safebox_id = document.createElement("td");
      const customer_id = document.createElement("td");
      const remaining_time = document.createElement("td");
      const payment_conrtact = document.createElement("td");
      safebox_id.innerHTML = json[item].safebox_id;
      customer_id.innerHTML = json[item].customer_id;
      remaining_time.innerHTML = json[item].remaining_time;
      payment_conrtact.innerHTML = json[item].payment_conrtact;
      tr.appendChild(safebox_id);
      tr.appendChild(customer_id);
      tr.appendChild(remaining_time);
      tr.appendChild(payment_conrtact);
      table.appendChild(tr);
    }
  }
window.onload = () => {
  getData();
}