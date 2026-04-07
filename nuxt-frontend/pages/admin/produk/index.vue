<script setup>
definePageMeta({
  layout: 'admin'
})

// Ambil konfigurasi API dari nuxt.config.ts
const config = useRuntimeConfig()

// Ambil data produk dari Product Service (port 8081)
const { data: productsData, pending, error } = await useFetch(`${config.public.apiBase}/products`)

// Fungsi pembantu untuk format mata uang Rupiah
const formatCurrency = (value) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0
  }).format(value)
}
</script>

<template>
  <div class="dashboard-content">
    <div class="header">
      <h1>Kelola Produk</h1>
      <p>Manajemen stok dan harga produk Anda.</p>
    </div>

    <!-- Tampilkan Loading jika data sedang dimuat -->
    <div v-if="pending" class="glass-panel" style="text-align: center; padding: 40px;">
      <p>Sedang memuat data produk...</p>
    </div>

    <!-- Tampilkan Error jika gagal mengambil data -->
    <div v-else-if="error" class="glass-panel" style="text-align: center; color: #ef4444; padding: 40px;">
      <p>Gagal mengambil data produk: {{ error.message }}</p>
    </div>

    <!-- Tampilkan data jika sukses -->
    <div v-else class="glass-panel">
      <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px;">
        <h3>Daftar Produk ({{ productsData?.data?.length || 0 }})</h3>
        <button class="btn-primary" style="width: auto; padding: 10px 20px;">+ Tambah Produk</button>
      </div>
      
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Nama</th>
              <th>Deskripsi</th>
              <th>Harga</th>
              <th>Stok</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="p in productsData?.data" :key="p.id">
              <td>{{ p.id }}</td>
              <td>{{ p.name }}</td>
              <td style="max-width: 300px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;">
                {{ p.description || '-' }}
              </td>
              <td>{{ formatCurrency(p.price) }}</td>
              <td>
                <span :style="{ color: p.stock < 10 ? '#ef4444' : '#10b981', fontWeight: 'bold' }">
                  {{ p.stock }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard-content {
  animation: fadeIn 0.8s ease-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.header h1 {
  font-size: 32px;
  font-weight: 700;
  margin-bottom: 8px;
  background: linear-gradient(90deg, #fff, #94a3b8);
  background-clip: text; /* Fixed lint warning */
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.header p {
  color: #94a3b8;
  font-size: 16px;
}

.glass-panel {
  margin-top: 40px;
  background: rgba(255, 255, 255, 0.03);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-radius: 20px;
  padding: 32px;
}

.glass-panel h3 {
  margin: 0;
  font-size: 20px;
  color: #fff;
}

/* Table Styling */
.table-container {
  overflow-x: auto;
  margin-top: 24px;
}

table {
  width: 100%;
  border-collapse: collapse;
  color: #cbd5e1;
}

th {
  text-align: left;
  padding: 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  color: #94a3b8;
  font-weight: 600;
  font-size: 14px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

td {
  padding: 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  font-size: 15px;
}

tr:hover td {
  background: rgba(255, 255, 255, 0.02);
  color: #fff;
}

/* Button Styling */
.btn-primary {
  background: linear-gradient(135deg, #0ea5e9, #3b82f6);
  color: white;
  border: none;
  padding: 12px 24px;
  border-radius: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(14, 165, 233, 0.2);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(14, 165, 233, 0.4);
}
</style>
