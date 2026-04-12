// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  runtimeConfig: {
    public: {
      // Nuxt otomatis membaca NUXT_PUBLIC_API_BASE dari environment variable
      // Fallback ke localhost untuk development lokal
      apiBase: 'http://127.0.0.1:8081/api/v1',
      userApiBase: 'http://127.0.0.1:8085',
      orderApiBase: 'http://127.0.0.1:8082/api/v1'
    }
  },
  css: ['~/assets/css/admin.css'],
})
