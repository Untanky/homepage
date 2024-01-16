<template>
    <NuxtLink
        :to="link"
        class="space-y-2 md:space-y-4"
    >
        <img
            v-if="previewImage"
            class="w-full rounded-xl"
            :src="previewImage.source"
            :alt="previewImage.alt"
        >
        <h2 class="text-white text-2xl md:text-4xl font-semibold">
            {{ title }}
        </h2>
        <p>
            {{ description }}
        </p>
    </NuxtLink>
    <nav class="flex justify-between">
        <a :href="github" class="bg-white dark:bg-zinc-950 border border-rose-200 px-4 py-2 rounded-lg">
            <IconGitHub class="w-4 h-4"/>
            <span class="ml-2 align-middle">GitHub</span>
        </a>
        <NuxtLink
            :to="link"
            class="bg-white dark:bg-rose-500 px-4 py-2 rounded-lg"
        >
            {{ $t('projects.learnMore') }}
        </NuxtLink>
    </nav>
</template>

<script setup lang="ts">
    defineProps<{
        link: string;
        title: string;
        description: string;
        previewImage?: { source: string; alt: string; };
        github: string; 
    }>();
</script>
