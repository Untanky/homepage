import React from 'react';
import type { Experience } from './model';

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
            {experience.startDate.toDateString()}
            &nbsp;-&nbsp;
            {experience.endDate ? experience.endDate.toDateString() : 'now'}
          </p>
          <p>
            {experience.description}
          </p>
        </li>
      ))}
    </ul>
  );
};
