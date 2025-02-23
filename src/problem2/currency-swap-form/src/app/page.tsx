import SwapForm from "@/components/form/swap-form";
import Image from "next/image";

export default function Home() {
  return (
    <div className="flex flex-col items-center h-screen w-screen justify-center min-h-screen p-5 font-[family-name:var(--font-geist-sans)] bg-blue-400">
      <main className="sm:w-full md:w-4/5 flex flex-col justify-center items-center">
        <SwapForm />
      </main>
    </div>
  );
}
