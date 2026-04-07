<script setup>
import { ref, onMounted } from "vue";

const adminName = ref("Admin");

onMounted(() => {
  const name = localStorage.getItem("user_name");
  if (name) {
    adminName.value = name;
  }
});
</script>

<template>
  <div class="admin-layout">
    <AdminSidebar />
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
