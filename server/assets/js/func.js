function addSource() {
    var url = $("input#source-url");
    if (url.length == 0) {
        console.log("couldnt find input#source-url");
        return;
    }
    url = url[0].value;
    console.log("url: ", url);
    $.post ("/api", {action: "addsource", url: url }, function (data) {
        // { "status_code": -1, error_msg: "kjds"}
        console.log("addsource-received-result: ", data);
        if (data.StatusCode > 0) {
            // display error_msg
            var error_msg = $("p#addsource-alert-msg");
            error_msg.text(data.ErrorMsg);
            $("#addsource-alert-error").alert()
        } else {
            // display success msg
            $("#addsource-alert-success").show();
            $("#addsource-alert-success").alert();
        }
    }, "json");
}