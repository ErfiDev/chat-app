const app = require("express")();
const server = require("http").Server(app);
const next = require("next");
const dev = process.env.NODE_ENV !== "production";
const nextApp = next({ dev });
const handle = nextApp.getRequestHandler();
const io = require("socket.io")(server);

nextApp.prepare().then(() => {
  app.get("*", (req, res) => {
    return handle(req, res);
  });

  io.on("connect", (socket) => {
    socket.emit("users", `${socket.id} is connected`);

    const chats = [];
    socket.on("chat", (data) => {
      chats.push(data);
    });

    socket.emit("all-chats", chats);
  });

  server.listen(3000, (err) => {
    if (err) throw err;
    console.log("on 3000");
  });
});
