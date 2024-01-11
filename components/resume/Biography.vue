<template>
    <p>{{ biography.about }}</p>
    <ul class="mt-2 space-y-2">
        <li>{{ biography.location }}</li>
        <li>
            <a :href="biography.github">Github</a>
        </li>
        <li>
            <a :href="biography.linkedIn">LinkedIn</a>
        </li>
        <li>
            <a :href="biography.x">X</a>
        </li>
    </ul>
</template>

<script setup lang="ts">
    import type { BiographyModel } from './models';

    defineProps<{ biography: BiographyModel; }>();
</script>
