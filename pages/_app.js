import "../styles/App.scss";
import "tailwindcss/tailwind.css";
import "react-toastify/dist/ReactToastify.css";
import { ToastContainer } from "react-toastify";
import { useEffect } from "react";
import getCookie from "utils/cookie";
import Io from "socket.io-client";

function MyApp({ Component, pageProps }) {
  const socket = Io();

  useEffect(() => {
    return () => {
      let usernameCookie = getCookie("username");

      if (!usernameCookie) return null;
      else {
        socket.emit("dis", usernameCookie);
        document.cookie = `username= ;`;
      }
    };
  }, []);

  return (
    <>
      <Component {...pageProps} />
      <ToastContainer />
    </>
  );
}

export default MyApp;
