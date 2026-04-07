<script setup>
definePageMeta({
  layout: 'admin'
})

// Mock data pesanan
const orders = [
  { id: '#ORD-001', customer: 'John Doe', items: 'Bayam, Tomat', total: 'Rp 13.000', status: 'Pending' },
  { id: '#ORD-002', customer: 'Jane Smith', items: 'Apel Fuji', total: 'Rp 15.000', status: 'Success' },
  { id: '#ORD-003', customer: 'Budi Santoso', items: 'Wortel', total: 'Rp 7.000', status: 'Canceled' },
]

const getStatusColor = (status) => {
  switch (status) {
    case 'Pending': return '#f59e0b';
    case 'Success': return '#10b981';
    case 'Canceled': return '#ef4444';
    default: return '#94a3b8';
  }
}
</script>

<template>
  <div class="dashboard-content">
    <div class="header">
      <h1>Pesanan Pelanggan</h1>
      <p>Kelola dan pantau status pesanan yang masuk.</p>
    </div>

    <div class="glass-panel">
      <h3>Riwayat Pesanan</h3>
      <table style="width: 100%; border-collapse: collapse; color: #cbd5e1; margin-top: 20px;">
        <thead>
          <tr style="text-align: left; border-bottom: 1px solid rgba(255,255,255,0.1);">
            <th style="padding: 12px;">Order ID</th>
            <th style="padding: 12px;">Pelanggan</th>
            <th style="padding: 12px;">Item</th>
            <th style="padding: 12px;">Total</th>
            <th style="padding: 12px;">Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="o in orders" :key="o.id" style="border-bottom: 1px solid rgba(255,255,255,0.05);">
            <td style="padding: 12px; color: #38bdf8;">{{ o.id }}</td>
            <td style="padding: 12px;">{{ o.customer }}</td>
            <td style="padding: 12px;">{{ o.items }}</td>
            <td style="padding: 12px;">{{ o.total }}</td>
            <td style="padding: 12px;">
              <span class="status-badge" :style="{ color: getStatusColor(o.status) }">
                {{ o.status }}
              </span>
            </td>
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
</style>
