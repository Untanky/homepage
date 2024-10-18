import { ChevronRightIcon } from '@heroicons/react/16/solid';
import Link from 'next/link';
import type { ReactNode } from 'react';

export type NavigationItem = {
  icon: ReactNode;
  label: string;
  href: string;
};

export const Navigation = ({ navigationItems }: { navigationItems: NavigationItem[] }) => {
  return (
    <nav className="flex items-center space-x-2 mb-2">
      {navigationItems.map((item, index) => (
        <>
          {index !== 0 && <ChevronRightIcon className="size-4 text-zinc-700 dark:text-zinc-300" />}
          <Link
            href={item.href}
            className="flex items-center space-x-2 pl-3 pr-4 py-2 rounded-lg hover:bg-zinc-200 dark:hover:bg-zinc-800"
          >
            {item.icon}
            <span>{item.label}</span>
          </Link>
        </>
      ))}
    </nav>
  );
};
