<template>
    <main class="max-w-[820px] mx-4 md:mx-auto">
        <h1 class="text-white text-4xl md:text-6xl font-bold mb-4 mt-8">
            {{ $t('blog.title') }}
        </h1>
        <p>
            {{ $t('blog.subtitle') }}
        </p>
        <div class="mt-8" v-if="blog?.length == 0">
            {{ $t('blog.noContent') }}
        </div>
        <ol class="divide-y divide-rose-200 dark:divide-rose-300">
            <li
                v-for="blogEntry in blog"
                class="space-y-4 py-8"
            >
                <ProjectPreview
                    :link="blogEntry._path"
                    :title="blogEntry.title"
                    :description="blogEntry.description"
                    :previewImage="blogEntry.previewImage"
                    :github="blogEntry.github"
                />
            </li>
        </ol>
    </main>
</template>

<script setup lang="ts">
    useHead({
        title: 'Lukas Grimm - Blog',
        meta: [
            { name: 'description', content: 'Blog of Lukas Grimm' }
        ]
    });

    type BlogContent = {
        _path: string;
        title: string;
        description: string;
        github: string;
        previewImage?: {
            source: string;
            alt: string;
        }
    };

    const { data: blog } = await useAsyncData('blog', () => queryContent<BlogContent>('blog').without('body').find().then(data => data.reverse()));
</script>
