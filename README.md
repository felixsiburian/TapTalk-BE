# TapTalk-BE
## Read the assumption.txt first
## You can Find The API Documentation Here : https://documenter.getpostman.com/view/4760244/TVzVhFor

### How to Deploy this Backend
1. Buat sebuah Database di server anda mis: Azure, GCP, AWS
2. Ubah Connection String pada folder database>connection.go
3. Import database yang telah disediakan pada repository ini (database.pgsql)
4. Build dockerfile yang telah disediakan dengan menjalankan perintah ***go build --tag nama_image:version .*** 
		n.b : **1.** nama_image dan version anda yang tentukan
					**2.** titik diakhir command diikutkan
5. Setelah image selesai di build, push image anda ke registry. Jika anda ingin image anda bersifat public, anda dapat menggunakan docker hub.
	 Jika tidak, anda dapat menggunakan registry pada cloud yg anda gunakan (misalnya: anda menggunakan Azure, anda dapat menggunakan Azure Container Registry)
6. Setelah image anda berhasil di push ke registry yang anda mau, selanjutnya anda hanya tinggal menghubungkan image yang ada pada registry anda ke app service yang anda gunakan. Pastikan tag yang anda gunakan adalah tag yang anda push terakhir ke registry anda.
7. Selanjutnya anda hanya tinggal mencoba menggunakan API yang telah di upload ke cloud dengan cara mengganti base_url(localhost) dengan url yang disediakan oleh cloud yang anda gunakan.
