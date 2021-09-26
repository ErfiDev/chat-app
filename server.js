const app = require("express")();
const server = require("http").Server(app);
const next = require("next");
const dev = process.env.NODE_ENV !== "production";
const nextApp = next({ dev });
const handle = nextApp.getRequestHandler();
const io = require("socket.io")(server);
const { addUser, getUser } = require("./utils/user");

nextApp.prepare().then(() => {
  app.get("*", (req, res) => {
    return handle(req, res);
  });

  io.on("connection", (socket) => {
    socket.on("join", ({ username, room }) => {
      const { error, user } = addUser({ id: socket.id, username, room });

      if (error)
        return socket.emit("systemError", "this username used please change!");
      else {
        socket.emit("message", {
          user: "admin",
          text: `${user.username} welcome to the ${user.room} room`,
        });

        socket.join(room);
        socket.broadcast.to(user.room).emit("message", {
          user: "admin",
          text: `${user.username}, has joined!`,
        });
      }
    });

    socket.on("sendMessage", ({ message, username }) => {
      const { error, user } = getUser(username);

      if (error) return socket.emit("systemError", "can't find this user!");
      else {
        socket.broadcast
          .to(user.room)
          .emit("message", { user: user.username, text: message });
      }
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
