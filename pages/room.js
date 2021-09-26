import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import Io from "socket.io-client";
import Head from "next/head";
import { toast } from "react-toastify";

export default function Room() {
  const [chats, setChats] = useState([]);
  const [data, setData] = useState("");
  const { query, push } = useRouter();
  const { username, room } = query;
  const socket = Io();

  function handleSend() {
    if (!data) return null;
    else {
      socket.emit("sendMessage", { message: data, username });
    }
  }

  useEffect(() => {
    socket.emit("join", { username, room });

    socket.on("systemError", (msg) => {
      toast.error(msg, {
        position: "bottom-left",
        closeOnClick: true,
      });
      push("/");
    });

    socket.on("message", (msg) => {
      setChats([...chats, msg]);
      console.log(msg);
    });

    return () => {
      socket.disconnect();
    };
  }, []);

  return (
    <>
      <Head>
        <title>Room {room}</title>
      </Head>
      <div className="w-screen h-screen flex justify-center items-center">
        <div className="w-9/12 min-h-1/2 flex flex-col justify-between items-center bg-white">
          <header className="w-full h-10 bg-blue-500 flex justify-between items-center">
            <span className="ml-5">{username}</span>
            <span className="mr-5 cursor-pointer">close</span>
          </header>
          <div className="w-full min-h-1/2 bg-white"></div>
          <div className="chat-inputs-container w-full h-10 bg-blue-500">
            <input
              type="text"
              placeholder="message..."
              className="chat-input"
              value={data}
              onChange={(e) => setData(e.target.value)}
            />
            <button onClick={handleSend}>Send</button>
          </div>
        </div>
      </div>
    </>
  );
}

export async function getServerSideProps({ query }) {
  const { username, room } = query;

  if (!username || !room) {
    return {
      notFound: true,
    };
  } else {
    return {
      props: {
        auth: true,
      },
    };
  }
}
