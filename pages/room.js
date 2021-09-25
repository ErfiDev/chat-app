import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import Io from "socket.io-client";
import ChatContainer from "comp/chatContainer";
import Head from "next/head";

export default function Room() {
  const [chats, setChats] = useState([]);
  const { query } = useRouter();
  const { username, room } = query;
  const socket = Io();

  useEffect(() => {
    socket.emit("join", { username, room });
    socket.on("message", (data) => {
      setChats(data);
    });

    return () => {
      socket.disconnect();
      socket.off();
    };
  }, []);

  return (
    <>
      <Head>
        <title>Room {room}</title>
      </Head>
      <ChatContainer username={username} chats={chats} />
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
