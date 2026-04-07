<script setup>
import { ref, onMounted } from "vue";

const adminName = ref("Admin");

onMounted(() => {
  const name = localStorage.getItem("user_name");
  if (name) {
    adminName.value = name;
  }
});

const doLogout = () => {
  localStorage.removeItem("token");
  localStorage.removeItem("user_name");
  navigateTo("/");
};
</script>

<template>
  <div class="admin-layout">
    <aside class="sidebar">
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
    <!-- konten kanan -->
    <main class="main-wrapper">
      <header class="topbar">
        <div class="topbar-left">
          <h3>Panel Administrator</h3>
        </div>
        <div class="topbar-right">
          <div class="user-profile">
            <span class="admin-name">{{ adminName }}</span>
            <div class="avatar">
              {{ adminName.charAt(0) }}
            </div>
          </div>
        </div>
      </header>

      <section class="content-area">
        <slot />
      </section>
    </main>
  </div>
</template>

<style scoped>
/* Reset & Variabel Warna - Disinkronkan dengan Login */
.admin-layout {
  --primary: #38bdf8; /* Sky Blue dari Login */
  --primary-dark: #0ea5e9;
  --bg-color: #0f172a;
  --bg-gradient: linear-gradient(135deg, #0f172a 0%, #1e1b4b 100%);
  --sidebar-bg: rgba(30, 41, 59, 0.7);
  --glass-bg: rgba(255, 255, 255, 0.03);
  --text-main: #f8fafc;
  --text-muted: #94a3b8;

  display: flex;
  min-height: 100vh;
  background: var(--bg-gradient);
  color: var(--text-main);
  font-family: "Inter", sans-serif;
}

/* Sidebar Styling */
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

/* Class Vue Router otomatis untuk link aktif */
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

/* Area Konten Utama */
.main-wrapper {
  flex: 1;
  margin-left: 260px;
  display: flex;
  flex-direction: column;
}

/* Header Topbar */
.topbar {
  height: 80px;
  background: rgba(15, 23, 42, 0.8);
  backdrop-filter: blur(10px);
  padding: 0 40px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  position: sticky;
  top: 0;
  z-index: 5;
}

.topbar-left h3 {
  margin: 0;
  font-size: 18px;
  color: var(--text-muted);
  font-weight: 600;
}

.user-profile {
  display: flex;
  align-items: center;
  gap: 16px;
}

.admin-name {
  font-weight: 500;
  color: var(--text-main);
}

.avatar {
  background: linear-gradient(135deg, #0ea5e9, #3b82f6);
  color: white;
  width: 40px;
  height: 40px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  font-size: 18px;
  box-shadow: 0 4px 12px rgba(14, 165, 233, 0.3);
}

/* Area render NuxtPage */
.content-area {
  padding: 40px;
  flex: 1;
  overflow-y: auto;
}
</style>
