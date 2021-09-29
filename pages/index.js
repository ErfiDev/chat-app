import Head from "next/head";
import Form from "comp/form";
import Io from "socket.io-client";
import { useEffect } from "react";
import getCookie from "utils/cookie";

export default function Home() {
  const socket = Io();

  useEffect(() => {
    let usernameCookie = getCookie("username");

    if (!usernameCookie) return null;
    else {
      socket.emit("dis", usernameCookie);
    }
  }, []);

  return (
    <>
      <Head>
        <title>Joining to room</title>
      </Head>
      <Form />
    </>
  );
}
