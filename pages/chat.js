import { useEffect, useState } from "react";
import Io from "socket.io-client";

export default function Chat() {
  const socket = Io();
  const [inputData, setInputData] = useState("");
  const [all, setAllChats] = useState([]);

  function sendToServer() {
    socket.emit("chat", inputData);

    socket.on("all-chats", (data) => {
      setAllChats(data);
      console.log(data);
    });
  }

  return (
    <>
      <ul>
        {all.map((item) => {
          <li>{item}</li>;
        })}
      </ul>

      <input
        type="text"
        id="text"
        value={inputData}
        onChange={(e) => setInputData(e.target.value)}
      />
      <button onClick={sendToServer}>send</button>
    </>
  );
}
