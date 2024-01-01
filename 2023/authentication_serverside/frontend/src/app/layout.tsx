import "./globals.css";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className="flex min-h-screen flex-col items-center p-24">
      <body>{children}</body>
    </html>
  );
}
