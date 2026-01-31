package validator

import (
	"errors"

	enLocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

// Validator adalah struct kustom untuk membungkus library validator pihak ketiga.
// Kita menggunakan ini untuk memvalidasi data input (misalnya: email harus berformat email).
type Validator struct {
	Validator  *validator.Validate // Library utama validator.
	Translator ut.Translator       // Penerjemah untuk pesan error (misal: ke Bahasa Inggris).
}

// NewValidator membuat instance validator baru dan mengatur penerjemah bahasa Inggris secara default.
func NewValidator() *Validator {
	en := enLocale.New()
	uni := ut.New(en, en)

	// Mencoba mendapatkan penerjemah Bahasa Inggris (en).
	trans, found := uni.GetTranslator("en")
	if !found {
		return nil
	}

	validate := validator.New()

	return &Validator{
		Validator:  validate,
		Translator: trans,
	}
}

// Validate melakukan pemeriksaan pada struct data yang diberikan.
// Jika ada field yang tidak sesuai aturan (tag `validate` di struct), fungsi ini mengembalikan error.
func (v *Validator) Validate(data interface{}) error {
	// Melakukan validasi struct.
	err := v.Validator.Struct(data)
	if err != nil {
		// Jika ada error validasi, kita parsing error-nya.
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		// Loop melalui semua error yang ditemukan, tapi kita hanya kembalikan error yang pertama.
		for _, e := range validationErrors {
			// Mencatat error ke log untuk debugging.
			log.Infof("[Validator] %s : %s", e.Field(), e.Translate(v.Translator))
			// Mengembalikan error dalam bentuk teks yang sudah diterjemahkan.
			return errors.New(e.Translate(v.Translator))
		}
	}
	// Mengembalikan nil jika tidak ada error (validasi sukses).
	return nil
}
