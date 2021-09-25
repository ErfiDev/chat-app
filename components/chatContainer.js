export default function ChatContainer({ username, chats }) {
  return (
    <div className="w-screen h-screen flex justify-center items-center">
      <div className="w-9/12 min-h-1/2 flex flex-col justify-between items-center bg-white">
        <header className="w-full h-10 bg-blue-500 flex justify-between items-center">
          <span className="ml-5">{username}</span>
          <span className="mr-5 cursor-pointer">close</span>
        </header>
        <div className="w-full min-h-1/2 bg-white">
          {/* {chats.map((item) => {
            <span className="block">{item.text}</span>;
          })} */}
        </div>
        <div className="chat-inputs-container w-full h-10 bg-blue-500">
          <input type="text" placeholder="message..." className="chat-input" />
          <button>Send</button>
        </div>
      </div>
    </div>
  );
}
