import { GitHubIcon } from '@/components/github-icon';
import { LinkedInIcon } from '@/components/linkedin-icon';
import type { Biography as BiographyModel } from '@/models/resume';
import { MapPinIcon } from '@heroicons/react/16/solid';
import { type ReactNode, useMemo } from 'react';

export const Biography = ({ biography }: { biography: BiographyModel }) => {
  const links = useMemo(() => biography.links.map((link) => {
    let icon: ReactNode = null;
    if (link.href.startsWith('https://github.com')) {
      icon = <GitHubIcon className="size-4" />;
    } else if (link.href.startsWith('https://linkedin.com')) {
      icon = <LinkedInIcon className="size-4" />;
    } else {
      console.warn(`unknown link ${link}`);
    }

    return {
      ...link,
      icon,
    };
  }), [biography]);

  return (
    <>
      <p>
        {biography.description}
      </p>
      <ul>
        <li className="flex items-center space-x-2">
          <MapPinIcon className="size-4" />
          <span>Berlin, Germany</span>
        </li>
        {links.map((link) => (
          <li key={link.href}>
            <a href={link.href} className="flex items-center space-x-2">
              {link.icon}
              <span>{link.label}</span>
            </a>
          </li>
        ))}
      </ul>
    </>
  );
};
