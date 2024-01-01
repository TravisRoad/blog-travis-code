import { isLogin } from "@/lib/auth";
import { redirect } from "next/navigation";
import Logout from "./logout";

export default async function Home() {
  await isLogin().then((res: boolean) => {
    if (!res) {
      redirect("/login");
    }
  });

  return (
    <main >
      <>
        You has logged in
        <div>
          <Logout />
        </div>
      </>
    </main>
  );
}
