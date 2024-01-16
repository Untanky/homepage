<template>
    <section>
        <h2 class="text-xl md:text-2xl font-medium underline decoration-rose-300 mb-2">
            {{ title }}
        </h2>
        <slot></slot>
    </section>
</template>

<script setup lang="ts">
    defineProps<{title: string}>();
</script>
