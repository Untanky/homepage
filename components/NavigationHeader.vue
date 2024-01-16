<template>
    <nav class="text-lg bg-gradient-to-r from-rose-300 via-amber-300 to-green-300 dark:text-black md:max-w-[820px] px-1.5 py-1.5 rounded-xl my-3 mx-4 md:mx-auto">
        <ul class="flex justify-start">
            <li v-for="link in links" class="hover:backdrop-saturate-200 hover:backdrop-brightness-90 rounded-lg">
                <NuxtLink class="inline-block px-3 py-1.5" :to="link.to">
                    {{ link.name }}
                </NuxtLink>
            </li>
        </ul>
    </nav>
</template>

<script setup lang="ts">
    defineProps<{
        links: Array<{ to: string, name: string }>
    }>();
</script>
