<template>
    <main class="max-w-[820px] mx-4 md:mx-auto">
        <h1 class="text-white text-4xl md:text-6xl font-bold mb-4 mt-8">
            {{ $t('projects.title') }}
        </h1>
        <p>
            {{ $t('projects.subtitle') }}
        </p>
        <div class="mt-8" v-if="projects?.length == 0">
            {{ $t('project.noContent') }}
        </div>
        <ol class="divide-y divide-rose-200 dark:divide-rose-300">
            <li
                v-for="project in projects"
                class="space-y-4 py-8"
            >
                <ProjectPreview
                    :link="project._path"
                    :title="project.title"
                    :description="project.description"
                    :previewImage="project.previewImage"
                    :github="project.github"
                />
            </li>
        </ol>
    </main>
</template>

<script setup lang="ts">
    useHead({
        title: 'Lukas Grimm - Projects',
        meta: [
            { name: 'description', content: 'Projects of Lukas Grimm' }
        ]
    });

    type ProjectContent = {
        _path: string;
        title: string;
        description: string;
        github: string;
        previewImage?: {
            source: string;
            alt: string;
        }
    };

    const { data: projects } = await useAsyncData('projects', () => queryContent<ProjectContent>('projects').without('body').find().then(data => data.reverse()));
</script>
