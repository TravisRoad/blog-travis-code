"use client";

import { useRouter } from "next/navigation";

export default function Logout() {
  const router = useRouter();
  return (
    <>
      <div
        className="cursor-pointer underline"
        onClick={() => {
          fetch("/api/logout", { method: "POST" }).then((res) => {
            if (res.status === 200) {
              router.refresh();
              alert("logout success");
              return;
            }
            alert("logout failed");
          });
        }}
      >
        logout
      </div>
    </>
  );
}
