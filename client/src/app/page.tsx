"use client";

import { useEffect, useState } from "react";

type User = {
  id: number;
  name: string;
};

export default function Home() {
  const [users, setUsers] = useState<User[]>([]);

  const getApi = async () => {
    try {
      const response = await fetch("http://localhost:8080/hoge");
      const data = await response.json();
      setUsers(data);
    } catch (error) {
      console.log(error);
    }
  };

  console.log(users);

  useEffect(() => {
    getApi();
  }, []);

  return (
    <>
      <div>
        <h1>Hello World!!</h1>
        {users.map((user) => {
          return <div key={user.id}>{user.name}</div>;
        })}
      </div>
    </>
  );
}
