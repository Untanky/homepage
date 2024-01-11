<template>
    <nav class="bg-gradient-to-r from-rose-200 via-amber-200 to-green-200 md:max-w-[820px] px-1.5 py-1.5 rounded-xl my-3 mx-4 md:mx-auto">
        <ul class="flex justify-start">
            <li v-for="link in links" class="px-3 py-1.5 hover:backdrop-saturate-200 hover:backdrop-brightness-90 rounded-lg">
                <NuxtLink :to="link.to">
                    {{ link.name }}
                </NuxtLink>
            </li>
        </ul>
    </nav>
    <slot />
</template>

<script setup lang="ts">
const links = ref([
    {
        to: '/',
        name: 'Home',
    },
    {
        to: '/resume',
        name: 'Resumé',
    },
    {
        to: '/projects',
        name: 'Projects',
    },
    {
        to: '/blog',
        name: 'Blog',
    },
]);
</script>
