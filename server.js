const app = require("express")();
const server = require("http").Server(app);
const next = require("next");
const dev = process.env.NODE_ENV !== "production";
const nextApp = next({ dev });
const handle = nextApp.getRequestHandler();
const io = require("socket.io")(server);
const { addUser, getUser } = require("utils/user");

nextApp.prepare().then(() => {
  app.get("*", (req, res) => {
    return handle(req, res);
  });

  io.on("connect", (socket) => {
    socket.on("join", ({ username, room }, cb) => {
      const { error, user } = addUser(socket.id, username, room);

      if (error) return cb(error);

      socket.emit("message", {
        user: "admin",
        text: `${user.username} welcome to the ${user.room} room`,
      });

      socket.broadcast
        .to(user.room)
        .emit("message", {
          user: "admin",
          text: `${user.username}, has joined!`,
        });
      socket.join(user.room);
    });

    socket.on("sendMessage", (message, cb) => {
      const user = getUser(socket.id);

      io.to(user.room).emit("message", { user: user.username, text: message });

      cb();
    });

    socket.on("disconnect", () => {
      console.log("user disconnected");
    });
  });

  server.listen(3000, (err) => {
    if (err) throw err;
    console.log("on 3000");
  });
});
