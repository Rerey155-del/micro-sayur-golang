<script setup>
definePageMeta({
  layout: 'admin'
})

// Ambil konfigurasi API dari nuxt.config.ts
const config = useRuntimeConfig()

// Ambil data produk dari Product Service (port 8081)
const { data: productsData, pending, error, refresh } = await useFetch(`${config.public.apiBase}/products`)

// Fungsi pembantu untuk format mata uang Rupiah
const formatCurrency = (value) => {
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

// ------------ STATE MODAL PRODUK ------------
const isModalOpen = ref(false)
const isEditMode = ref(false)
const currentEditId = ref(null)
const isSubmitting = ref(false)
const formError = ref('')

const form = reactive({
  name: '',
  description: '',
  price: '',
  stock: '',
  image: null
})

const handleFileChange = (e) => {
  if (e.target.files.length > 0) {
    form.image = e.target.files[0]
  }
}

const openAddModal = () => {
  isEditMode.value = false
  currentEditId.value = null
  form.name = ''
  form.description = ''
  form.price = ''
  form.stock = ''
  form.image = null
  formError.value = ''
  isModalOpen.value = true
}

const openEditModal = (product) => {
  isEditMode.value = true
  currentEditId.value = product.id
  form.name = product.name
  form.description = product.description
  form.price = product.price
  form.stock = product.stock
  form.image = null // reset image field for edit
  formError.value = ''
  isModalOpen.value = true
}

const submitProduct = async () => {
  isSubmitting.value = true
  formError.value = ''
  
  const formData = new FormData()
  formData.append('name', form.name)
  formData.append('description', form.description)
  formData.append('price', form.price.toString())
  formData.append('stock', form.stock.toString())
  if (form.image) {
    formData.append('image', form.image)
  }

  const url = isEditMode.value 
    ? `${config.public.apiBase}/products/${currentEditId.value}`
    : `${config.public.apiBase}/products`
  
  const method = isEditMode.value ? 'PUT' : 'POST'

  try {
     const { error: apiError } = await useFetch(url, {
       method: method,
       body: formData
     })
     
     if (apiError.value) {
        throw new Error(apiError.value.data?.error || 'Gagal menyimpan produk')
     }
     
     isModalOpen.value = false
     await refresh()
  } catch (e) {
     formError.value = e.message
  } finally {
     isSubmitting.value = false
  }
}

const confirmDelete = async (id, name) => {
  if (confirm(`Apakah Anda yakin ingin menghapus produk "${name}"?`)) {
    try {
      const { error: apiError } = await useFetch(`${config.public.apiBase}/products/${id}`, {
        method: 'DELETE'
      })
      
      if (apiError.value) {
        throw new Error(apiError.value.data?.error || 'Gagal menghapus produk')
      }
      
      await refresh()
    } catch (e) {
      alert(e.message)
    }
  }
}
</script>

<template>
  <div class="dashboard-content">
    <div class="header">
      <h1>Kelola Produk</h1>
      <p>Manajemen stok dan harga produk Anda.</p>
    </div>

    <!-- Tampilkan Loading jika data sedang dimuat -->
    <div v-if="pending" class="glass-panel" style="text-align: center; padding: 40px;">
      <p>Sedang memuat data produk...</p>
    </div>

    <!-- Tampilkan Error jika gagal mengambil data -->
    <div v-else-if="error" class="glass-panel" style="text-align: center; color: #ef4444; padding: 40px;">
      <p>Gagal mengambil data produk: {{ error.message }}</p>
    </div>

    <!-- Tampilkan data jika sukses -->
    <div v-else class="glass-panel">
      <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px;">
        <h3>Daftar Produk ({{ productsData?.data?.length || 0 }})</h3>
        <button class="btn-primary" style="width: auto; padding: 10px 20px;" @click="openAddModal">+ Tambah Produk</button>
      </div>
      
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>Gambar</th>
              <th>ID</th>
              <th>Nama</th>
              <th>Deskripsi</th>
              <th>Harga</th>
              <th>Stok</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="p in productsData?.data" :key="p.id">
              <td>
                 <img v-if="p.image_url" :src="getImageUrl(p.image_url)" alt="Product image" style="width: 50px; height: 50px; object-fit: cover; border-radius: 8px;">
                 <div v-else style="width: 50px; height: 50px; background: #334155; border-radius: 8px; display: flex; align-items: center; justify-content: center; font-size: 10px; color: #94a3b8; text-align: center; line-height: 1;">No<br>Img</div>
              </td>
              <td>{{ p.id }}</td>
              <td>{{ p.name }}</td>
              <td style="max-width: 300px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;">
                {{ p.description || '-' }}
              </td>
              <td>{{ formatCurrency(p.price) }}</td>
              <td>
                <span :style="{ color: p.stock < 10 ? '#ef4444' : '#10b981', fontWeight: 'bold' }">
                  {{ p.stock }}
                </span>
              </td>
              <td>
                <div style="display: flex; gap: 8px;">
                  <button class="btn-edit" @click="openEditModal(p)" title="Edit">
                    Edit
                  </button>
                  <button class="btn-delete" @click="confirmDelete(p.id, p.name)" title="Hapus">
                    Hapus
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- MODAL PRODUK -->
    <div v-if="isModalOpen" class="modal-overlay" @click.self="isModalOpen = false">
      <div class="modal-content glass-panel" style="margin-top: 0; padding: 32px; width: 100%; max-width: 500px; animation: slideIn 0.3s ease-out;">
        <h3 style="margin-top: 0; margin-bottom: 24px;">{{ isEditMode ? 'Edit Produk' : 'Tambah Produk Baru' }}</h3>

        <div v-if="formError" style="background: rgba(239,68,68,0.1); color: #ef4444; padding: 12px; border-radius: 8px; margin-bottom: 20px; font-size: 14px;">
          {{ formError }}
        </div>

        <form @submit.prevent="submitProduct" class="product-form">
          <div class="form-group">
            <label>Nama Produk</label>
            <input type="text" v-model="form.name" required placeholder="Contoh: Bayam Segar">
          </div>
          
          <div class="form-group">
            <label>Harga (Rp)</label>
            <input type="number" v-model="form.price" required placeholder="Contoh: 15000">
          </div>

          <div class="form-group">
            <label>Stok</label>
            <input type="number" v-model="form.stock" required placeholder="Contoh: 50">
          </div>

          <div class="form-group">
            <label>Deskripsi Singkat</label>
            <textarea v-model="form.description" rows="3" placeholder="Informasi detail tentang sayuran ini..."></textarea>
          </div>

          <div class="form-group">
            <label>Gambar Produk</label>
            <input type="file" @change="handleFileChange" accept="image/*" style="padding: 10px 0;">
          </div>

          <div style="display: flex; gap: 12px; margin-top: 32px;">
            <button type="button" @click="isModalOpen = false" class="btn-secondary" style="flex: 1;">Batal</button>
            <button type="submit" class="btn-primary" style="flex: 1;" :disabled="isSubmitting">
              {{ isSubmitting ? 'Menyimpan...' : (isEditMode ? 'Simpan Perubahan' : 'Tambah Produk') }}
            </button>
          </div>
        </form>
      </div>
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
  background-clip: text; /* Fixed lint warning */
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

/* Button Styling */
.btn-primary {
  background: linear-gradient(135deg, #0ea5e9, #3b82f6);
  color: white;
  border: none;
  padding: 12px 24px;
  border-radius: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(14, 165, 233, 0.2);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(14, 165, 233, 0.4);
}

.btn-secondary {
  background: transparent;
  color: #94a3b8;
  border: 1px solid rgba(255, 255, 255, 0.1);
  padding: 12px 24px;
  border-radius: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.05);
  color: white;
}

.btn-edit {
  background: rgba(14, 165, 233, 0.1);
  color: #0ea5e9;
  border: 1px solid rgba(14, 165, 233, 0.2);
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  transition: 0.3s;
}

.btn-edit:hover {
  background: #0ea5e9;
  color: white;
}

.btn-delete {
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
  border: 1px solid rgba(239, 68, 68, 0.2);
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  transition: 0.3s;
}

.btn-delete:hover {
  background: #ef4444;
  color: white;
}

/* Modal Stles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(15, 23, 42, 0.8);
  backdrop-filter: blur(8px);
  z-index: 2000;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

@keyframes slideIn {
  from { opacity: 0; transform: translateY(-20px) scale(0.95); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}

.product-form .form-group {
  margin-bottom: 20px;
}

.product-form label {
  display: block;
  margin-bottom: 8px;
  color: #cbd5e1;
  font-size: 14px;
  font-weight: 500;
}

.product-form input[type="text"],
.product-form input[type="number"],
.product-form textarea {
  width: 100%;
  padding: 14px;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  color: white;
  font-family: inherit;
  transition: 0.3s;
}

.product-form input:focus,
.product-form textarea:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2);
}
</style>
