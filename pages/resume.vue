<template>
    <main class="max-w-[820px] mx-4 md:mx-auto space-y-4">
        <ResumeHeader name="Lukas Grimm" role="Software Engineer from Berlin, Germany" />
        <div class="flex flex-col md:flex-row space-y-4 md:space-x-6 md:space-y-0">
            <div class="flex flex-col md:min-w-[240px] md:w-1/3 md:max-w-[360px] space-y-4">
                <ResumeSection title="Biography">
                    <ResumeBiography :biography="biography!" />
                </ResumeSection>
                <ResumeSection title="Strengths">
                    <ul class="flex flex-col space-y-2">
                        <li v-for="strength in resume?.strengths">
                            <ResumeStrengthItem :strength="strength" />
                        </li>
                    </ul>
                </ResumeSection>
            </div>
            <div class="flex flex-col space-y-4">
                <ResumeSection title="Experience">
                    <ResumeExperienceList :experiences="resume?.experience || []" />
                </ResumeSection>
                <ResumeSection title="Education">
                    <ResumeEducationList :educations="resume?.education || []" />
                </ResumeSection>
                <ResumeSection title="Skills">
                    <ul class="grid grid-cols-3 gap-4">
                        <li v-for="skill in resume?.skills">
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

    type ResumeModel = {
        experience: ExperienceItemModel[];
        education: EducationItemModel[];
        skills: SkillItemModel[];
        strengths: StrengthItemModel[];
    }

    const { data: resume } = await useAsyncData('resume', () => queryContent<ResumeModel>('resume').findOne().then((data) => ({
        ...data,
    })));
    
    if (resume.value) {
        resume.value.experience = resume.value?.experience?.map((experience) => ({
            ...experience,
            start: new Date(experience.start),
            end: experience.end ? new Date(experience.end) : undefined
        })) || [];

        resume.value.education = resume.value?.education?.map((education) => ({
            ...education,
            start: new Date(education.start),
            end: education.end ? new Date(education.end) : undefined
        })) || [];

        console.log(resume.value)
    }

    const { data: biography } = await useAsyncData('biography', () => queryContent<BiographyModel>('biography').findOne())

    useHead({
        title: `${biography.value?.name} - ${biography.value?.role}`,
        meta: [
            { name: 'description', content: biography.value?.about }
        ],
    });
</script>
