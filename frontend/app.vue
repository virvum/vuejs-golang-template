<template>
  <div id="app">
    <header>
      <nav>
        <router-link v-for="route in $router.options.routes.filter(r => !('meta' in r) || !('hidden' in r.meta) || r.meta.hidden === false)" :to="route.path" :key="route.path">
          {{ route.meta.title }}
        </router-link>
      </nav>
      <aside>
        Header
      </aside>
    </header>

    <main>
      <transition name="fade" mode="out-in">
        <router-view />
      </transition>
    </main>

    <footer>
      <a :href="app.homepage" target="_blank">{{ app.name }}</a> v{{ app.version }}
    </footer>
  </div>
</template>

<script>
export default {
  data: () => ({
    app: {
      name: process.env.NAME,
      version: process.env.VERSION,
      homepage: process.env.HOMEPAGE,
    },
    themes: require('@/themes.js'),
    theme: 'dark',
  }),
  methods: {
    toggleTheme () {
      this.theme = this.theme === 'dark' ? 'bright' : 'dark'
      this.setTheme()
    },
    setTheme () {
      for (const property in this.themes[this.theme]) {
        document.documentElement.style.setProperty(`--${property}`, this.themes[this.theme][property])
      }
    },
  },
  mounted () {
    this.setTheme()
  },
}
</script>

<style lang="sass">
*
  margin: 0
  padding: 0
  outline: 0
  border: 0
  color: inherit
  background: inherit
  text-decoration: none
  box-sizing: border-box
  font-family: Karla, sans-serif
  font-style: normal
  font-weight: normal
  font-size: $fontsize

html, body
  width: 100%
  height: 100%
  cursor: default
  background: $bg
  color: $fg

#app
  height: 100%
  display: flex
  flex-direction: column

  & > header
    display: flex
    background: $bg-surface

    & > nav
      flex: 1
      display: flex

      & > a
        padding: $padding

      & > a.active
        background: $bg-primary

    & > aside
      padding: $padding

  & > main
    flex: 1
    padding: $padding
    overflow: auto

  & > footer
    padding: $padding
    background: $bg-surface

.fade-enter-active, .fade-leave-active
  transition: opacity .1s

.fade-enter, .fade-leave-to
  opacity: 0
</style>
