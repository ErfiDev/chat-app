import Head from "next/head";
import Form from "comp/form";

export default function Home() {
  function handleSubmit(e) {
    e.preventDefault();

    console.log("submitted");
  }

  return (
    <>
      <Head>
        <title>joining to room</title>
      </Head>
      <Form handleSubmit={handleSubmit} />
    </>
  );
}
