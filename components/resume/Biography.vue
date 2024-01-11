<template>
    <p>{{ biography.about }}</p>
    <ul class="mt-2 space-y-1">
        <li>
            <IconLocation class="mx-2" />
            <span class="align-middle">
                {{ biography.location }}
            </span>
        </li>
        <li>
            <a :href="biography.github">
                <IconGitHub class="mx-2" />
                <span class="align-middle">
                    Github
                </span>
            </a>
        </li>
        <li>
            <a :href="biography.linkedIn">
                <IconLinkedIn class="mx-2" />
                <span class="align-middle">
                    LinkedIn
                </span>
            </a>
        </li>
        <li>
            <a :href="biography.x">
                <IconTwitter class="mx-2" />
                <span class="align-middle">
                    X (formerly Twitter)
                </span>
            </a>
        </li>
    </ul>
</template>

<script setup lang="ts">
    import type { BiographyModel } from './models';

    defineProps<{ biography: BiographyModel; }>();
</script>
