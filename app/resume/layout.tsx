import { SmartNavigation } from '@/components/smart-navigation';
import type { ReactNode } from 'react';

export default function ResumeLayout({ children }: { children: ReactNode[] }) {
  return (
    <div className="w-[400px] px-4 mt-4 mb-8 mx-auto md:w-[800px]">
      <SmartNavigation />
      {children}
    </div>
  );
};
