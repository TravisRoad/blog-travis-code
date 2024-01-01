"use Server";

import { cookies } from "next/headers";

export async function isLogin(): Promise<boolean> {
  console.log("check login");
  console.log(cookies().toString());

  const hasLogin = await fetch("http://localhost:8080/api/islogin", {
    cache: "no-store",
    headers: {
      Cookie: cookies().toString(),
    },
  }).then((res) => {
    return res.status === 200;
  });
  return hasLogin;
}
