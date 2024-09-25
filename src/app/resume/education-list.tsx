import React from 'react';
import type { Education } from './model';

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
            {education.startDate.toDateString()}
            &nbsp;-&nbsp;
            {education.endDate ? education.endDate.toDateString() : 'now'}
          </p>
          <p>
            {education.description}
          </p>
        </li>
      ))}
    </ul>
  );
};
