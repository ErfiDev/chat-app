// const express = require("express");
// const app = express();
// const socket = require("socket.io");

// const io = socket();

// app.use(express.static("public"));

// io.on("connection", (socket) => {
//   console.log("user connected");
// });

// app.listen(3000, () => console.log("listening on *:3000"));

const express = require("express");
const app = express();
const http = require("http");
const server = http.createServer(app);
const { Server } = require("socket.io");
const io = new Server(server);

app.get("/", (req, res) => {
  res.sendFile(__dirname + "/index.html");
});

let users = [];
io.on("connection", (socket) => {
  console.log(socket.id);
  users.push(socket.id);

  socket.on("chat", (data) => {
    console.log(data);
  });

  socket.on("disconnect", () => {
    users.pop();
  });

  socket.emit("user-count", users.length);
});

server.listen(3000, () => {
  console.log("listening on *:3000");
});
