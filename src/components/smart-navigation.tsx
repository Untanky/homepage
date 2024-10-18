'use client';

import { CodeBracketIcon, HomeIcon } from '@heroicons/react/20/solid';
import { usePathname } from 'next/navigation';
import { Navigation, type NavigationItem } from './navigation';

const homeNavigationItem: NavigationItem = {
  icon: <HomeIcon className="size-5" />,
  href: '/',
  label: 'Home',
};

const useNavigationItems = (): NavigationItem[] => {
  const pathname = usePathname();

  if (pathname.startsWith('/projects/')) {
    const projectNavigationItem: NavigationItem = {
      icon: <CodeBracketIcon className="size-5" />,
      href: '/projects',
      label: 'Projects',
    };
    return [homeNavigationItem, projectNavigationItem];
  }

  if (pathname.startsWith('/projects')) {
    return [homeNavigationItem];
  }

  return [];
};

export const SmartNavigation = () => {
  const navigationItems = useNavigationItems();

  if (navigationItems.length === 0) {
    return null;
  }

  return (
    <Navigation navigationItems={navigationItems} />
  );
};
