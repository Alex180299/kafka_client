const socket = io();

socket.on('connect', function(){
    console.log("connected")
});

socket.on("/", function (msg) {
    const event = JSON.parse(msg)
    console.log(event)

    $("#nav-tab").append("<button class=\"nav-link active\" id=\"nav-" + event.topic + "\" data-bs-toggle=\"tab\" data-bs-target=\"#nav-" + event.topic + "\" type=\"button\" role=\"tab\" aria-controls=\"nav-" + event.topic + "\" aria-selected=\"true\">" + event.topic +"</button>")
    $("#nav-tabContent").append("<div class=\"tab-pane fade show active\" id=\"nav-" + event.topic + "\" role=\"tabpanel\" aria-labelledby=\"nav-" + event.topic + "\">" + event.value + "</div>")
})
