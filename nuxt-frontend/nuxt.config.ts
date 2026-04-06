// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  runtimeConfig: {
    public: {
      apiBase: 'http://127.0.0.1:8081/api/v1',
      userApiBase: 'http://127.0.0.1:8085'
    }
  },
  css: ['~/assets/css/admin.css']
})
