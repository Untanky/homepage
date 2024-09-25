import React from 'react';
import type { Education } from './model';

const formatDate = (date: Date) => {
  return date.toLocaleDateString('en-US', { day: undefined, month: 'long', year: 'numeric' });
};

export const EducationList = ({ educationList }: { educationList: Education[] }) => {
  return (
    <ul className="space-y-3">
      {educationList.map((education) => (
        <li key={Math.random()}>
          <h4 className="text-lg font-medium">
            {education.course}
            &nbsp;-&nbsp;
            <span className="italic">{education.school}</span>
            (
            {education.grade}
            )
          </h4>
          <p>
            {formatDate(education.startDate)}
            &nbsp;-&nbsp;
            {formatDate(education.endDate)}
          </p>
          <p className="text-zinc-700 dark:text-zinc-300">
            {education.description}
          </p>
        </li>
      ))}
    </ul>
  );
};
