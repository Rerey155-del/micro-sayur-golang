<script setup>
definePageMeta({
    layout: false
})

const config = useRuntimeConfig()

// Ambil data produk
const { data: productsData, pending, error } = await useFetch(`${config.public.apiBase}/products`)

const isOrderModalOpen = ref(false)
const isSubmittingOrder = ref(false)
const orderError = ref('')
const orderSuccess = ref('')
const selectedProduct = ref(null)

const orderForm = reactive({
  quantity: 1,
  address: '',
  notes: ''
})

// Format mata uang
const formatCurrency = (value) => {
  if (!value) return 'Rp 0'
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0
  }).format(value)
}

// Helper untuk URL gambar static
const getImageUrl = (path) => {
  if (!path) return ''
  const baseUrl = config.public.apiBase.replace('/api/v1', '')
  return `${baseUrl}${path}`
}

const logout = () => {
    localStorage.clear()
    navigateTo('/')
}

const openOrderModal = (product) => {
    selectedProduct.value = product
    orderForm.quantity = 1
    orderForm.address = ''
    orderForm.notes = ''
    orderError.value = ''
    orderSuccess.value = ''
    isOrderModalOpen.value = true
}

const closeOrderModal = () => {
    isOrderModalOpen.value = false
    selectedProduct.value = null
}

const selectedSubtotal = computed(() => {
    if (!selectedProduct.value) return 0
    return Number(selectedProduct.value.price) * Number(orderForm.quantity || 0)
})

const submitOrder = async () => {
    if (!selectedProduct.value) return

    const userId = Number(localStorage.getItem('user_id') || 0)
    const userName = localStorage.getItem('user_name') || 'Pelanggan'

    if (!userId) {
        orderError.value = 'Sesi login tidak lengkap. Silakan login ulang.'
        return
    }

    isSubmittingOrder.value = true
    orderError.value = ''
    orderSuccess.value = ''

    try {
        await $fetch(`${config.public.orderApiBase}/orders`, {
            method: 'POST',
            body: {
                user_id: userId,
                customer_name: userName,
                address: orderForm.address,
                notes: orderForm.notes,
                items: [
                    {
                        product_id: selectedProduct.value.id,
                        product_name: selectedProduct.value.name,
                        product_price: Number(selectedProduct.value.price),
                        quantity: Number(orderForm.quantity)
                    }
                ]
            }
        })

        orderSuccess.value = 'Pesanan berhasil dibuat.'
        setTimeout(() => {
            closeOrderModal()
            navigateTo('/customer/pesanan')
        }, 900)
    } catch (err) {
        orderError.value = err?.data?.error || err?.message || 'Gagal membuat pesanan'
    } finally {
        isSubmittingOrder.value = false
    }
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

        <section class="products-section" id="produk">
            <div class="section-header">
                <h2 class="section-title">Produk Segar Kami <span class="accent">.</span></h2>
                <p class="section-subtitle">Sayuran organik pilihan langsung dari kebun, dipanen setiap pagi.</p>
            </div>

            <div v-if="pending" class="loading-state">
                <p>Memuat produk berkualitas untuk Anda...</p>
            </div>

            <div v-else-if="error" class="error-state">
                <p>Waduh, gagal memuat produk. Coba refresh halaman ini, ya!</p>
            </div>

            <div v-else class="products-grid">
                <div v-for="product in productsData?.data" :key="product.id" class="product-card glass">
                    <div class="product-image">
                        <!-- Show image if exists -->
                        <img v-if="product.image_url" :src="getImageUrl(product.image_url)" alt="Product" style="width: 100%; height: 100%; object-fit: cover;">
                        
                        <!-- Placeholder if no image -->
                        <div v-else class="img-badge">#{{ product.id }} (Tanpa Gambar)</div>
                    </div>
                    <div class="product-info">
                        <h3 class="product-name">{{ product.name }}</h3>
                        <p class="product-desc">{{ product.description || 'Sayuran segar berkualitas tinggi.' }}</p>
                        <div class="product-meta">
                            <span class="product-price">{{ formatCurrency(product.price) }}</span>
                            <span class="product-stock" :class="{ 'low-stock': product.stock < 10 }">
                                Sisa Stok: {{ product.stock }}
                            </span>
                        </div>
                        <button class="btn-add-cart" :disabled="product.stock < 1" @click="openOrderModal(product)">
                            {{ product.stock < 1 ? 'Stok Habis' : 'Beli Sekarang' }}
                        </button>
                    </div>
                </div>
            </div>
        </section>

        <div v-if="isOrderModalOpen && selectedProduct" class="modal-overlay" @click.self="closeOrderModal">
            <div class="order-modal glass">
                <div class="modal-head">
                    <div>
                        <p class="eyebrow">Checkout Cepat</p>
                        <h3>{{ selectedProduct.name }}</h3>
                    </div>
                    <button class="modal-close" @click="closeOrderModal">×</button>
                </div>

                <div class="order-summary">
                    <span>Harga</span>
                    <strong>{{ formatCurrency(selectedProduct.price) }}</strong>
                </div>

                <div v-if="orderError" class="alert alert-error">{{ orderError }}</div>
                <div v-if="orderSuccess" class="alert alert-success">{{ orderSuccess }}</div>

                <form class="order-form" @submit.prevent="submitOrder">
                    <label>
                        Jumlah
                        <input v-model="orderForm.quantity" type="number" min="1" :max="selectedProduct.stock" required>
                    </label>

                    <label>
                        Alamat Pengiriman
                        <textarea v-model="orderForm.address" rows="3" placeholder="Contoh: Jl. Melati No. 12, Jakarta" required></textarea>
                    </label>

                    <label>
                        Catatan
                        <textarea v-model="orderForm.notes" rows="2" placeholder="Opsional, misalnya minta dikirim sore"></textarea>
                    </label>

                    <div class="order-total">
                        <span>Total</span>
                        <strong>{{ formatCurrency(selectedSubtotal) }}</strong>
                    </div>

                    <div class="modal-actions">
                        <button type="button" class="btn-secondary" @click="closeOrderModal">Batal</button>
                        <button type="submit" class="btn-add-cart" :disabled="isSubmittingOrder">
                            {{ isSubmittingOrder ? 'Memproses...' : 'Buat Pesanan' }}
                        </button>
                    </div>
                </form>
            </div>
        </div>
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
}

