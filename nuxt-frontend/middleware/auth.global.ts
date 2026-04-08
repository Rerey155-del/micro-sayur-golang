export default defineNuxtRouteMiddleware((to, from) => {
  // Hanya jalankan di browser (client-side)
  if (import.meta.client) {
    const token = localStorage.getItem('token')
    const role = localStorage.getItem('user_role')

    // Pengecekan akses halaman Admin
    if (to.path.startsWith('/admin')) {
      if (!token) return navigateTo('/')
      // Cegah customer masuk ke wilayah admin
      if (role === 'Customer') return navigateTo('/customer/home')
    }

    // Pengecekan akses halaman Customer
    if (to.path.startsWith('/customer')) {
      if (!token) return navigateTo('/')
      // Jika butuh, bisa tambahkan cegah admin masuk wilayah customer di sini nantinya
    }

    // Jika sudah login dan mencoba ke halaman login (/), redirect sesuai rolenya
    if (to.path === '/' && token) {
      if (role === 'Customer') {
        return navigateTo('/customer/home')
      } else {
        return navigateTo('/admin')
      }
    }
  }
})
