import { useRouter } from "next/router";
import { useEffect } from "react";
import Io from "socket.io-client";

export default function Room() {
  const { query } = useRouter();
  const { username, room } = query;
  const socket = Io();

  useEffect(() => {
    socket.emit("join", { username, room });
    socket.on("message", (data) => {
      console.log(data);
    });

    return () => {
      socket.disconnect();
      socket.off();
    };
  }, []);

  return (
    <>
      salam {username}
      room {room}
    </>
  );
}
