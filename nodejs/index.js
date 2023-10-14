var server = require("ws").Server;
var s = new server({port:5050});
s.on("connection", (ws) => {
    ws.on("message", (msg) => {
        console.log("Recv: " + msg);
        s.clients.forEach((client) => {
            console.log("Send: " + msg);
            client.send(msg.toString());
        });
    });
    ws.on("close", () => {
        console.log("Closed.");
    });
});