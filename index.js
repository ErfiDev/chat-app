const express = require("express");
const app = express();
const server = require("http").createServer(app);
const { Server } = require("socket.io");
const io = new Server(server);
const cors = require("cors");

const PORT = process.env.PORT || 3000;

app.set("view engine", "ejs");
// app.use(express.static("public"));
app.use(express.json({ limit: "25mb" }));
app.use(express.urlencoded({ extended: true }));
app.use(cors());

app.get("/", (req, res) => {
  res.sendFile(__dirname + "/index.html");
});

io.on("connection", (socket) => {
  console.log(socket.id);
});

app.listen(PORT, () => console.log(`running on ${PORT} port`));
