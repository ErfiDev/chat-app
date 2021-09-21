import io from "socket.io-client";
import { useEffect, useState } from "react";

export default function About() {
  const socket = io();
  const [data, setData] = useState();

  function toServer() {
    socket.emit("file-stream", data);
    console.log("to server working");
  }

  return (
    <>
      <input type="file" onChange={(e) => setData(e.target.value)} />
      <button onClick={toServer}>Click</button>
    </>
  );
}
