<script setup>
definePageMeta({
    layout : false 
})

// Definisikan state reaktif untuk input dan error
const email = ref('')
const password = ref('')
const error = ref(null)

const doLogin = async () => {
    error.value = null // Reset error tiap kali mencoba login
    
    try {
        // Melakukan request POST ke User Service (port 8085)
        // Pastikan User Service sudah menyala!
        const response = await $fetch('http://localhost:8085/signin', {
            method: 'POST',
            body: {
                email: email.value,
                password: password.value
            }
        })

        // Jika login sukses (mendapat token)
        if (response.data && response.data.access_token) {
            // Simpan token ke dalam storage browser (agar kedetect di Application > Local Storage)
            localStorage.setItem('token', response.data.access_token)
            localStorage.setItem('user_name', 'Super Admin')
            
            // Redirect ke halaman admin
            navigateTo('/admin')
        }
    } catch (err) {
        // Jika data gagal/tidak cocok dari database
        console.error("Login Error:", err)
        error.value = "Email atau Password salah! Silakan coba lagi."
        
        // Mengosongkan password agar user bisa mencoba kembali
        password.value = ''
    }
}
</script>

<template>
    <div class="login-container">
        <div class="login-card glass">
            <h1 class="title">Login</h1>
            <p class="subtitle">Buka aktivitas anda disini</p>
            
            <!-- Tampilkan pesan error jika login gagal -->
            <div v-if="error" class="error-alert">
                {{ error }}
            </div>

            <form @submit.prevent="doLogin" class="form" >
                <div class="input-group">
                    <label>Email Admin</label>
                    <!-- v-model menghubungkan input dengan variabel di script -->
                    <input v-model="email" type="email" placeholder="admin@microsayur.com" required>
                </div>
                <div class="input-group">
                    <label>Kata Sandi</label>
                    <input v-model="password" type="password" placeholder="••••••••" required>
                </div>
                <button type="submit" class="btn-primary">Masuk & Dapatkan Token</button>
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

@keyframes shake {
    0%, 100% { transform: translateX(0); }
    25% { transform: translateX(-5px); }
    75% { transform: translateX(5px); }
}
</style>

