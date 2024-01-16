// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@nuxt/content',
    '@nuxtjs/i18n',
    '@nuxtjs/tailwindcss'
  ],
  tailwindcss: {
    exposeConfig: true
  },
  i18n: {
    vueI18n: './i18n.config.ts'
  }
})
