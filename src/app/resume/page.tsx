import React from 'react';
import { ResumeSection } from './resume-section';
import { ExperienceList } from './experience-list';
import { EducationList } from './education-list';
import type { Resume } from './model';
import { Biography } from './biography';
import { SkillList } from './skill-list';
import { StrengthList } from './strength-list';

const resume: Resume = {
  biography: {
    description: 'Bla bla',
    links: [
      { label: 'GitHub', href: 'https://github.com/untanky' },
      { label: 'LinkedIn', href: 'https://linkedin.com/in/lukasgrimm' },
    ],
  },
  experienceList: [
    {
      startDate: new Date('2024-11-1'),
      company: 'Stealth Startup',
      role: 'Fullstack Developer',
      description: 'Development and maintenance of a social media platform hosted in the cloud.',
    },
    {
      startDate: new Date('2024-10-31'),
      endDate: new Date('2023-9-1'),
      company: 'Empolis Solutions GmbH',
      role: 'Fullstack Engineer',
      description: 'Design, development and maintenance of large scale, on-premise enterprise application written in Java (Spring Boot), Angular and VueJS.',
    },
    {
      startDate: new Date('2024-8-31'),
      endDate: new Date('2021-10-1'),
      company: 'DIPKO GmbH',
      role: 'Fullstack Developer',
      description: 'Design, development and maintenance of cloud platform with Angular and Express. ',
    },
    {
      startDate: new Date('2021-9-30'),
      endDate: new Date('2020-9-1'),
      company: 'Pinuts media+science Multimedia Agentur GmbH',
      role: 'Fullstack Engineer',
      description: 'Development of applications',
    },
  ],
  educationList: [
    {
      startDate: new Date('2017-10-01'),
      endDate: new Date('2021-02-01'),
      school: 'University of Applied Science Berlin (HTW)',
      grade: '1.7',
      course: 'Applied Computer Science',
      description: 'Bla Bla',
    },
  ],
  skills: [
    { label: 'TypeScript', value: 1 },
  ],
  strengths: [
    { label: 'Determination', description: 'Bla bla' },
  ],
};

export default function Resume() {
  return (
    <main className="w-[400px] mx-4 grid grid-cols-1 grid-rows-[auto_min-content_min-content_auto] gap-4 grid-flow-dense mb-8 sm:mx-auto md:grid-cols-3 md:w-[800px]">
      <header className="mt-4 md:col-span-3">
        <h1 className="text-4xl font-bold text-center">
          Lukas Grimm
        </h1>
        <h2 className="text-2xl font-medium text-center">
          Software Engineer
        </h2>
      </header>
      <ResumeSection className="mt-2 md:col-span-1" title="Biography">
        <Biography biography={resume.biography} />
      </ResumeSection>
      <ResumeSection className="mt-2 md:col-span-2 md:col-start-2 md:row-span-3" title="Experience">
        <ExperienceList experienceList={resume.experienceList} />
      </ResumeSection>
      <ResumeSection className="mt-2 md:col-span-2 md:col-start-2" title="Education">
        <EducationList educationList={resume.educationList} />
      </ResumeSection>
      <ResumeSection className="mt-2 md:col-span-2 md:col-start-2" title="Skills">
        <SkillList skills={resume.skills} />
      </ResumeSection>
      <ResumeSection className="mt-2 md:col-span-1" title="Strengths">
        <StrengthList strengths={resume.strengths} />
      </ResumeSection>
    </main>
  );
};