.glass {
    background: rgba(255, 255, 255, 0.03);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(255, 255, 255, 0.08);
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

.products-section {
    padding: 120px 8% 80px; 
}

.section-header {
    text-align: center;
    margin-bottom: 50px;
}

.section-title {
    font-size: 36px;
    font-weight: 800;
    margin-bottom: 12px;
}

.accent {
    color: var(--primary);
    text-shadow: 0 0 30px var(--primary-glow);
}

.section-subtitle {
    color: var(--text-muted);
    font-size: 16px;
    max-width: 600px;
    margin: 0 auto;
}

.loading-state, .error-state {
    text-align: center;
    padding: 40px;
    background: rgba(255, 255, 255, 0.03);
    border-radius: 16px;
    border: 1px solid rgba(255, 255, 255, 0.08);
    color: var(--text-muted);
}

.error-state {
    color: #ef4444;
}

.products-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 30px;
}

.product-card {
    border-radius: 20px;
    overflow: hidden;
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    display: flex;
    flex-direction: column;
}

.product-card:hover {
    transform: translateY(-10px);
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
    border-color: rgba(16, 185, 129, 0.3);
}

.product-image {
    height: 180px;
    background: linear-gradient(135deg, rgba(16, 185, 129, 0.2), rgba(59, 130, 246, 0.2));
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
}

.img-badge {
    position: absolute;
    top: 16px;
    right: 16px;
    background: rgba(0, 0, 0, 0.5);
    backdrop-filter: blur(4px);
    padding: 4px 10px;
    border-radius: 20px;
    font-size: 12px;
    font-weight: 700;
}

.product-info {
    padding: 24px;
    display: flex;
    flex-direction: column;
    flex: 1;
}

.product-name {
    font-size: 20px;
    font-weight: 700;
    margin-bottom: 8px;
    color: #fff;
}

.product-desc {
    color: var(--text-muted);
    font-size: 14px;
    line-height: 1.5;
    margin-bottom: 20px;
    flex: 1;
}

.product-meta {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}

.product-price {
    font-size: 20px;
    font-weight: 800;
    color: var(--primary);
}

.product-stock {
    font-size: 13px;
    color: var(--text-muted);
    background: rgba(255, 255, 255, 0.05);
    padding: 4px 10px;
    border-radius: 8px;
}

.product-stock.low-stock {
    color: #ef4444;
    background: rgba(239, 68, 68, 0.1);
}

.btn-add-cart {
    width: 100%;
    background: rgba(16, 185, 129, 0.1);
    color: var(--primary);
    border: 1px solid rgba(16, 185, 129, 0.3);
    padding: 12px;
    border-radius: 12px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s;
}

.btn-add-cart:hover {
    background: var(--primary);
    color: white;
}

.btn-add-cart:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.btn-secondary {
    width: 100%;
    background: transparent;
    color: var(--text-main);
    border: 1px solid var(--glass-border);
    padding: 12px;
    border-radius: 12px;
    font-weight: 600;
    cursor: pointer;
}

.modal-overlay {
    position: fixed;
    inset: 0;
    background: rgba(15, 23, 42, 0.82);
    backdrop-filter: blur(10px);
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 24px;
    z-index: 1200;
}

.order-modal {
    width: min(100%, 520px);
    border-radius: 24px;
    padding: 28px;
}

.modal-head {
    display: flex;
    justify-content: space-between;
    gap: 16px;
    margin-bottom: 24px;
}

.modal-head h3 {
    margin: 4px 0 0;
    font-size: 28px;
}

.eyebrow {
    margin: 0;
    color: var(--primary);
    font-size: 13px;
    text-transform: uppercase;
    letter-spacing: 0.12em;
}

.modal-close {
    width: 42px;
    height: 42px;
    border-radius: 50%;
    border: 1px solid var(--glass-border);
    background: transparent;
    color: var(--text-main);
    font-size: 24px;
    cursor: pointer;
}

.order-summary,
.order-total {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 18px;
    padding: 14px 16px;
    background: rgba(255, 255, 255, 0.04);
    border-radius: 14px;
}

.order-form {
    display: grid;
    gap: 16px;
}

.order-form label {
    display: grid;
    gap: 8px;
    color: var(--text-muted);
    font-size: 14px;
}

.order-form input,
.order-form textarea {
    width: 100%;
    border: 1px solid var(--glass-border);
    background: rgba(0, 0, 0, 0.18);
    color: var(--text-main);
    border-radius: 12px;
    padding: 14px 16px;
    font-family: inherit;
}

.modal-actions {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 12px;
}

.alert {
    margin-bottom: 16px;
    padding: 12px 14px;
    border-radius: 12px;
    font-size: 14px;
}

.alert-error {
    color: #fecaca;
    background: rgba(239, 68, 68, 0.16);
    border: 1px solid rgba(239, 68, 68, 0.25);
}

.alert-success {
    color: #bbf7d0;
    background: rgba(16, 185, 129, 0.16);
    border: 1px solid rgba(16, 185, 129, 0.25);
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
        background: var(--bg-dark);
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

    .modal-actions {
        grid-template-columns: 1fr;
    }
}
</style>
