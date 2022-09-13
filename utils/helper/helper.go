package helper

func Success(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":     "Berhasil",
		"Keterangan": msg,
	}
}

func SuccessGet(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":     "Berhasil",
		"Keterangan": msg,
		"Data":       data,
	}
}

func Failed(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":     "Gagal",
		"Keterangan": msg,
	}
}

func SuccessLogin(msg string, name string, data interface{}) map[string]interface{} {
	sambut := "Selamat Datang " + name
	return map[string]interface{}{
		"status":     "Berhasil",
		"welcome":    sambut,
		"Keterangan": msg,
		"Data":       data,
	}
}
