export default defineNuxtRouteMiddleware((to, from) => {
  // Hanya jalankan di browser (client-side)
  if (import.meta.client) {
    const token = localStorage.getItem('token')

    // Jika mencoba akses halaman admin tanpa token
    if (to.path.startsWith('/admin') && !token) {
      return navigateTo('/')
    }

    // Jika sudah login dan mencoba ke halaman login, redirect ke admin
    if (to.path === '/' && token) {
      return navigateTo('/admin')
    }
  }
})
