<script setup>
definePageMeta({
    layout: false
})

const userName = ref('')
const userRole = ref('')

onMounted(() => {
    userName.value = localStorage.getItem('user_name') || 'Pelanggan'
    userRole.value = localStorage.getItem('user_role') || 'Customer'
})

const logout = () => {
    localStorage.clear()
    navigateTo('/')
}

const isMenuOpen = ref(false)
const toggleMenu = () => {
    isMenuOpen.value = !isMenuOpen.value
}
</script>

<template>
    <div class="landing-container">
        <!-- Navbar -->
        <nav class="navbar glass">
            <div class="logo">MicroSayur</div>
            
            <!-- Hamburger Button -->
            <button class="hamburger" @click="toggleMenu" :class="{ 'is-active': isMenuOpen }">
                <span></span>
                <span></span>
                <span></span>
            </button>

            <div :class="['nav-links', { 'mobile-open': isMenuOpen }]">
                <NuxtLink to="/customer/home" @click="isMenuOpen = false">Beranda</NuxtLink>
                <NuxtLink to="/customer/produk" @click="isMenuOpen = false">Produk</NuxtLink>
                <NuxtLink to="/customer/pesanan" @click="isMenuOpen = false">Pesanan</NuxtLink>
                <button @click="logout" class="btn-logout">Keluar</button>
            </div>
        </nav>

        <!-- Hero Section -->
        <main class="hero">
            <div class="hero-content">
                <p class="welcome-text">Selamat Datang, {{ userName }}!</p>
                <h1 class="title">Belanja Sayur Segar <br> <span class="accent">Jadi Lebih Mudah</span></h1>
                <p class="description">
                    Dapatkan sayuran organik berkualitas langsung dari petani lokal. 
                    Segar, higienis, dan diantar sampai depan pintu rumah Anda.
                </p>
                <div class="cta-buttons">
                    <button class="btn-primary" @click="navigateTo('/customer/produk')">Mulai Belanja</button>
                    <button class="btn-secondary">Lihat Promo</button>
                </div>
            </div>
            
            <div class="hero-visual">
                <div class="circle-gradient"></div>
                <div class="glass-card floating">
                    <div class="card-header">
                        <div class="dot red"></div>
                        <div class="dot yellow"></div>
                        <div class="dot green"></div>
                    </div>
                    <div class="card-body">
                        <h3>Pesanan Anda</h3>
                        <div class="mock-item">
                            <div class="img-placeholder"></div>
                            <div class="text-lines">
                                <div class="line long"></div>
                                <div class="line short"></div>
                            </div>
                        </div>
                        <div class="status-badge">Sedang Dikirim</div>
                    </div>
                </div>
            </div>
        </main>

        <!-- Stats Section -->
        <section class="quick-stats">
            <div class="stat-item">
                <h2>10k+</h2>
                <p>Pengguna Aktif</p>
            </div>
            <div class="stat-item">
                <h2>500+</h2>
                <p>Petani Lokal</p>
            </div>
            <div class="stat-item">
                <h2>100%</h2>
                <p>Segar & Organik</p>
            </div>
        </section>
    </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap');

.landing-container {
    --primary: #10b981;
    --primary-glow: rgba(16, 185, 129, 0.4);
    --bg-dark: #0f172a;
    --text-main: #f8fafc;
    --text-muted: #94a3b8;
    --glass-bg: rgba(255, 255, 255, 0.03);
    --glass-border: rgba(255, 255, 255, 0.08);

    min-height: 100vh;
    background-color: var(--bg-dark);
    color: var(--text-main);
    font-family: 'Plus Jakarta Sans', sans-serif;
    background-image: radial-gradient(circle at 0% 0%, rgba(16, 185, 129, 0.08) 0%, transparent 50%),
                radial-gradient(circle at 100% 100%, rgba(59, 130, 246, 0.08) 0%, transparent 50%);
    overflow-x: hidden;
}

.glass {
    background: var(--glass-bg);
    backdrop-filter: blur(12px);
    border: 1px solid var(--glass-border);
}

/* Navbar */
.navbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px 8%;
    position: fixed;
    top: 0;
    width: 100%;
    z-index: 1000;
    border-top: none;
    border-left: none;
    border-right: none;
}

.logo {
    font-size: 24px;
    font-weight: 800;
    color: var(--primary);
    letter-spacing: -1px;
}

.nav-links {
    display: flex;
    align-items: center;
    gap: 32px;
}

.nav-links a {
    text-decoration: none;
    color: var(--text-muted);
    font-size: 14px;
    font-weight: 500;
    transition: color 0.3s;
}

.nav-links a:hover, .nav-links a.router-link-active {
    color: var(--primary);
}

.btn-logout {
    background: transparent;
    border: 1px solid var(--glass-border);
    color: var(--text-main);
    padding: 8px 16px;
    border-radius: 8px;
    cursor: pointer;
    font-size: 14px;
    transition: all 0.3s;
}

.btn-logout:hover {
    background: rgba(239, 68, 68, 0.1);
    border-color: rgba(239, 68, 68, 0.2);
    color: #ef4444;
}

/* Hamburger Styles */
.hamburger {
    display: none;
    flex-direction: column;
    justify-content: space-between;
    width: 30px;
    height: 21px;
    background: transparent;
    border: none;
    cursor: pointer;
    padding: 0;
    z-index: 1001;
}

.hamburger span {
    width: 100%;
    height: 3px;
    background-color: var(--text-main);
    border-radius: 10px;
    transition: all 0.3s ease;
}

.hamburger.is-active span:nth-child(1) {
    transform: translateY(9px) rotate(45deg);
}

.hamburger.is-active span:nth-child(2) {
    opacity: 0;
}

