import { SmartNavigation } from '@/components/smart-navigation';
import type { ReactNode } from 'react';

export default function ResumeLayout({ children }: { children: ReactNode[] }) {
  return (
    <div className="w-[400px] m-4 mb-8 sm:mx-auto md:w-[800px] ">
      <SmartNavigation />
      {children}
    </div>
  );
};
