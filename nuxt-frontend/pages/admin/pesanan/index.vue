<script setup>
definePageMeta({
  layout: 'admin'
})

const config = useRuntimeConfig()
const statusOptions = ['Pending', 'Diproses', 'Dikirim', 'Selesai', 'Dibatalkan']
const updatingOrderId = ref(null)

const { data: ordersData, pending, error, refresh } = await useFetch(`${config.public.orderApiBase}/orders`)

const formatCurrency = (value) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0
  }).format(value || 0)
}

const formatDate = (value) => {
  if (!value) return '-'
  return new Intl.DateTimeFormat('id-ID', {
    dateStyle: 'medium',
    timeStyle: 'short'
  }).format(new Date(value))
}

const getStatusColor = (status) => {
  switch (status) {
    case 'Pending': return '#f59e0b';
    case 'Diproses': return '#38bdf8';
    case 'Dikirim': return '#60a5fa';
    case 'Selesai': return '#10b981';
    case 'Dibatalkan': return '#ef4444';
    default: return '#94a3b8';
  }
}

const getItemsSummary = (items) => {
  return items?.map(item => `${item.product_name} x${item.quantity}`).join(', ') || '-'
}

const updateStatus = async (orderId, status) => {
  updatingOrderId.value = orderId

  try {
    await $fetch(`${config.public.orderApiBase}/orders/${orderId}/status`, {
      method: 'PATCH',
      body: { status }
    })
    await refresh()
  } catch (err) {
    alert(err?.data?.error || err?.message || 'Gagal mengubah status pesanan')
  } finally {
    updatingOrderId.value = null
  }
}
</script>

<template>
  <div class="dashboard-content">
    <div class="header">
      <h1>Pesanan Pelanggan</h1>
      <p>Kelola dan pantau status pesanan yang masuk.</p>
    </div>

    <div v-if="pending" class="glass-panel">
      <h3>Riwayat Pesanan</h3>
      <p style="margin-top: 20px; color: #94a3b8;">Sedang memuat data pesanan...</p>
    </div>

    <div v-else-if="error" class="glass-panel">
      <h3>Riwayat Pesanan</h3>
      <p style="margin-top: 20px; color: #ef4444;">Gagal memuat data pesanan: {{ error.message }}</p>
    </div>

    <div v-else class="glass-panel">
      <div class="panel-head">
        <h3>Riwayat Pesanan</h3>
        <button class="refresh-btn" @click="refresh">Muat Ulang</button>
      </div>
      <table style="width: 100%; border-collapse: collapse; color: #cbd5e1; margin-top: 20px;">
        <thead>
          <tr style="text-align: left; border-bottom: 1px solid rgba(255,255,255,0.1);">
            <th style="padding: 12px;">Order ID</th>
            <th style="padding: 12px;">Pelanggan</th>
            <th style="padding: 12px;">Item</th>
            <th style="padding: 12px;">Total</th>
            <th style="padding: 12px;">Status</th>
            <th style="padding: 12px;">Waktu</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="!ordersData?.data?.length">
            <td colspan="6" style="padding: 24px; color: #94a3b8; text-align: center;">Belum ada pesanan masuk.</td>
          </tr>
          <tr v-for="o in ordersData?.data" :key="o.id" style="border-bottom: 1px solid rgba(255,255,255,0.05);">
            <td style="padding: 12px; color: #38bdf8;">{{ o.order_code }}</td>
            <td style="padding: 12px;">
              <strong>{{ o.customer_name }}</strong>
              <p style="margin: 6px 0 0; color: #94a3b8; font-size: 13px;">User ID: {{ o.user_id }}</p>
            </td>
            <td style="padding: 12px;">{{ getItemsSummary(o.items) }}</td>
            <td style="padding: 12px;">{{ formatCurrency(o.total_amount) }}</td>
            <td style="padding: 12px;">
              <div class="status-cell">
                <span class="status-badge" :style="{ color: getStatusColor(o.status) }">
                  {{ o.status }}
                </span>
                <select class="status-select" :value="o.status" :disabled="updatingOrderId === o.id" @change="updateStatus(o.id, $event.target.value)">
                  <option v-for="status in statusOptions" :key="status" :value="status">{{ status }}</option>
                </select>
              </div>
            </td>
            <td style="padding: 12px;">{{ formatDate(o.created_at) }}</td>
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
  background-clip: text;
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

.panel-head {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
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

.status-badge {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  border: 1px solid currentColor;
  background: rgba(255, 255, 255, 0.05);
}

.status-cell {
  display: grid;
  gap: 10px;
}

.status-select,
.refresh-btn {
  background: rgba(15, 23, 42, 0.9);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: #e2e8f0;
  border-radius: 10px;
  padding: 10px 12px;
}

.refresh-btn {
  cursor: pointer;
}
</style>
