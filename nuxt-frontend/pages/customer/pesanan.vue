<script setup>
definePageMeta({
    layout: false
})

const config = useRuntimeConfig()
const userId = useState('customer-user-id', () => 0)
const userName = useState('customer-user-name', () => 'Pelanggan')
const isMenuOpen = ref(false)

if (import.meta.client) {
    userId.value = Number(localStorage.getItem('user_id') || 0)
    userName.value = localStorage.getItem('user_name') || 'Pelanggan'
}

const ordersUrl = computed(() => `${config.public.orderApiBase}/orders/user/${userId.value || 0}`)
const { data: ordersData, pending, error, refresh } = await useFetch(ordersUrl, {
    server: false,
    default: () => ({ data: [] }),
    watch: [ordersUrl]
})

const logout = () => {
    localStorage.clear()
    navigateTo('/')
}

const toggleMenu = () => {
    isMenuOpen.value = !isMenuOpen.value
}

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

const getStatusClass = (status) => {
    switch (status) {
        case 'Selesai': return 'done'
        case 'Dikirim': return 'shipping'
        case 'Diproses': return 'process'
        case 'Dibatalkan': return 'cancel'
        default: return 'pending'
    }
}
</script>

<template>
    <div class="orders-page">
        <nav class="navbar glass">
            <div class="logo">MicroSayur</div>

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

        <section class="hero">
            <div>
                <p class="eyebrow">Riwayat Belanja</p>
                <h1>Pesanan {{ userName }}</h1>
                <p class="subtitle">Pantau status pemesanan Anda dari satu halaman.</p>
            </div>
            <button class="refresh-btn" @click="refresh">Muat Ulang</button>
        </section>

        <section v-if="pending" class="state-card glass">
            Memuat data pesanan...
        </section>

        <section v-else-if="error" class="state-card glass error-card">
            Gagal memuat pesanan: {{ error.message }}
        </section>

        <section v-else-if="!ordersData?.data?.length" class="state-card glass">
            Belum ada pesanan. Mulai belanja dari halaman produk.
        </section>

        <section v-else class="orders-grid">
            <article v-for="order in ordersData.data" :key="order.id" class="order-card glass">
                <div class="order-top">
                    <div>
                        <p class="order-code">{{ order.order_code }}</p>
                        <h3>{{ formatCurrency(order.total_amount) }}</h3>
                    </div>
                    <span class="status-pill" :class="getStatusClass(order.status)">{{ order.status }}</span>
                </div>

                <p class="meta">{{ formatDate(order.created_at) }}</p>
                <p class="address">{{ order.address || 'Alamat belum tersedia' }}</p>

                <div class="items-list">
                    <div v-for="item in order.items" :key="item.id" class="item-row">
                        <span>{{ item.product_name }} x{{ item.quantity }}</span>
                        <strong>{{ formatCurrency(item.subtotal) }}</strong>
                    </div>
                </div>

                <p v-if="order.notes" class="notes">Catatan: {{ order.notes }}</p>
            </article>
        </section>
    </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap');

.orders-page {
    --primary: #10b981;
    --bg-dark: #0f172a;
    --text-main: #f8fafc;
    --text-muted: #94a3b8;
    --glass-border: rgba(255, 255, 255, 0.08);
    min-height: 100vh;
    background:
        radial-gradient(circle at top left, rgba(16, 185, 129, 0.1), transparent 32%),
        radial-gradient(circle at bottom right, rgba(59, 130, 246, 0.12), transparent 30%),
        var(--bg-dark);
    color: var(--text-main);
    font-family: 'Plus Jakarta Sans', sans-serif;
    padding-bottom: 48px;
}

.glass {
    background: rgba(255, 255, 255, 0.04);
    backdrop-filter: blur(14px);
    border: 1px solid var(--glass-border);
}

.navbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px 8%;
    position: sticky;
    top: 0;
    z-index: 1000;
}

.logo {
    font-size: 24px;
    font-weight: 800;
    color: var(--primary);
}

.nav-links {
    display: flex;
    align-items: center;
    gap: 32px;
}

.nav-links a {
    text-decoration: none;
    color: var(--text-muted);
}

.nav-links a:hover,
.nav-links a.router-link-active {
    color: var(--primary);
}

.btn-logout,
.refresh-btn {
    border: 1px solid var(--glass-border);
    background: transparent;
    color: var(--text-main);
    padding: 10px 16px;
    border-radius: 10px;
    cursor: pointer;
}

.hero {
    padding: 56px 8% 28px;
    display: flex;
    justify-content: space-between;
    gap: 20px;
    align-items: end;
}

.eyebrow {
    margin: 0 0 10px;
    color: var(--primary);
    text-transform: uppercase;
    letter-spacing: 0.12em;
    font-size: 12px;
}

.hero h1 {
    margin: 0;
    font-size: 46px;
}

.subtitle,
.meta,
.address,
.notes {
    color: var(--text-muted);
}

.state-card {
    margin: 0 8%;
    border-radius: 24px;
    padding: 28px;
}

.error-card {
    color: #fecaca;
}

.orders-grid {
    padding: 0 8%;
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
    gap: 24px;
}

.order-card {
    border-radius: 24px;
    padding: 24px;
}

.order-top,
.item-row {
    display: flex;
    justify-content: space-between;
    gap: 12px;
}

.order-code {
    margin: 0 0 8px;
    color: #86efac;
    font-size: 13px;
    letter-spacing: 0.08em;
}

.order-card h3 {
    margin: 0;
    font-size: 28px;
}

.status-pill {
    height: fit-content;
    padding: 8px 12px;
    border-radius: 999px;
    font-size: 12px;
    font-weight: 700;
}

.status-pill.pending {
    background: rgba(245, 158, 11, 0.16);
    color: #fcd34d;
}

.status-pill.process {
    background: rgba(14, 165, 233, 0.16);
    color: #7dd3fc;
}

.status-pill.shipping {
    background: rgba(59, 130, 246, 0.16);
    color: #93c5fd;
}

.status-pill.done {
    background: rgba(16, 185, 129, 0.16);
    color: #86efac;
}

.status-pill.cancel {
    background: rgba(239, 68, 68, 0.16);
    color: #fca5a5;
}

.items-list {
    margin-top: 18px;
    display: grid;
    gap: 12px;
}

.item-row {
    padding: 12px 14px;
    background: rgba(255, 255, 255, 0.04);
    border-radius: 14px;
}

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
        background: var(--bg-dark);
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

    .hero {
        flex-direction: column;
        align-items: flex-start;
    }

    .hero h1 {
        font-size: 36px;
    }
}
</style>
