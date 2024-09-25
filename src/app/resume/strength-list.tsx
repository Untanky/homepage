import React from 'react';
import type { Strength } from './model';

export const StrengthList = ({ strengths }: { strengths: Strength[] }) => {
  return (
    <ul className="flex flex-col space-y-2">
      {strengths.map((strength) => (
        <li key={strength.label}>
          <span className="text-lg font-medium mr-2">
            { strength.label }
          </span>
          <span className="text-zinc-700 dark:text-zinc-300">
            { strength.description }
          </span>
        </li>
      ))}
    </ul>
  );
};
