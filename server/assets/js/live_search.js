function showSearch(e) {
    e.length > 0 && ($(".livesearchresult").show(), $(".livesearchresult").html('<span class="saving">Loading<span>.</span><span>.</span><span>.</span></span>'), $.ajax({
        type: "POST",
        url: "/nu",
        data : {
            searchTerm: e
        },
        success: function(e) {
            $(".livesearchresult").html(e)
        }
    }))
}

function search_btn_click() {
    console.log("search_btn_click")
}