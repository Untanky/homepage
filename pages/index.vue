<template>
    <div class="max-w-[820px] mx-4 md:mx-auto grid grid-cols-1 grid-rows-[400px_auto_auto_300px_auto_auto_auto_auto_auto] md:grid-rows-[auto_auto_auto_1fr_auto_auto] md:grid-cols-12 gap-4">
        <LandingPicture class="md:col-span-6 md:row-span-4" :subtitle="content!.name" as="h1" />
        <LandingCard class="md:col-span-6">
            <h2 class="text-2xl md:text-4xl font-medium">
                {{ content!.role }}
            </h2>
        </LandingCard>
        <NuxtLink to="/resume" class="bento-card !px-4 !py-2 rounded-xl col-span-3">
            See Resumé
        </NuxtLink>
        <NuxtLink to="/projects" class="bento-card !px-4 !py-2 rounded-xl col-span-3">
            Explore Projects
        </NuxtLink>
        <LandingCard class="md:col-span-6">
            <div class="text-gray-700 dark:text-zinc-400 mb-2">About me</div>
            <div>
                {{ content!.about }}
            </div>
        </LandingCard>
        <div class="relative md:col-start-7 md:col-span-6 rounded-xl md:row-span-3 bg-[url('/map.png')] bg-cover bg-center">
            <div class="absolute bottom-0 left-0 right-0 text-center p-3 rounded-b-xl backdrop-blur-lg text-black text-xl">
                {{ content!.location }}
            </div>
        </div>
        <LandingCard class="md:col-span-3">
            <div class="text-zinc-700 dark:text-zinc-400 mb-2">Experience</div>
            <div class="text-xl">
                {{ content!.experience }}
            </div>
        </LandingCard>
        <LandingCard class="md:col-span-3">
            <div class="text-zinc-700 dark:text-zinc-400 mb-2">Education</div>
            <div class="text-xl">
                {{ content!.education }}
            </div>
        </LandingCard>
        <a :href="content!.github" class="bento-card !px-4 !py-2 rounded-xl col-start-1 col-span-2 row-span-1">
            GitHub
        </a>
        <a :href="content!.linkedIn" class="bento-card !px-4 !py-2 rounded-xl col-span-2 row-span-1">
            LinkedIn
        </a>
        <a :href="content!.x" class="bento-card !px-4 !py-2 rounded-xl col-span-2 row-span-1">
            X
        </a>
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
    linkedIn: string;
    x: string;
};

const { data: content } = await useAsyncData(() => queryContent<BiograhphyContent>('biography').findOne())

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
