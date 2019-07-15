package app

import "fmt"

func (r *Repository) findDospem() string {
	resp := r.serviceSiakad.GetDospem(r.phoneNumber)

	if resp.Status != 1 {
		return MESSAGE[STATUS_ERROR_SYSTEM]
	}

	answer := fmt.Sprintf("Data dosen pembimbing akademik anda : \n\nNama :  %s \n\nNID :  %s \n\nNo.Tlp : %s \n", resp.Data.NamaDosen, resp.Data.Nid, resp.Data.TlpDosen)

	return answer
}

func (r *Repository) findBriva() string {

	resp := r.serviceSiakad.GetBriva(r.phoneNumber)

	if resp.Status != 1 {
		return MESSAGE[resp.Status]
	}

	answer := fmt.Sprintf("nomer briva %s adalah %s", resp.Data.Nama, resp.Data.NomorBriva)

	return answer
}

func (r *Repository) findSpp() string {

	resp := r.serviceSiakad.GetSpp(r.phoneNumber)

	if resp.Status != 1 {
		return MESSAGE[resp.Status]
	}

	answer := fmt.Sprintf("Informasi status bayaran kuliah anda saat ini adalah :\n\n Tahun Ajaran : %s \n Semester : %s \n Status : %s \n Tanggal Validasi : %s", resp.Data.Tahunajaran, resp.Data.Semester, resp.Data.Status, resp.Data.TglValidasi)

	return answer
}
