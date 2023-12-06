package request

type Booking struct {
	TglPinjam       string `json:"tgl_pinjam"`
	UserID          int    `json:"user_id"`
	TglKembali      string `json:"tgl_kembali"`
	TglPengembalian string `json:"tgl_pengembalian"`
	Status          string `json:"status"`
	TotalDenda      int    `json:"total_denda"`
}
