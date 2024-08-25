"use client";

import { useEffect, useState } from "react";

export default function Home() {
  const [data, setData] = useState();

  const getApi = async () => {
    try {
      const response = await fetch("http://localhost:8080/hoge");
      const a = await response.json();
      setData(a);
    } catch (error) {
      console.log(error);
    }
  };

  console.log(data);

  useEffect(() => {
    getApi();
  }, []);

  return (
    <>
      <div>
        <h1>Hello World!!</h1>
      </div>
    </>
  );
}
