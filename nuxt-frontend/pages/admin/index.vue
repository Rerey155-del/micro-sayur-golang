<script setup>
definePageMeta({
  layout: 'admin'
})

const config = useRuntimeConfig()

// Ambil data produk dan user secara paralel
const { data: productsData } = await useFetch(`${config.public.apiBase}/products`)
const { data: usersData } = await useFetch(`${config.public.userApiBase}/users`)

// Hitung statistik riil
const totalProduk = computed(() => productsData.value?.data?.length || 0)
const totalPelanggan = computed(() => usersData.value?.data?.length || 0)

// Data statistik untuk ditampilkan di kartu
const stats = computed(() => [
  { title: 'Total Produk', value: totalProduk.value.toString(), color: '#10b981' },
  { title: 'Pesanan Baru', value: '0', color: '#3b82f6' },
  { title: 'Total Pelanggan', value: totalPelanggan.value.toString(), color: '#8b5cf6' },
  { title: 'Pendapatan (Rp)', value: '0', color: '#f59e0b' }
])

// Kosongkan aktivitas terbaru sesuai permintaan
const recentActivities = ref([])
</script>

<template>
  <div class="dashboard-content">
    <div class="header">
      <h1>Selamat Datang di Dashboard!</h1>
      <p>Berikut adalah ringkasan performa toko Microsayur hari ini.</p>
    </div>

    <div class="stats-grid">
      <div v-for="stat in stats" :key="stat.title" class="stat-card" :style="{ borderTopColor: stat.color }">
        <h4>{{ stat.title }}</h4>
        <h2>{{ stat.value }}</h2>
      </div>
    </div>

    <div class="glass-panel">
      <h3>Aktivitas Terbaru</h3>
      <div class="activity-list">
        <div v-if="recentActivities.length === 0" style="color: #94a3b8; text-align: center; padding: 20px;">
          Belum ada aktivitas saat ini.
        </div>
        <div v-for="activity in recentActivities" :key="activity.id" class="activity-item">
          <div class="bullet" :style="{ backgroundColor: activity.color }"></div>
          <div class="info">{{ activity.action }}</div>
          <div class="time">{{ activity.time }}</div>
        </div>
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
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.header p {
  color: #94a3b8;
  font-size: 16px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 24px;
  margin-top: 40px;
}

.stat-card {
  background: rgba(255, 255, 255, 0.03);
  backdrop-filter: blur(10px);
  padding: 24px;
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-top: 4px solid transparent;
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-5px);
  background: rgba(255, 255, 255, 0.07);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3);
}

.stat-card h4 {
  margin: 0;
  color: #94a3b8;
  font-size: 14px;
  font-weight: 500;
  letter-spacing: 0.5px;
}

.stat-card h2 {
  margin: 12px 0 0 0;
  font-size: 32px;
  font-weight: 700;
  color: #fff;
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
  margin: 0 0 24px 0;
  font-size: 20px;
  color: #fff;
}

.activity-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.activity-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.02);
  transition: 0.3s;
}

.activity-item:hover {
  background: rgba(255, 255, 255, 0.05);
  border-color: rgba(255, 255, 255, 0.1);
}

.bullet {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  box-shadow: 0 0 10px currentColor;
}

.info {
  flex: 1;
  color: #e2e8f0;
  font-size: 14px;
}

.time {
  color: #64748b;
  font-size: 12px;
}
</style>
