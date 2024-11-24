let xhr = new XMLHttpRequest();
xhr.open("GET", "/api/review", true);
xhr.onload = function () {
  var xml = xhr.responseXML;
  var reviews = xml.getElementsByTagName("Reviews");
  var arrayNodes = [...reviews[0].children];
  arrayNodes.reverse();
  for (let review of arrayNodes) {
    var message = review.getElementsByTagName("message")[0].textContent;
    var email = review.getElementsByTagName("email")[0].textContent;
    var rating = review.getElementsByTagName("rating")[0].textContent;
    $("#comments").append(
      DOMPurify.sanitize(`
            <div class="p-3 border border-top-0 comment mb-0">
                <div class="d-flex justify-content-between border-bottom">
                    <div class=""><b>${email}</b></div> <div class="">
                        <div class="list-unstyled d-flex justify-content-center mb-0">
                            <i class="fa fa-star${
                              rating > 0 ? "" : "-o"
                            }" aria-hidden="true"></i>
                            <i class="fa fa-star${
                              rating > 1 ? "" : "-o"
                            }" aria-hidden="true"></i>
                            <i class="fa fa-star${
                              rating > 2 ? "" : "-o"
                            }" aria-hidden="true"></i>
                            <i class="fa fa-star${
                              rating > 3 ? "" : "-o"
                            }" aria-hidden="true"></i>
                            <i class="fa fa-star${
                              rating > 4 ? "" : "-o"
                            }" aria-hidden="true"></i>

                        </div>
                    </div>
                </div>
                <div class="row"><p>${message}</p></div>
                
            </div>
        `)
    );
  }
};
xhr.send();
