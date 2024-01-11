<template>
    <main class="max-w-[820px] mx-4 md:mx-auto space-y-4">
        <ResumeHeader name="Lukas Grimm" role="Software Engineer from Berlin, Germany" />
        <div class="flex flex-col md:flex-row space-y-4 md:space-x-6 md:space-y-0">
            <div class="flex flex-col md:min-w-[240px] md:w-1/3 md:max-w-[360px] space-y-4">
                <ResumeSection title="Biography">
                    <ResumeBiography :biography="biography" />
                </ResumeSection>
                <ResumeSection title="Strengths">
                    <ul class="flex flex-col space-y-2">
                        <li v-for="strength in strengths">
                            <ResumeStrengthItem :strength="strength" />
                        </li>
                    </ul>
                </ResumeSection>
            </div>
            <div class="flex flex-col space-y-4">
                <ResumeSection title="Experience">
                    <ResumeExperienceList :experiences="list" />
                </ResumeSection>
                <ResumeSection title="Education">
                    <ResumeEducationList :educations="edulist" />
                </ResumeSection>
                <ResumeSection title="Skills">
                    <ul class="grid grid-cols-3 gap-4">
                        <li v-for="skill in skills">
                            <ResumeSkillItem :skill="skill" />
                        </li>
                    </ul>
                </ResumeSection>
            </div>
        </div>
    </main>
</template>

<script setup lang="ts">
    import type { BiographyModel, EducationItemModel, ExperienceItemModel, SkillItemModel, StrengthItemModel } from '~/components/resume/models';

    const list = [
        { start: new Date(2023, 8), company: 'Some Company Ltd', role: 'Software Engineer', description: 'Lorem Ipsum' },
        { start: new Date(2023, 7), end: new Date(2021, 9), company: 'Some Company GmbH', role: 'Software Engineer', description: 'Lorem Ipsum' },
    ] satisfies ExperienceItemModel[];

    const edulist = [
        { start: new Date(2023, 8), school: 'Some Company Ltd', course: 'Software Engineer', grade: '', description: 'Lorem Ipsum' },
        { start: new Date(2023, 7), end: new Date(2021, 9), school: 'Some Company GmbH', course: 'Software Engineer', grade: '', description: 'Lorem Ipsum' },
    ] satisfies EducationItemModel[];

    const skills = [
        { name: 'Go', value: 1 },
        { name: 'TypeScript', value: 1 },
        { name: 'CI/CD', value: 0.9 },
    ] satisfies SkillItemModel[];

    const strengths = [
        { name: 'Go', description: 'Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum' },
        { name: 'TypeScript', description: 'Lorem Ipsum' },
        { name: 'CI/CD', description: 'Lorem Ipsum' },
    ] satisfies StrengthItemModel[];

    const biography = {
        about: 'Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum',
        location: 'Berlin, Germany',
        github: 'https://github.com/Untanky',
        linkedIn: 'https://linkedin.com/in/lukasgrimm',
        x: 'https://twitter.com/Untanky'
    } satisfies BiographyModel;
</script>
