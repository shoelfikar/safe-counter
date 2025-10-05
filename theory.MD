
# Safe Counter

Safe counter merupakan cara/mekanisme yang digunakan oleh goroutine untuk melakukan counter pada suatu data yang di akses oleh banyak goroutine secara bersamaan untuk menghindari terjadinya race condition.

Berikut ini beberapa cara Untuk membuatnya thread-safe: 

## Menggunakan sync.Mutex

Mutex (mutual exclusion) adalah mekanisme lock di Go.

Saat satu goroutine mengakses atau mengubah data, goroutine lain harus menunggu hingga lock dilepas.

### Kenapa Menggunakan sync.Mutex

Karena Mutex cocok untuk data kompleks (misalnya struct dengan banyak field) karena bisa melindungi lebih dari satu variabel.

Contoh teori:

Inc() → menambah nilai counter dengan melakukan lock mutex pada data yang akan di update.

### Contoh Penggunaan

```golang
func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}
```

Value() → membaca nilai counter dengan cara yang sama.

### Contoh Penggunaan

```golang
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
```

## Kesimpulan

Mutex cocok untuk data kompleks (misalnya struct dengan banyak field) karena bisa melindungi lebih dari satu variabel.

Kita juga bisa menggunakan Atomic operation untuk counter sederhana, lebih cepat karena tidak ada mekanisme lock/unlock.

Jadi, Safe Counter di Go adalah implementasi counter yang bisa diakses oleh banyak goroutine secara bersamaan tanpa race condition, jika menggunakan struct yang kompleks, gunakan mutex dan gunakan atomic operations untuk operasi yang lebih sederhana.

Dengan begitu, walaupun ada ratusan goroutine yang melakukan akses terhadap data yang sama, tidak akan terjadi race condition.

