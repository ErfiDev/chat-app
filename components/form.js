import Link from "next/link";
import { useState } from "react";

export default function Form({ handleSubmit }) {
  const [username, setUsername] = useState("");
  const [room, setRoom] = useState("");

  return (
    <div className="w-screen h-screen flex justify-center items-center">
      <form
        onSubmit={(e) => handleSubmit(e)}
        className="form-get-room w-3/4 flex flex-col justify-around items-center p-20"
      >
        <div className="md:flex md:items-center mb-6">
          <div className="md:w-1/3">
            <label
              className="block text-white font-bold md:text-right mb-1 md:mb-0 pr-4"
              htmlFor="username"
            >
              Username
            </label>
          </div>
          <div className="md:w-2/3">
            <input
              className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
              id="username"
              type="text"
              placeholder="Jane Doe"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
        </div>
        <div className="md:flex md:items-center mb-6">
          <div className="md:w-1/3">
            <label
              className="block text-white font-bold md:text-right mb-1 md:mb-0 pr-4"
              htmlFor="room"
            >
              Room
            </label>
          </div>
          <div className="md:w-2/3">
            <input
              className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
              id="room"
              type="text"
              placeholder="room"
              value={room}
              onChange={(e) => setRoom(e.target.value)}
            />
          </div>
        </div>
        <div className="flex items-center">
          <div>
            <button
              className="shadow bg-purple-500 hover:bg-purple-400 focus:shadow-outline focus:outline-none text-white font-bold py-2 px-4 rounded"
              type="button"
            >
              <Link href={`/room?username=${username}&room=${room}`}>
                Sign In
              </Link>
            </button>
          </div>
        </div>
      </form>
    </div>
  );
}
