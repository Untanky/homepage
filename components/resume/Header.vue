<template>
    <header class="py-8 px-6 bg-zinc-950 border border-rose-300 rounded-xl">
        <h1 class="text-3xl md:text-6xl font-semibold md:font-bold">{{ name }}</h1>
        <h2 class="text-xl md:text-2xl font-medium md:font-semibold">{{ role }}</h2>
    </header>
</template>

<script setup lang="ts">
    defineProps<{ name: string; role: string }>();
</script>