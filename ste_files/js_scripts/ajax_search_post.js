function showResult(e, a, s) {
    e.length >= 2 && ("series" == s && ($(".livesearch." + a).show(), $(".livesearch." + a).html('<span style="padding:5px;">Loading...</span>')), "group" == s && ($(".livesearchgroup." + a).show(), $(".livesearchgroup." + a).html('<span style="padding:5px;">Loading...</span>')), "addrelease" == s && ($(".livesearch").show(), $(".livesearch").html('<span style="padding:5px;">Loading...</span>')), $.ajax({
        type: "POST",
        url: "https://www.novelupdates.com/wp-admin/admin-ajax.php",
        data: {
            action: "nd_ajaxsearch",
            str: e,
            strID: a,
            strType: s
        },
        success: function (e) {
            e = e.slice(0, -1), $(".livesearch").hide(), $(".livesearchgroup").hide(), "series" == s && ($(".livesearch." + a).show(), $(".livesearch." + a).html(e)), "group" == s && ($(".livesearchgroup." + a).show(), $(".livesearchgroup." + a).html(e)), "addrelease" == s && ($(".livesearch").show(), $(".livesearch").html(e))
        }
    }))
}

function changeitem(e, a, s, i) {
    "title" == s && ($("#title" + e).val(a), $("#title_change_" + e).val($(i).text()), $(".livesearch").hide(), $(".livesearchgroup").hide()), "group" == s && ($("#group" + e).val(a), $("#group_change_" + e).val($(i).text()), $(".livesearch").hide(), $(".livesearchgroup").hide()), "addrelease" == s && ($("#artitle").val(a), $("#title_change_").val($(i).text()), $(".livesearch").hide(), $(".livesearchgroup").hide())
}

function showSearch(e) {
    e.length > 0 && ($(".livesearchresult").show(), $(".livesearchresult").html('<span class="saving">Loading<span>.</span><span>.</span><span>.</span></span>'), $.ajax({
        type: "POST",
        url: "https://www.novelupdates.com/wp-admin/admin-ajax.php",
        data: {
            action: "nd_ajaxsearchmain",
            strType: "desktop",
            strOne: e
        },
        success: function (e) {
            e = e.slice(0, -1), $(".livesearchresult").html(e)
        }
    }))
}

function showSearch_mb(e) {
    e.length > 0 && ($(".livesearch_mb").show(), $(".livesearch_mb").html('<span class="saving" style="font-size: 15px;">Loading<span>.</span><span>.</span><span>.</span></span>'), $.ajax({
        type: "POST",
        url: "https://www.novelupdates.com/wp-admin/admin-ajax.php",
        data: {
            action: "nd_ajaxsearchmain",
            strType: "mobile",
            strOne: e
        },
        success: function (e) {
            e = e.slice(0, -1), $(".livesearch_mb").html(e)
        }
    }))
}
$(document).ready(function () {
    $("body").click(function (e) {
        var a = $(".lo_main").css("display");
        ("user_avatar" != e.target.id && "menu_right_item" != e.target.id && "user_menu_links" != e.target.id && "fa-fontawesome" != e.target.id && "block" == a && toggleUserMenu(), "searchme-rl" == e.target.id) ? $("#searchme-rl").val().length > 1 && $(".livesearchresult").show(): ($(".livesearchresult").hide(), $(".livesearch").hide(), $(".livesearchgroup").hide())
    })
});