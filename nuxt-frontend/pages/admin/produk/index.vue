<script setup>
definePageMeta({
  layout: 'admin'
})

// Mock data produk
const products = [
  { id: 1, name: 'Bayam Organik', category: 'Sayuran', price: 'Rp 5.000', stock: 50 },
  { id: 2, name: 'Tomat Merah', category: 'Sayuran', price: 'Rp 8.000', stock: 30 },
  { id: 3, name: 'Apel Fuji', category: 'Buah', price: 'Rp 15.000', stock: 20 },
]
</script>

<template>
  <div class="dashboard-content">
    <div class="header">
      <h1>Kelola Produk</h1>
      <p>Manajemen stok dan harga produk Anda.</p>
    </div>

    <div class="glass-panel">
      <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px;">
        <h3>Daftar Produk</h3>
        <button class="btn-primary" style="width: auto; padding: 10px 20px;">+ Tambah Produk</button>
      </div>
      
      <table style="width: 100%; border-collapse: collapse; color: #cbd5e1;">
        <thead>
          <tr style="text-align: left; border-bottom: 1px solid rgba(255,255,255,0.1);">
            <th style="padding: 12px;">ID</th>
            <th style="padding: 12px;">Nama</th>
            <th style="padding: 12px;">Kategori</th>
            <th style="padding: 12px;">Harga</th>
            <th style="padding: 12px;">Stok</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="p in products" :key="p.id" style="border-bottom: 1px solid rgba(255,255,255,0.05);">
            <td style="padding: 12px;">{{ p.id }}</td>
            <td style="padding: 12px;">{{ p.name }}</td>
            <td style="padding: 12px;">{{ p.category }}</td>
            <td style="padding: 12px;">{{ p.price }}</td>
            <td style="padding: 12px;">{{ p.stock }}</td>
          </tr>
        </tbody>
      </table>
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
