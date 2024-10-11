import type { Experience } from '@/models/resume';

const formatDate = (date: Date) => {
  return date.toLocaleDateString('en-US', { day: undefined, month: 'long', year: 'numeric' });
};

export const ExperienceList = ({ experienceList }: { experienceList: Experience[] }) => {
  return (
    <ul className="space-y-3">
      {experienceList.map((experience) => (
        <li key={Math.random()}>
          <h4 className="text-lg font-medium">
            {experience.role}
            &nbsp;-&nbsp;
            <span className="italic">{experience.company}</span>
          </h4>
          <p>
            {formatDate(experience.startDate)}
            &nbsp;-&nbsp;
            {experience.endDate ? formatDate(experience.endDate) : 'now'}
          </p>
          <p className="text-zinc-700 dark:text-zinc-300">
            {experience.description}
          </p>
        </li>
      ))}
    </ul>
  );
};