.hamburger.is-active span:nth-child(3) {
    transform: translateY(-9px) rotate(-45deg);
}

/* Hero Section */
.hero {
    display: flex;
    align-items: center;
    padding: 160px 8% 80px;
    gap: 60px;
}

.hero-content {
    flex: 1.2;
}

.welcome-text {
    color: var(--primary);
    font-weight: 600;
    margin-bottom: 16px;
    display: inline-block;
    padding: 4px 12px;
    background: rgba(16, 185, 129, 0.1);
    border-radius: 100px;
    font-size: 14px;
}

.title {
    font-size: 64px;
    line-height: 1.1;
    font-weight: 800;
    margin-bottom: 24px;
}

.accent {
    color: var(--primary);
    text-shadow: 0 0 30px var(--primary-glow);
}

.description {
    color: var(--text-muted);
    font-size: 18px;
    line-height: 1.6;
    margin-bottom: 40px;
    max-width: 500px;
}

.cta-buttons {
    display: flex;
    gap: 16px;
}

.btn-primary {
    background: var(--primary);
    color: white;
    border: none;
    padding: 16px 32px;
    border-radius: 12px;
    font-weight: 700;
    font-size: 16px;
    cursor: pointer;
    box-shadow: 0 10px 20px var(--primary-glow);
    transition: transform 0.3s, box-shadow 0.3s;
}

.btn-primary:hover {
    transform: translateY(-2px);
    box-shadow: 0 15px 30px var(--primary-glow);
}

.btn-secondary {
    background: var(--glass-bg);
    color: var(--text-main);
    border: 1px solid var(--glass-border);
    padding: 16px 32px;
    border-radius: 12px;
    font-weight: 600;
    font-size: 16px;
    cursor: pointer;
    transition: background 0.3s;
}

.btn-secondary:hover {
    background: rgba(255, 255, 255, 0.08);
}

/* Hero Visual */
.hero-visual {
    flex: 1;
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
}

.circle-gradient {
    position: absolute;
    width: 400px;
    height: 400px;
    background: linear-gradient(135deg, var(--primary), #3b82f6);
    border-radius: 50%;
    filter: blur(80px);
    opacity: 0.15;
    animation: pulse 8s infinite alternate;
}

.glass-card {
    width: 300px;
    background: rgba(255, 255, 255, 0.05);
    backdrop-filter: blur(20px);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 24px;
    padding: 24px;
    position: relative;
    z-index: 1;
}

.floating {
    animation: floating 4s infinite ease-in-out;
}

.card-header {
    display: flex;
    gap: 6px;
    margin-bottom: 20px;
}

.dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
}

.dot.red { background: #ff5f57; }
.dot.yellow { background: #ffbd2e; }
.dot.green { background: #27c93f; }

.card-body h3 {
    font-size: 18px;
    margin-bottom: 16px;
}

.mock-item {
    display: flex;
    gap: 12px;
    align-items: center;
    background: rgba(255, 255, 255, 0.03);
    padding: 12px;
    border-radius: 12px;
    margin-bottom: 20px;
}

.img-placeholder {
    width: 40px;
    height: 40px;
    background: var(--primary);
    border-radius: 8px;
    opacity: 0.6;
}

.line {
    background: var(--glass-border);
    height: 6px;
    border-radius: 4px;
}

.line.long { width: 120px; margin-bottom: 6px; }
.line.short { width: 60px; }

.status-badge {
    background: rgba(16, 185, 129, 0.1);
    color: var(--primary);
    font-size: 12px;
    font-weight: 700;
    padding: 6px 12px;
    border-radius: 6px;
    display: inline-block;
}

/* Quick Stats */
.quick-stats {
    display: flex;
    justify-content: center;
    gap: 100px;
    padding: 60px 8%;
    border-top: 1px solid var(--glass-border);
}

.stat-item h2 {
    font-size: 32px;
    font-weight: 800;
    margin-bottom: 4px;
}

.stat-item p {
    color: var(--text-muted);
    font-size: 14px;
}

/* Animations */
@keyframes floating {
    0%, 100% { transform: translateY(0); }
    50% { transform: translateY(-20px); }
}

@keyframes pulse {
    from { transform: scale(1); opacity: 0.15; }
    to { transform: scale(1.2); opacity: 0.25; }
}

@keyframes fadeIn {
    from { opacity: 0; transform: translateY(20px); }
    to { opacity: 1; transform: translateY(0); }
}

.hero-content {
    animation: fadeIn 1s ease-out;
}

/* Mobile Responsive */
@media (max-width: 1024px) {
    .hamburger {
        display: flex;
    }

    .nav-links {
        position: fixed;
        top: 0;
        right: -100%;
        width: 100%;
        height: 100vh;
        background: #0f172a;
        backdrop-filter: blur(20px);
        flex-direction: column;
        justify-content: center;
        gap: 40px;
        transition: right 0.4s cubic-bezier(0.77, 0, 0.175, 1);
        z-index: 1000;
        padding: 8%;
    }

    .nav-links.mobile-open {
        right: 0;
    }

    .nav-links a {
        font-size: 24px;
        font-weight: 700;
        color: var(--text-main);
    }

    .btn-logout {
        width: 100%;
        padding: 16px;
        font-size: 18px;
    }

    .hero {
        flex-direction: column;
        text-align: center;
        padding-top: 140px;
    }
    
    .hero-content {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    
    .title {
        font-size: 48px;
    }
    
    .cta-buttons {
        justify-content: center;
    }

    .quick-stats {
        gap: 30px; 
        padding: 40px 5%;
        flex-wrap: wrap; 
    }

    .stat-item h2 {
        font-size: 24px; 
    }

    .stat-item p {
        font-size: 11px; 
    }
}
</style>
