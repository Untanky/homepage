<template>
  <footer class="legal-footer">
    <hr class="rule">
    <div class="footer-container">
      <nuxt-link
        class="footer-link"
        :to="localePath('/imprint')"
      >{{ $t('pages.imprint') }}</nuxt-link>
      <nuxt-link
        class="footer-link"
        :to="localePath('/data-protection')"
      >{{ $t('pages.dataProtection') }}</nuxt-link>
      <nuxt-link
        class="footer-link"
        :to="localePath('/sitemap')"
      >{{ $t('pages.sitemap') }}</nuxt-link>
    </div>
    <div class="footer-container">
      <span>
        {{ `${$t('language')}:` }}
      </span>
      <nuxt-link
        class="footer-link"
        v-for="(link, index, key) in languages"
        :key="key"
        :to="switchLocalePath(link.code)"
      >{{ link.name }}</nuxt-link>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'LegalFooter',
  data() {
    return {
      selected: this.$i18n.locale,
      languages: this.$i18n.locales,
    };
  },
  methods: {
    changeLanguage(evt) {
      this.$router.push(this.switchLocalePath(evt));
    },
  },
};
</script>

<style lang='scss' scoped>
@import "../assets/scss/variables";

.legal-footer {
  .rule {
    border: 1px solid $primary;
  }

  .footer-container {
    .footer-link {
      margin: 1em;
    }

    .language-switcher {
      display: inline-block;
      padding: 1em;
    }
  }
}
</style>
