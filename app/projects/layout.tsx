import { Navigation } from '@/components/navigation';

export default function ProjectLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="m-4 md:w-[700px] md:mx-auto">
      <Navigation />
      {children}
    </div>
  );
}
