<template>
    <main class="max-w-[820px] mx-4 md:mx-auto">
        <img
            v-if="project?.previewImage"
            class="w-full rounded-xl"
            :src="project?.previewImage.source"
            :alt="project?.previewImage.alt"
        >
        <ContentRenderer
            class="mt-8 prose lg:prose-xl dark:prose-invert prose-zinc prose-a:decoration-rose-300"
            :value="project!"
        />
    </main>
</template>

<script setup lang="ts">
    const { params } = useRoute();

    const { data: project } = await useAsyncData(`projects/${params.title}`, () => queryContent('projects', params.title as string).findOne());
    
    useHead({
        title: project.value?.title,
        meta: [
            { name: 'description', content: project.value?.description }
        ]
    });
</script>
