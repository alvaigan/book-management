import Sidebar from "@/components/Sidebar";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
        <Sidebar>
          <div className="m-5">{children}</div>
        </Sidebar>
  );
}
