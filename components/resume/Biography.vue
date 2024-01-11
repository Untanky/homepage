<template>
    <p>{{ biography.about }}</p>
    <ul class="mt-2 space-y-1">
        <li>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-4 h-4 mx-2 inline-block">
                <path fill-rule="evenodd" d="m9.69 18.933.003.001C9.89 19.02 10 19 10 19s.11.02.308-.066l.002-.001.006-.003.018-.008a5.741 5.741 0 0 0 .281-.14c.186-.096.446-.24.757-.433.62-.384 1.445-.966 2.274-1.765C15.302 14.988 17 12.493 17 9A7 7 0 1 0 3 9c0 3.492 1.698 5.988 3.355 7.584a13.731 13.731 0 0 0 2.273 1.765 11.842 11.842 0 0 0 .976.544l.062.029.018.008.006.003ZM10 11.25a2.25 2.25 0 1 0 0-4.5 2.25 2.25 0 0 0 0 4.5Z" clip-rule="evenodd" />
            </svg>
            <span class="align-middle">
                {{ biography.location }}
            </span>
        </li>
        <li>
            <a :href="biography.github">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4 mx-2 inline-block">
                    <path fill-rule="evenodd" d="M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1ZM5.657 3.023a5.5 5.5 0 1 0 7.584 3.304l-.947-.63a.431.431 0 0 0-.544.053.431.431 0 0 1-.544.054l-.467-.312a.475.475 0 0 0-.689.608l.226.453a2.119 2.119 0 0 1 0 1.894L10.1 8.8a.947.947 0 0 0-.1.424v.11a2 2 0 0 1-.4 1.2L8.8 11.6A1 1 0 0 1 7 11v-.382a1 1 0 0 0-.553-.894l-.422-.212A1.854 1.854 0 0 1 6.855 6h.707a.438.438 0 1 0-.107-.864l-.835.209a1.129 1.129 0 0 1-1.305-1.553l.342-.77Z" clip-rule="evenodd" />
                </svg>
                <span class="align-middle">
                    Github
                </span>
            </a>
        </li>
        <li>
            <a :href="biography.linkedIn">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4 mx-2 inline-block">
                    <path fill-rule="evenodd" d="M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1ZM5.657 3.023a5.5 5.5 0 1 0 7.584 3.304l-.947-.63a.431.431 0 0 0-.544.053.431.431 0 0 1-.544.054l-.467-.312a.475.475 0 0 0-.689.608l.226.453a2.119 2.119 0 0 1 0 1.894L10.1 8.8a.947.947 0 0 0-.1.424v.11a2 2 0 0 1-.4 1.2L8.8 11.6A1 1 0 0 1 7 11v-.382a1 1 0 0 0-.553-.894l-.422-.212A1.854 1.854 0 0 1 6.855 6h.707a.438.438 0 1 0-.107-.864l-.835.209a1.129 1.129 0 0 1-1.305-1.553l.342-.77Z" clip-rule="evenodd" />
                </svg>
                <span class="align-middle">
                    LinkedIn
                </span>
            </a>
        </li>
        <li>
            <a :href="biography.x">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4 mx-2 inline-block">
                    <path fill-rule="evenodd" d="M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1ZM5.657 3.023a5.5 5.5 0 1 0 7.584 3.304l-.947-.63a.431.431 0 0 0-.544.053.431.431 0 0 1-.544.054l-.467-.312a.475.475 0 0 0-.689.608l.226.453a2.119 2.119 0 0 1 0 1.894L10.1 8.8a.947.947 0 0 0-.1.424v.11a2 2 0 0 1-.4 1.2L8.8 11.6A1 1 0 0 1 7 11v-.382a1 1 0 0 0-.553-.894l-.422-.212A1.854 1.854 0 0 1 6.855 6h.707a.438.438 0 1 0-.107-.864l-.835.209a1.129 1.129 0 0 1-1.305-1.553l.342-.77Z" clip-rule="evenodd" />
                </svg>
                <span class="align-middle">
                    X
                </span>
            </a>
        </li>
    </ul>
</template>

<script setup lang="ts">
    import type { BiographyModel } from './models';

    defineProps<{ biography: BiographyModel; }>();
</script>
