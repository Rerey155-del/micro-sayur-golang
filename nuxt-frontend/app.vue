<script setup>
const config = useRuntimeConfig()

// Mengakses API Product Service (port 8081)
const { data: products, pending: productsPending, error: productsError } = await useFetch(`${config.public.apiBase}/products`, {
  key: 'products-data',
  server: false
})

// Mengakses API User Service (port 8080)
const { data: userData, pending: usersPending, error: usersError } = await useFetch(`${config.public.userApiBase}/users`, {
  key: 'users-data',
  server: false
})

// Fungsi helper format mata uang IDR
const formatIDR = (price) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0
  }).format(price)
}
</script>

<template>
  <div class="container">
    <header>
      <h1>MicroSayur Premium</h1>
      <p style="color: #94a3b8">Sistem E-Commerce Microservices (Golang + Nuxt 3)</p>
    </header>

    <main>
      <!-- Section Daftar Produk -->
      <section>
        <h2 class="section-title">🥬 Katalog Sayur Segar</h2>
        
        <div v-if="productsPending" class="loading">Memuat produk...</div>
        
        <div v-else-if="productsError" class="loading" style="color: #ef4444">
          Gagal mengambil data produk. Pastikan Product Service (8081) menyala.
        </div>
        
        <div v-else class="grid">
          <div v-for="product in products?.data" :key="product.id" class="card">
            <div>
              <h3>{{ product.name }}</h3>
              <p>{{ product.description }}</p>
            </div>
            <div class="price-row">
              <span class="price">{{ formatIDR(product.price) }}</span>
              <span class="stock">📦 Stok: {{ product.stock }}</span>
            </div>
            <button class="btn-add">✨ Tambah ke Keranjang</button>
          </div>
        </div>
      </section>

      <!-- Section Daftar User (Data dari User Service) -->
      <section style="margin-top: 4rem">
        <h2 class="section-title">👥 Pengguna Terdaftar (User Service)</h2>
        
        <div v-if="usersPending" class="loading">Memuat data user...</div>
        
        <div v-else-if="usersError" class="loading" style="color: #ef4444">
          Gagal mengambil data user. Pastikan User Service (8080) menyala.
        </div>
        
        <div v-else class="user-list">
          <div v-for="user in userData?.data" :key="user.id" class="user-item">
            <div class="user-avatar">{{ user.name.charAt(0) }}</div>
            <div class="user-info">
              <h4>{{ user.name }}</h4>
              <p>{{ user.email }} • <span class="badge">{{ user.role_name }}</span></p>
            </div>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>


