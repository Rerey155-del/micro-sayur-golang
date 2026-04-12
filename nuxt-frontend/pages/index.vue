<script setup>
definePageMeta({
    layout : false 
})

// State untuk toggle Login / Register
const isLoginMode = ref(true)

// State untuk form
const name = ref('')
const email = ref('')
const password = ref('')
const role = ref('Customer') // Default role
const error = ref(null)
const successMsg = ref(null)

const toggleMode = () => {
    isLoginMode.value = !isLoginMode.value
    error.value = null
    successMsg.value = null
    password.value = ''
}

const submitForm = async () => {
    error.value = null
    successMsg.value = null

    if (isLoginMode.value) {
        // Mode Login
        try {
            const response = await $fetch('http://localhost:8085/signin', {
                method: 'POST',
                body: { email: email.value, password: password.value }
            })

            if (response.data && response.data.access_token) {
                localStorage.setItem('token', response.data.access_token)
                localStorage.setItem('user_id', String(response.data.id || ''))
                localStorage.setItem('user_name', response.data.name || 'User')
                localStorage.setItem('user_email', response.data.email || '')
                localStorage.setItem('user_role', response.data.role)
                
                // Cek role untuk navigasi
                if (response.data.role === 'Customer') {
                    navigateTo('/customer/home') // Halaman landing page untuk user/customer
                } else {
                    navigateTo('/admin') // Halaman dashboard admin
                }
            }
        } catch (err) {
            console.error("Login Error:", err)
            error.value = "Email atau Password salah! Silakan coba lagi."
            password.value = ''
        }
    } else {
        // Mode Register
        try {
            const response = await $fetch('http://localhost:8085/signup', {
                method: 'POST',
                body: { 
                    name: name.value,
                    email: email.value, 
                    password: password.value,
                    role_name: role.value
                }
            })

            if (response.message) {
                successMsg.value = "Pendaftaran berhasil! Silakan login."
                // Beralih kembali ke mode login setelah sukses
                setTimeout(() => {
                    isLoginMode.value = true
                    password.value = ''
                }, 2000)
            }
        } catch (err) {
            console.error("Register Error:", err)
            error.value = "Gagal mendaftar, pastikan data benar dan email belum digunakan."
        }
    }
}
</script>

<template>
    <div class="login-container">
        <div class="login-card glass">
            <h1 class="title">{{ isLoginMode ? 'Login' : 'Register' }}</h1>
            <p class="subtitle">{{ isLoginMode ? 'Buka aktivitas anda disini' : 'Buat akun baru untuk memulai' }}</p>
            
            <!-- Tampilkan pesan error atau sukses -->
            <div v-if="error" class="error-alert">
                {{ error }}
            </div>
            <div v-if="successMsg" class="success-alert">
                {{ successMsg }}
            </div>

            <!-- Tab Toggle -->
            <div class="toggle-container">
                <button :class="['toggle-btn', { active: isLoginMode }]" @click.prevent="toggleMode">Login</button>
                <button :class="['toggle-btn', { active: !isLoginMode }]" @click.prevent="toggleMode">Register</button>
            </div>

            <form @submit.prevent="submitForm" class="form" >
                
                <div v-if="!isLoginMode" class="input-group">
                    <label>Nama Lengkap</label>
                    <input v-model="name" type="text" placeholder="John Doe" required>
                </div>
                
                <div class="input-group">
                    <label>Email</label>
                    <input v-model="email" type="email" placeholder="admin@microsayur.com" required>
                </div>
                
                <div class="input-group">
                    <label>Kata Sandi</label>
                    <input v-model="password" type="password" placeholder="••••••••" required>
                </div>

                <div v-if="!isLoginMode" class="input-group">
                    <label>Daftar Sebagai (Role)</label>
                    <select v-model="role" class="custom-select" required>
                        <option value="Customer">Customer</option>
                        <option value="Super Admin">Super Admin</option>
                    </select>
                </div>

                <button type="submit" class="btn-primary">
                    {{ isLoginMode ? 'Masuk & Dapatkan Token' : 'Daftar Sekarang' }}
                </button>
            </form>
        </div>
    </div>
</template>

<style>
/* Tambahan CSS untuk alert error agar terlihat premium */
.error-alert {
    background: rgba(239, 68, 68, 0.1);
    border: 1px solid rgba(239, 68, 68, 0.4);
    color: #ef4444;
    padding: 12px;
    border-radius: 8px;
    font-size: 14px;
    margin-bottom: 20px;
    text-align: center;
    animation: shake 0.5s ease-in-out;
}

.success-alert {
    background: rgba(16, 185, 129, 0.1);
    border: 1px solid rgba(16, 185, 129, 0.4);
    color: #10b981;
    padding: 12px;
    border-radius: 8px;
    font-size: 14px;
    margin-bottom: 20px;
    text-align: center;
}

.toggle-container {
    display: flex;
    background: rgba(255, 255, 255, 0.05);
    border-radius: 12px;
    padding: 4px;
    margin-bottom: 24px;
}

.toggle-btn {
    flex: 1;
    padding: 10px;
    border: none;
    background: transparent;
    color: #a0aec0;
    font-size: 14px;
    font-weight: 600;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s ease;
}

.toggle-btn.active {
    background: rgba(59, 130, 246, 0.2);
    color: #acc1ff; /* Accent color matching premium look */
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.custom-select {
    width: 100%;
    padding: 12px 16px;
    background: rgba(255, 255, 255, 0.05); /* Match the dark premium inputs */
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 12px;
    font-size: 14px;
    color: #ffffff; /* Assuming dark mode based on 'premium' / 'glass' classes */
    outline: none;
    transition: all 0.3s ease;
    appearance: none;
}
.custom-select option {
    background: #1e293b; /* Dark dropdown options */
    color: #fff;
}
.custom-select:focus {
    border-color: rgba(59, 130, 246, 0.5);
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2);
}

@keyframes shake {
    0%, 100% { transform: translateX(0); }
    25% { transform: translateX(-5px); }
    75% { transform: translateX(5px); }
}
</style>

