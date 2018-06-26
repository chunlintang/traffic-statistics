/**
 * before include jquery.js
 * in this project we include jquery.3.3.1.min.js /static/jquery.3.3.1.min.js
 */
$(document).ready(function () {
    // localhost port:8888
    $.get("http://dev.dig.com/dig", {
        "time": gettime(),
        "url": geturl(),
        "refer": getrefer(),
        "ua": getuser_agent()
    })
});

function gettime() {
    var nowDate = new Date();
    return nowDate.toLocaleString();
}

function geturl() {
    return window.location.href;
}

function getrefer() {
    return document.referrer;
}

function getuser_agent() {
    return navigator.userAgent;
}