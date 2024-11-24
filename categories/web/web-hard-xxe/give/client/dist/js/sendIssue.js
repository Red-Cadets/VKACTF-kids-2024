$("#sendIssue").on("click", () => {
  let email = $("#issueEmail").val();
  let message = $("#issueMessage").val();
  var xml = `<?xml version="1.0" encoding="UTF-8"?>
    <Issue>
    <email>${email}</email>
    <message>${message}</message>
    </Issue>
    `;
  let xhr = new XMLHttpRequest();
  let res;
  xhr.open("POST", "/api/issue");
  xhr.onload = () => {
    if (xhr.readyState === xhr.DONE && xhr.status === 200) {
      console.log(xhr.response, xhr.responseXML);
      res = xhr.responseXML;
      alert(res.getElementsByTagName("message")[0].textContent);
      location = "/";
    }else{
      res = xhr.responseXML; 
      if(res){
        alert(res.getElementsByTagName("message")[0].textContent);
      }else{
          alert("Ошибка")
      }
    }
  };
  xhr.send(xml);
});
$("#sendReview").on("click", () => {
  let email = $("#reviewEmail").val();
  let message = $("#reviewMessage").val();
  let rating = $("#reviewRating").val();
  var xml = `<?xml version="1.0" encoding="UTF-8"?>
    <Review>
    <email>${email}</email>
    <message>${message}</message>
    <rating>${parseInt(rating)}</rating>
    </Review>
    `;
  let xhr = new XMLHttpRequest();
  let res;
  xhr.open("POST", "/api/review");
  xhr.onload = () => {
    if (xhr.readyState === xhr.DONE && xhr.status === 200) {
      console.log(xhr.response, xhr.responseXML);
      res = xhr.responseXML;
      alert(res.getElementsByTagName("message")[0].textContent);
      location = "/";
    }else{
      res = xhr.responseXML;
      if(res){
      alert(res.getElementsByTagName("message")[0].textContent);
      }else{
        alert("Ошибка")
      }
    }
  };
  xhr.send(xml);
});
