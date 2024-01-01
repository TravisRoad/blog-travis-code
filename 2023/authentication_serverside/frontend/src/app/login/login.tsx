"use client";

import { useRouter } from "next/navigation";
import { useState } from "react";

interface User {
  username: string;
  password: string;
}

export default function Login() {
  const [cred, setCred] = useState<User>({
    username: "",
    password: "",
  });
  const router = useRouter();

  return (
    <>
      <div>
        <label htmlFor="username" className="w-1/6">
          Username
        </label>
        <input
          type="text"
          onChange={(e) => setCred({ ...cred, username: e.target.value })}
          className="border p-2 w-[20rem] ml-2"
        />
      </div>
      <div>
        <label htmlFor="password" className="w-1/6">
          Password
        </label>
        <input
          type="password"
          onChange={(e) => setCred({ ...cred, password: e.target.value })}
          className="border p-2 w-[20rem] ml-2"
        />
      </div>
      <button
        className="border p-2 hover:bg-gray-100"
        onClick={() => {
          fetch("/api/login", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(cred),
          }).then((res) => {
            if (res.status === 200) {
              alert("Login success");
              router.push("/");
              return;
            }

            alert("Login failed");
          });
        }}
      >
        Login
      </button>
    </>
  );
}
