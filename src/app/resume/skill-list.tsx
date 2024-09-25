import React from 'react';
import type { Skill } from './model';

export const SkillList = ({ skills }: { skills: Skill[] }) => {
  return (
    <ul className="grid grid-cols-2 gap-6">
      {skills.map((skill) => (
        <li key={skill.label}>
          <div className="text-lg font-semibold">
            { skill.label }
            <div className="w-full h-4 bg-rose-100 dark:bg-rose-900 bg-opacity-10 rounded p-px">
              <div style={{ width: `${skill.value * 100}%` }} className="h-full bg-rose-500 rounded"></div>
            </div>
          </div>
        </li>
      ))}
    </ul>
  );
};
