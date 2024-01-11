<template>
    <main class="max-w-[820px] mx-4 md:mx-auto space-y-4">
        <ResumeHeader name="Lukas Grimm" role="Software Engineer from Berlin, Germany" />
        <ResumeSection title="Experience">
            <ResumeExperienceList :experiences="list" />
        </ResumeSection>
        <ResumeSection title="Education">
            <ResumeEducationList :educations="edulist" />
        </ResumeSection>
        <ResumeSection title="Skills" />
        <ResumeSection title="Tech Stack" />
        <ResumeSection title="Strengths" />
        <ResumeSection title="Biography" />
    </main>
</template>

<script setup lang="ts">
    import type { EducationItemModel, ExperienceItemModel } from '~/components/resume/models';

    const list = [
        { start: new Date(2023, 8), company: 'Some Company Ltd', role: 'Software Engineer', description: 'Lorem Ipsum' },
        { start: new Date(2023, 7), end: new Date(2021, 9), company: 'Some Company GmbH', role: 'Software Engineer', description: 'Lorem Ipsum' },
    ] satisfies ExperienceItemModel[]

    const edulist = [
        { start: new Date(2023, 8), school: 'Some Company Ltd', course: 'Software Engineer', grade: '', description: 'Lorem Ipsum' },
        { start: new Date(2023, 7), end: new Date(2021, 9), school: 'Some Company GmbH', course: 'Software Engineer', grade: '', description: 'Lorem Ipsum' },
    ] satisfies EducationItemModel[]
</script>
