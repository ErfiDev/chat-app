import { useRouter } from "next/router";

export default function Room() {
  const { query } = useRouter();
  const { username, room } = query;

  return (
    <>
      salam {username}
      room {room}
    </>
  );
}
