import React, { type ReactNode } from 'react';

type ResumeSectionProps = {
  title: string;
  className?: string;
  children: ReactNode[] | ReactNode;
};

export const ResumeSection = (props: ResumeSectionProps) => {
  return (
    <section className={props.className}>
      <h1 className="text-xl font-medium underline decoration-rose-500">
        {props.title}
      </h1>
      {props.children}
    </section>
  );
};
