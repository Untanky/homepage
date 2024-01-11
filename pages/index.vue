<template>
    <div class="max-w-[820px] mx-4 md:mx-auto grid grid-cols-1 grid-rows-[400px_auto_auto_auto_300px_auto_auto_auto_auto] md:grid-rows-[auto_auto_1fr_auto_auto] md:grid-cols-6 gap-4">
        <div class="relative md:min-h-[400px] md:col-span-3 rounded-xl md:row-span-3 bg-[url('/profile.jpg')] bg-cover bg-center">
            <h1 class="absolute bottom-0 left-0 right-0 text-center p-3 rounded-b-xl backdrop-blur-lg text-white text-4xl md:text-6xl font-bold">
                {{ content!.name }}
            </h1>
        </div>
        <h2 class="bento-card rounded-xl text-2xl md:text-4xl font-medium md:col-span-3">
            {{ content!.role }}
        </h2>
        <div class="bento-card rounded-xl">
            <div class="text-zinc-700 dark:text-zinc-400">Experience</div>
            <div class="text-xl">
                {{ content!.experience }}
            </div>
        </div>
        <div class="bento-card rounded-xl md:col-span-2">
            <div class="text-zinc-700 dark:text-zinc-400">Education</div>
            <div class="text-xl">
                {{ content!.education }}
            </div>
        </div>
        <div class="relative md:col-span-3 rounded-xl md:row-span-2 bg-[url('/map.png')] bg-cover">
            <div class="absolute bottom-0 left-0 right-0 text-center p-3 rounded-b-xl backdrop-blur-lg text-white text-xl">
                {{ content!.location }}
            </div>
        </div>
        <div class="bento-card rounded-xl md:col-span-3 md:row-span-2">
            <div class="text-gray-700 dark:text-zinc-400 mb-2">About me</div>
            <div>
                {{ content!.about }}
            </div>
        </div>
        <a :href="content!.github" class="bento-card !px-4 !py-2 rounded-xl col-span-1 row-span-1">
            GitHub
        </a>
        <a :href="content!.linkedIn" class="bento-card !px-4 !py-2 rounded-xl col-span-1 row-span-1">
            LinkedIn
        </a>
        <NuxtLink to="/resume" class="bento-card !px-4 !py-2 rounded-xl col-span-1 row-span-1">
            Resumé
        </NuxtLink>
    </div>
</template>

<script setup lang="ts">
type BiograhphyContent = {
    name: string;
    role: string;
    experience: string;
    education: string;
    location: string;
    about: string;
    github: string;
    linkedIn: string
};

const { data: content } = await useAsyncData('biography', () => queryContent<BiograhphyContent>('biography').findOne())

useHead({
    title: `${content.value?.name} - ${content.value?.role}`,
    meta: [
        { name: 'description', content: content.value?.about }
    ],
})
</script>

<style lang="postcss" scoped>
.bento-card {
    @apply dark:bg-zinc-950 border dark:border-rose-300 border-rose-200 px-4 py-6;
}
</style>
