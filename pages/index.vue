<template>
    <div class="max-w-[820px] mx-4 md:mx-auto grid grid-cols-1 grid-rows-[400px_auto_auto_300px_auto_auto_auto_auto_auto] md:grid-rows-[auto_auto_auto_1fr_auto_auto] md:grid-cols-[8fr_12fr_1fr_19fr_6fr_8fr_7fr_8fr_4fr_15fr_8fr_6fr] gap-4">
        <LandingPicture class="md:col-span-6 md:row-span-4 bg-[url('/profile.jpg')]">
            <h1 class="text-white text-4xl md:text-6xl font-bold">
                {{ content!.name }}
            </h1>
        </LandingPicture>
        <LandingCard class="md:col-span-6">
            <h2 class="text-2xl md:text-4xl font-medium">
                {{ content!.role }}
            </h2>
        </LandingCard>
        <LandingCard class="!p-0 md:col-span-3">
            <NuxtLink to="/resume" class="block !px-4 !py-2">
                See Resumé
            </NuxtLink>
        </LandingCard>
        <LandingCard class="!p-0 md:col-span-3">
            <NuxtLink to="/projects" class="block !px-4 !py-2 ">
                Explore Projects
            </NuxtLink>
        </LandingCard>
        <LandingCard class="md:col-span-6">
            <div class="text-gray-700 dark:text-zinc-400 mb-2">About me</div>
            <div>
                {{ content!.about }}
            </div>
        </LandingCard>
        <LandingPicture class="md:col-start-7 min-h-[250px] md:col-span-6 md:row-span-3 bg-[url('/map.png')]">
            <div class="text-black text-xl">
                {{ content!.location }}
            </div>
        </LandingPicture>
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
        <LandingCard class="md:col-start-1 md:col-span-2 md:row-span-1 !p-0">
            <a :href="content!.github" class="block px-4 py-2">
                GitHub
            </a>
        </LandingCard>
        <LandingCard class="md:col-span-2 md:row-span-1 !p-0">
            <a :href="content!.linkedIn" class="block px-4 py-2">
                LinkedIn
            </a>
        </LandingCard>
        <LandingCard class="md:col-span-2 md:row-span-1 !p-0">
            <a :href="content!.x" class="block px-4 py-2">
                X
            </a>
        </LandingCard>
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
    });
</script>
