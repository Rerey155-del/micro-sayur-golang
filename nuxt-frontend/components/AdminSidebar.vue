<script setup>
const props = defineProps(['isOpen'])

const doLogout = () => {
  localStorage.removeItem("token");
  localStorage.removeItem("user_name");
  navigateTo("/");
};
</script>

<template>
  <aside class="sidebar" :class="{ 'is-open': isOpen }">
    <div class="sidebar-brand">
      <h2>Micro<span>sayur</span></h2>
    </div>
    <nav class="sidebar-nav">
      <NuxtLink to="/admin" class="nav-item">
        <span class="label">Dashboard</span>
      </NuxtLink>
      <NuxtLink to="/admin/produk" class="nav-item">
        <span class="label">Kelola Produk</span>
      </NuxtLink>
      <NuxtLink to="/admin/pesanan" class="nav-item">
        <span class="label">Pesanan</span>
      </NuxtLink>
    </nav>
    <div class="sidebar-footer">
      <button @click="doLogout" class="btn-logout">
        <span>Logout</span>
      </button>
    </div>
  </aside>
</template>

<style scoped>
.sidebar {
  width: 260px;
  background: var(--sidebar-bg);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-right: 1px solid rgba(255, 255, 255, 0.05);
  display: flex;
  flex-direction: column;
  position: fixed;
  height: 100vh;
  z-index: 10;
}

.sidebar-brand {
  padding: 32px 24px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  text-align: center;
}

.sidebar-brand h2 {
  font-size: 24px;
  color: var(--text-main);
  margin: 0;
  font-weight: 700;
}

.sidebar-brand span {
  color: var(--primary);
}

.sidebar-nav {
  flex: 1;
  padding: 32px 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 18px;
  border-radius: 12px;
  color: var(--text-muted);
  text-decoration: none;
  font-weight: 500;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.05);
  color: var(--primary);
  transform: translateX(4px);
}

.router-link-exact-active {
  background: linear-gradient(90deg, rgba(56, 189, 248, 0.2), rgba(56, 189, 248, 0.05));
  color: var(--primary) !important;
  border-left: 4px solid var(--primary);
  border-radius: 4px 12px 12px 4px;
}

.sidebar-footer {
  padding: 24px;
  border-top: 1px solid rgba(255, 255, 255, 0.05);
}

.btn-logout {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
  border: 1px solid rgba(239, 68, 68, 0.2);
  padding: 12px;
  border-radius: 10px;
  font-weight: 600;
  cursor: pointer;
  transition: 0.3s ease;
}

.btn-logout:hover {
  background: #ef4444;
  color: white;
}

@media (max-width: 1024px) {
  .sidebar {
    left: -260px;
    transition: left 0.3s ease;
  }

  .sidebar.is-open {
    left: 0;
  }
}
</style>
