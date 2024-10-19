import { loadJsonFile } from '@/lib/load-file';
import type { Resume } from '@/models/resume';
import { Metadata } from 'next';
import Image from 'next/image';
import { Biography } from './biography';
import { EducationList } from './education-list';
import { ExperienceList } from './experience-list';
import { ResumeSection } from './resume-section';
import { SkillList } from './skill-list';
import { StrengthList } from './strength-list';

export const metadata: Metadata = {
  title: 'Resumé Lukas Grimm',
};

export default async function Resume() {
  const resume = await loadJsonFile<Resume>('./data/resume.json').then((resume) => ({
    ...resume,
    experienceList: resume.experienceList.map((experience) => ({
      ...experience,
      startDate: new Date(experience.startDate),
      endDate: experience.endDate && new Date(experience.endDate),
    })),
    educationList: resume.educationList.map((education) => ({
      ...education,
      startDate: new Date(education.startDate),
      endDate: new Date(education.endDate),
    })),
  }));

  return (
    <main className="grid grid-cols-1 grid-rows-[auto_min-content_min-content_auto] gap-4 grid-flow-dense md:grid-cols-3">
      <header className="relative self-center md:col-span-3 py-4 border border-rose-500 rounded-2xl mt-[140px] md:mt-0">
        <h1 className="text-4xl font-bold text-center">
          Lukas Grimm
        </h1>
        <h2 className="text-2xl font-medium text-center">
          Software Engineer
        </h2>
        <Image
          src="/profile.jpg"
          alt="Profile Picture"
          width={150}
          height={150}
          priority={true}
          placeholder="empty"
          className="absolute top-[-140px] right-[109px] md:right-[-1px] md:top-[-23px] mx-auto rounded-2xl"
        />
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
