import type { Biography as BiographyModel } from '@/models/resume';

export const Biography = ({ biography }: { biography: BiographyModel }) => {
  return (
    <>
      <p>
        {biography.description}
      </p>
      <ul>
        <li>Berlin, Germany</li>
        {biography.links.map((link) => (
          <li key={link.href}>
            <a href={link.href}>{link.label}</a>
          </li>
        ))}
      </ul>
    </>
  );
};
