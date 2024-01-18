<template>
    <header class="flex flex-col-reverse md:relative md:my-12">
        <div class="text-center md:text-left py-8 px-6 bg-zinc-950 border border-rose-300 rounded-xl -mt-6 md:mt-0">
            <h1 class="text-4xl md:text-6xl font-semibold md:font-bold">{{ name }}</h1>
            <h2 class="text-2xl font-medium md:font-semibold">{{ role }}</h2>
        </div>
        <img class="self-center z-10 md:absolute md:-right-px md:-top-8 md:-bottom-8 rounded-xl h-[180px] md:h-[220px]" src="/profile.jpg" alt="Profile Picture" />
    </header>
</template>

<script setup lang="ts">
    defineProps<{ name: string; role: string }>();
</script>
